package metadata

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"spoke7-go/internal/metadata/config"
	"spoke7-go/internal/metadata/controllers"
	"spoke7-go/internal/metadata/pb"
	"spoke7-go/internal/metadata/repository"
	"spoke7-go/internal/metadata/services"
	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/grpc_client"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"
	"strings"
	"sync"

	"spoke7-go/pkg/request_id"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

type MetadataServer struct {
	version            string
	logger             logger.Logger
	e                  *echo.Echo
	repo               repository.DBClient
	grpcConn           *grpc.ClientConn
	grpcStorageService grpc_client.GrpcStorageClient

	fileUploadService services.UploadFileService

	// services
	dataSourceService       services.DataSourceService
	detectionSectionService services.DetectionSectionService
	detectionPointService   services.DetectionPointService
	roadNetworkService      services.RoadNetworkService

	// controllers
	dataSourceController       pb.DataSourceServiceServer
	detectionSectionController pb.DetectionSectionServiceServer
	detectionPointController   pb.DetectionPointServiceServer
	roadNetworkController      pb.RoadNetworkServiceServer
	healthController           pb.HealthServiceServer
	dashboardController        pb.DashboardServiceServer
}

func NewMetadataServer(e *echo.Echo, repo repository.DBClient, grpcStorageString string, version string, logger logger.Logger, organizationName string) *MetadataServer {
	grpcStorageService, err := grpc_client.NewGrpcStorageClient(grpcStorageString)
	if err != nil {
		panic(fmt.Errorf("error while connecting to data microservice"))
	}
	fileUploadService := services.NewUploadFileService(*grpcStorageService)

	// Initialize services
	dataSourceService := services.NewDataSourceService(repo, organizationName)
	detectionSectionService := services.NewDetectionSectionService(repo)
	detectionPointService := services.NewDetectionPointService(repo)
	roadNetworkService := services.NewRoadNetworkService(repo, logger)
	dashboardService := services.NewDashboardService(repo, organizationName)

	// Initialize controllers
	dataSourceController := controllers.NewDataSourceController(dataSourceService, logger)
	detectionSectionController := controllers.NewDetectionSectionController(detectionSectionService, fileUploadService, logger)
	detectionPointController := controllers.NewDetectionPointController(detectionPointService, fileUploadService, logger)
	roadNetworkController := controllers.NewRoadNetworkController(roadNetworkService, logger)
	healthController := controllers.NewHealthController(version, logger)
	dashboardController := controllers.NewDashboardController(dashboardService, logger)

	return &MetadataServer{
		e:       e,
		repo:    repo,
		version: version,
		logger:  logger,

		dataSourceService:       dataSourceService,
		detectionSectionService: detectionSectionService,
		detectionPointService:   detectionPointService,
		roadNetworkService:      roadNetworkService,

		dataSourceController:       dataSourceController,
		detectionSectionController: detectionSectionController,
		detectionPointController:   detectionPointController,
		roadNetworkController:      roadNetworkController,
		healthController:           healthController,
		dashboardController:        dashboardController,
	}
}

func (s *MetadataServer) Start() error {

	err := s.startGrpc()
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

func (s *MetadataServer) startGrpc() error {
	once.Do(func() {
		grpcPort := config.AppConfig.Service.GrpcPort
		address := fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, grpcPort)
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
				grpc.ChainUnaryInterceptor(
					request_id.RequestIDUnaryServerInterceptor(),
					authz.JWTAuthInterceptor(config.AppConfig.Authentication.SkipUrls, jwtAuthConfig, s.logger),
				),

				//jwtInterceptor(config.AppConfig.Authentication.SkipUrls)),
			)

			// Register reflection service on gRPC server.
			reflection.Register(grpcServer)

			// register grpc services
			pb.RegisterDataSourceServiceServer(grpcServer, s.dataSourceController)
			pb.RegisterDetectionSectionServiceServer(grpcServer, s.detectionSectionController)
			pb.RegisterDetectionPointServiceServer(grpcServer, s.detectionPointController)
			pb.RegisterRoadNetworkServiceServer(grpcServer, s.roadNetworkController)
			pb.RegisterHealthServiceServer(grpcServer, s.healthController)
			pb.RegisterDashboardServiceServer(grpcServer, s.dashboardController)

			// Serve gRPC server
			grpcServer.Serve(lis)
		}()
	})

	return nil
}

// func jwtInterceptor(urls []string) grpc.UnaryServerInterceptor {
// 	fmt.Printf("jwt Intercepct")
// 	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

// 		fmt.Printf("extract metadata")

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
// 		switch v := claims[config.AppConfig.Authentication.RolesClaim].(type) {
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

// 		// Call the handler if authentication passes
// 		return handler(ctx, req)
// 	}
// }

// connectToGrpcServer connects to the gRPC server.
func (s *MetadataServer) connectToGrpcServer(ctx context.Context) error {

	url := fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, config.AppConfig.Service.GrpcPort)
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		return err
	}
	s.grpcConn = conn

	return nil
}

func (s *MetadataServer) loadMiddlewares() {
	s.e.HideBanner = true
	s.e.Use(
		middleware.CORS(),
		middleware.Recover(),
		request_id.RequestIDMiddleware,
	)
}

func isSkipUrls(path string, urls []string) bool {
	for _, url := range urls {
		if path == url || path == "" || strings.HasPrefix(path, url) {
			return true
		}
	}
	return false

}

func (s *MetadataServer) LoadRoutes() error {

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

	err := pb.RegisterDataSourceServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterDetectionSectionServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterDetectionPointServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterHealthServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	err = pb.RegisterRoadNetworkServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	err = pb.RegisterDashboardServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	s.e.Any("/*", echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gwmux.ServeHTTP(w, r)
	})))

	return err
}
