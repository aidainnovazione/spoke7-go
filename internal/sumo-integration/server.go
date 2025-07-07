package data

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"spoke7-go/internal/sumo-integration/config"
	"spoke7-go/internal/sumo-integration/controllers"
	"spoke7-go/internal/sumo-integration/pb"
	"spoke7-go/internal/sumo-integration/services"
	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/grpc_client"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"
	"strings"
	"sync"

	"spoke7-go/pkg/request_id"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

var once sync.Once

type SumoIntegrationServer struct {
	version  string
	logger   logger.Logger
	e        *echo.Echo
	grpcConn *grpc.ClientConn

	//external client
	grpcDataService     grpc_client.GrpcDataClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	grpcStorageService  grpc_client.GrpcStorageClient

	// services
	currentTrafficDataByDetectionPointByLaneService services.SumoIntegrationCurrentTrafficDataByDetectionPointByLaneService
	currentTrafficDataByDetectionSectionService     services.SumoIntegrationCurrentTrafficDataByDetectionSectionService
	datasourceService                               services.DataSourceService
	fileUploadService                               services.UploadFileService

	// controllers
	currentTrafficDataByDetectionPointByLaneController pb.SumoIntegrationCurrentTrafficDataByDetectionPointByLaneServiceServer
	currentTrafficDataByDetectionSectionController     pb.SumoIntegrationCurrentTrafficDataByDetectionSectionServiceServer
	dataSourceController                               pb.SumoIntegrationDataSourceServiceServer
}

func NewSumoIntegrationServer(e *echo.Echo, grpcDataString string, grpcMetadataString string, grpcStorageString string, version string, logger logger.Logger) *SumoIntegrationServer {
	// Initialize grpc service
	grpcDataService, err := grpc_client.NewGrpcDataClient(grpcDataString)
	if err != nil {
		panic(fmt.Errorf("error while connecting to data microservice"))
	}

	grpcMetadataService, err := grpc_client.NewGrpcMetadataClient(grpcMetadataString)
	if err != nil {
		panic(fmt.Errorf("error while connecting to data microservice"))
	}

	grpcStorageService, err := grpc_client.NewGrpcStorageClient(grpcStorageString)
	if err != nil {
		panic(fmt.Errorf("error while connecting to data microservice"))
	}

	// Initialize services

	currentTrafficDataByDetectionPointByLaneService := services.NewSumoIntegrationCurrentTrafficDataByDetectionPointByLaneService(*grpcDataService, *grpcMetadataService)
	currentTrafficDataByDetectionSectionService := services.NewSumoIntegrationCurrentTrafficDataByDetectionSectionService(*grpcDataService, *grpcMetadataService)
	roadNetworkService := services.NewDataSourceService(*grpcDataService, *grpcMetadataService)
	fileUploadService := services.NewUploadFileService(*grpcStorageService)

	// Initialize controllers

	currentTrafficDataByDetectionPointByLaneController := controllers.NewSumoIntegrationCurrentTrafficDataByDetectionPointByLaneController(currentTrafficDataByDetectionPointByLaneService, fileUploadService, logger)
	currentTrafficDataByDetectionSectionController := controllers.NewSumoIntegrationCurrentTrafficDataByDetectionSectionController(currentTrafficDataByDetectionSectionService, fileUploadService, logger)
	dataSourceController := controllers.NewDataSourceController(roadNetworkService, fileUploadService, logger)

	return &SumoIntegrationServer{
		e:       e,
		version: version,
		logger:  logger,

		grpcDataService:     *grpcDataService,
		grpcMetadataService: *grpcMetadataService,
		grpcStorageService:  *grpcStorageService,

		currentTrafficDataByDetectionPointByLaneService: currentTrafficDataByDetectionPointByLaneService,
		currentTrafficDataByDetectionSectionService:     currentTrafficDataByDetectionSectionService,
		datasourceService: roadNetworkService,
		fileUploadService: fileUploadService,

		currentTrafficDataByDetectionPointByLaneController: currentTrafficDataByDetectionPointByLaneController,
		currentTrafficDataByDetectionSectionController:     currentTrafficDataByDetectionSectionController,
		dataSourceController:                               dataSourceController,
	}
}

