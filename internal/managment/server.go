package managment

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"spoke7-go/internal/managment/config"
	"spoke7-go/internal/managment/controller"
	"spoke7-go/internal/managment/pb"
	"spoke7-go/internal/managment/repository"
	"spoke7-go/internal/managment/services"

	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"spoke7-go/pkg/request_id"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var once sync.Once

type ManagmentServer struct {
	version  string
	logger   logger.Logger
	e        *echo.Echo
	grpcConn *grpc.ClientConn
	conf     config.AppConfiguration

	//external client
	casdoorClient *casdoorsdk.Client
	enforcer      casbin.IEnforcer
	authz         *authz.Authz

	//service
	userService  services.UserService
	groupService services.GroupService
	roleService  services.RoleService

	// controllers
	userController  pb.UserServiceServer
	groupController pb.GroupServiceServer
	roleController  pb.RoleServiceServer
}

func NewManagmentServer(e *echo.Echo, conf config.AppConfiguration, logger logger.Logger, version string) *ManagmentServer {

	casdoorClient := casdoorsdk.NewClient(
		conf.Authentication.Endpoint,
		conf.Authentication.ClientId,
		conf.Authentication.ClientSecret,
		conf.Authentication.Certificate,
		conf.Authentication.OrganizationName,
		conf.Authentication.ApplicationName,
	)

	//cc := NewCasdoorClient(config.AppConfig.Casdoor)
	dbconn, err := InitDBConnection(logger, config.AppConfig.Database)
	if err != nil {
		logger.Fatal("Error db connections")
	}

	authzConf := authz.AuthzConfig{
		ModelPath:     config.AppConfig.Authorization.ModelPath,
		InitRulePath:  config.AppConfig.Authorization.InitRulePath,
		JwksUrl:       config.AppConfig.Authentication.JwksUrl,
		AdminRoleName: "admin",
	}
	authz := authz.NewAuthz(authzConf, dbconn, logger)
	err = authz.InitCasbin()
	if err != nil {
		logger.Fatalf("Failed to start server. Error:%v\n", err)
	}

	// initialize services

	userService := services.NewUserService(casdoorClient, authz.GetEnforcer(), conf.Authentication.OrganizationName)
	roleService := services.NewRoleService(authz.GetEnforcer())
	groupService := services.NewGroupService(conf.Authentication.OrganizationName, casdoorClient)

	//initialize controllers
	userController := controller.NewUserController(userService, authz.GetEnforcer(), logger)
	groupController := controller.NewGroupController(groupService, logger)
	roleController := controller.NewRoleController(roleService, logger)

	//authz.RegisterController(groupController)

	return &ManagmentServer{
		e:       e,
		version: version,
		logger:  logger,
		conf:    conf,

		//external client
		casdoorClient: casdoorClient,
		enforcer:      authz.GetEnforcer(),
		authz:         authz,

		// services
		userService:  userService,
		groupService: groupService,
		roleService:  roleService,
		// controllers
		userController:  userController,
		groupController: groupController,
		roleController:  roleController,
	}
}

func (s *ManagmentServer) Start() error {

	s.logger.Infof("Starting Managment Server version %s", s.version)
	s.logger.Infof("Loading configuration")
	//print port
	s.logger.Infof("Port: %d", s.conf.Service.Port)
	s.logger.Infof("GRPC Port: %d", s.conf.Service.GrpcPort)

	err := s.startGrpc(s.conf.Service.GrpcPort, s.conf.Service.Host)
	if err != nil {
		return err
	}

	err = s.connectToGrpcServer(context.Background())
	if err != nil {
		return err
	}

	// Load echo middlewares
	s.loadMiddlewares()

	// Load routes
	s.LoadRoutes()

	host := config.AppConfig.Service.Host
	port := config.AppConfig.Service.Port
	s.logger.Infof("Starting server on %s:%d", host, port)
	return s.e.Start(fmt.Sprintf("%s:%d", host, port))
}