func (s *SumoIntegrationServer) Start() error {

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

func (s *SumoIntegrationServer) startGrpc() error {
	once.Do(func() {
		grpcPort := config.AppConfig.Service.GrpcPort
		address := fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, grpcPort)
		s.logger.Infof("Starting gRPC server on %d", grpcPort)

		go func() {
			lis, err := net.Listen("tcp", address)
			if err != nil {
				return
			}
			grpcServer := grpc.NewServer(
				grpc.ChainUnaryInterceptor(jwtInterceptor(config.AppConfig.Authentication.SkipUrls), request_id.RequestIDUnaryServerInterceptor()),
			)

			// Register reflection service on gRPC server.
			reflection.Register(grpcServer)

			// register grpc services
			pb.RegisterSumoIntegrationCurrentTrafficDataByDetectionPointByLaneServiceServer(grpcServer, s.currentTrafficDataByDetectionPointByLaneController)
			pb.RegisterSumoIntegrationCurrentTrafficDataByDetectionSectionServiceServer(grpcServer, s.currentTrafficDataByDetectionSectionController)
			pb.RegisterSumoIntegrationDataSourceServiceServer(grpcServer, s.dataSourceController)

			// Serve gRPC server
			grpcServer.Serve(lis)
		}()
	})

	return nil
}

func jwtInterceptor(urls []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		// Extract metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata")
		}
		// Extract the HTTP request path from metadata
		httpPath := ""
		if pathValues := md.Get("path"); len(pathValues) > 0 {
			httpPath = pathValues[0]
		}
		if isSkipUrls(httpPath, urls) {
			return handler(ctx, req)
		}

		// Get token from authorization header
		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
		}

		// Extract token (format: Bearer <token>)
		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
		if tokenString == authHeader[0] {
			return nil, status.Errorf(codes.Unauthenticated, "malformed authorization header")
		}

		// Load JWKS
		k, err := keyfunc.NewDefaultCtx(context.Background(), []string{config.AppConfig.Authentication.JwksUrl})
		if err != nil {
			return nil, err
		}

		// Parse and validate the JWT
		token, err := jwt.Parse(tokenString, k.Keyfunc)
		if err != nil || !token.Valid {
			return nil, errors.New("invalid or expired token")
		}

		// get the claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("invalid claims")
		}

		var roles []string
		switch v := claims[config.AppConfig.Authentication.RolesClaim].(type) {
		case string:
			roles = strings.Split(v, ",")
		case []interface{}:
			for _, role := range v {
				roles = append(roles, role.(string))
			}
		case []string:
			roles = v
		default:
			return nil, errors.New("invalid roles claim in jwt")
		}

		user := authz.User{
			Username: claims[config.AppConfig.Authentication.UsernameClaim].(string),
			Groups:   roles,
		}
		ctx = context.WithValue(ctx, authz.UserCtxKey, user)
		ctx = context.WithValue(ctx, "authorization", authHeader)

		// Call the handler if authentication passes
		return handler(ctx, req)
	}
}

// connectToGrpcServer connects to the gRPC server.
func (s *SumoIntegrationServer) connectToGrpcServer(ctx context.Context) error {

	url := fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, config.AppConfig.Service.GrpcPort)
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		return err
	}
	s.grpcConn = conn

	return nil
}

func (s *SumoIntegrationServer) loadMiddlewares() {
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

func (s *SumoIntegrationServer) LoadRoutes() error {

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

	err := pb.RegisterSumoIntegrationCurrentTrafficDataByDetectionPointByLaneServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	err = pb.RegisterSumoIntegrationCurrentTrafficDataByDetectionSectionServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	err = pb.RegisterSumoIntegrationDataSourceServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	s.e.Any("/*", echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gwmux.ServeHTTP(w, r)
	})))

	return err
}