func (s *ManagmentServer) startGrpc(grpcPort int64, host string) error {
	once.Do(func() {
		//grpcPort := config.AppConfig.Service.GrpcPort
		address := fmt.Sprintf("%s:%d", host, grpcPort)
		s.logger.Infof("Starting gRPC server on %d", grpcPort)

		go func() {
			lis, err := net.Listen("tcp", address)
			if err != nil {
				return
			}

			jwtAuthConfig := authz.JWTAuthInterceptorConfig{
				JwksUrl:          config.AppConfig.Authentication.JwksUrl,
				UsernameClaim:    config.AppConfig.Authentication.UsernameClaim,
				RolesClaim:       config.AppConfig.Authentication.RolesClaim,
				GroupsClaim:      config.AppConfig.Authentication.GroupsClaim,
				OrganizationName: config.AppConfig.Authentication.OrganizationName,
			}

			grpcServer := grpc.NewServer(
				grpc.UnaryInterceptor(
					grpc_middleware.ChainUnaryServer(
						LoggingInterceptor(s.logger),
						authz.JWTAuthInterceptor(config.AppConfig.Authentication.SkipUrls, jwtAuthConfig, s.logger),
						authz.ControllerInterceptorABAC(config.AppConfig.Authentication.SkipUrls, s.enforcer, s.logger),
						request_id.RequestIDUnaryServerInterceptor(),
					),
				),
				//jwtInterceptor(config.AppConfig.Authentication.SkipUrls, s.enforcer, s.logger)
				// ),
				//grpc.UnaryInterceptor(enforcer(s.enforcer, s.logger)),
			)

			// Register reflection service on gRPC server.
			reflection.Register(grpcServer)

			// register grpc services
			pb.RegisterUserServiceServer(grpcServer, s.userController)
			pb.RegisterGroupServiceServer(grpcServer, s.groupController)
			pb.RegisterRoleServiceServer(grpcServer, s.roleController)

			// Serve gRPC server
			grpcServer.Serve(lis)
		}()
	})

	return nil
}

// func enforcer(enforcer *casbin.Enforcer, logger logger.Logger) grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
// 		// Extract metadata
// 		md, ok := metadata.FromIncomingContext(ctx)
// 		if !ok {
// 			return nil, errors.New("missing metadata")
// 		}

// 		// Extract the object and action from the gRPC method name
// 		method := info.FullMethod
// 		parts := strings.Split(method, "/")
// 		if len(parts) < 3 {
// 			return nil, status.Errorf(codes.Internal, "invalid method name: %s", method)
// 		}
// 		domain := parts[1]
// 		object := parts[2]

// 		//log debug info
// 		logger.Debugf("Domain: %s, Action: %v", domain, object)
// 		logger.Debug("Method: %s", md)

// 		// Call the handler if authentication passes
// 		return handler(ctx, req)

// 	}
// }

func LoggingInterceptor(logger logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		logger.Infof("gRPC request: %s", info.FullMethod)
		resp, err := handler(ctx, req)
		logger.Infof("gRPC response: %s, error: %v", info.FullMethod, err)
		return resp, err
	}
}

// func jwtInterceptor(urls []string, enforcer *casbin.Enforcer, logger logger.Logger) grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

// 		// Extract metadata
// 		md, ok := metadata.FromIncomingContext(ctx)
// 		if !ok {
// 			return nil, errors.New("missing metadata")
// 		}
// 		// Extract the HTTP request path from metadata
// 		httpPath := ""
// 		if pathValues := md.Get("path"); len(pathValues) > 0 {
// 			httpPath = pathValues[0]
// 		}
// 		if isSkipUrls(httpPath, urls) {
// 			return handler(ctx, req)
// 		}

// 		// Get token from authorization header
// 		authHeader := md["authorization"]
// 		if len(authHeader) == 0 {
// 			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
// 		}

// 		// Extract token (format: Bearer <token>)
// 		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
// 		if tokenString == authHeader[0] {
// 			return nil, status.Errorf(codes.Unauthenticated, "malformed authorization header")
// 		}

// 		// Load JWKS
// 		k, err := keyfunc.NewDefaultCtx(context.Background(), []string{config.AppConfig.Authentication.JwksUrl})
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Parse and validate the JWT
// 		token, err := jwt.Parse(tokenString, k.Keyfunc)
// 		if err != nil || !token.Valid {
// 			return nil, errors.New("invalid or expired token")
// 		}

// 		// get the claims
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			return nil, errors.New("invalid claims")
// 		}

// 		var roles []string
// 		switch v := claims[config.RolesClaim].(type) {
// 		case string:
// 			roles = strings.Split(v, ",")
// 		case []interface{}:
// 			for _, role := range v {
// 				roles = append(roles, role.(string))
// 			}
// 		case []string:
// 			roles = v
// 		default:
// 			return nil, errors.New("invalid roles claim in jwt")
// 		}

// 		user := authz.User{
// 			Username: claims[config.AppConfig.Authentication.UsernameClaim].(string),
// 			Groups:   roles,
// 		}
// 		ctx = context.WithValue(ctx, authz.UserCtxKey, user)
// 		ctx = context.WithValue(ctx, "authorization", authHeader)

// 		//log debug info
// 		logger.Debugf("User: %s, Roles: %v", user.Username, user.Groups)

// 		// Call the handler if authentication passes
// 		return handler(ctx, req)
// 	}
// }

// connectToGrpcServer connects to the gRPC server.
func (s *ManagmentServer) connectToGrpcServer(ctx context.Context) error {

	url := fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, config.AppConfig.Service.GrpcPort)
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		return err
	}
	s.grpcConn = conn

	return nil
}

func (s *ManagmentServer) loadMiddlewares() {
	s.e.HideBanner = true
	s.e.Use(
		middleware.CORS(),
		middleware.Recover(),
		request_id.RequestIDMiddleware,
	)
}

func (s *ManagmentServer) LoadRoutes() error {

	// Serve Swagger UI at /swagger
	s.e.GET("/swagger/*", echoSwagger.EchoWrapHandler(
		echoSwagger.DocExpansion("none"),
		echoSwagger.URL("/swagger/doc.json"), // Directly point to your endpoint
	))

	// Serve the OpenAPI file
	s.e.GET("/swagger/doc.json", func(c echo.Context) error {
		return c.File("assets/swagger/metadata.swagger.json")
	})

	marshaler := &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: true},
	}

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption("multipart/form-data", &utils.MultipartFormPb{
			Marshaler: marshaler,
		}),
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := make(map[string]string)
			if method, ok := runtime.RPCMethod(ctx); ok {
				md["method"] = method
			}
			if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
				md["path"] = pattern
			}
			if requestID := r.Header.Get("X-Request-Id"); requestID != "" {
				md["x-request-id"] = requestID
			}
			return metadata.New(md)
		}),
	)

	/* Register handler */

	err := pb.RegisterUserServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		s.logger.Errorf("failed to register gRPC gateway: %v", err)
		return err
	}
	err = pb.RegisterGroupServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		s.logger.Errorf("failed to register gRPC gateway: %v", err)
		return err
	}

	err = pb.RegisterRoleServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		s.logger.Errorf("failed to register gRPC gateway: %v", err)
		return err
	}

	s.e.Any("/*", echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gwmux.ServeHTTP(w, r)
	})))

	return nil
}

type Object struct {
	Type   string
	Owner  string
	Groups []string
}

func InitDBConnection(logger logger.Logger, config config.DatabaseConfig) (db *gorm.DB, err error) {
	dsn := config.DSN()

	dbLogger := &repository.DBLogger{Logger: logger}

	logger.Infof("connecting to database on %s:%d", config.Host, config.Port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
	})

}
