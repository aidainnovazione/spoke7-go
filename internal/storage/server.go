package storage

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"spoke7-go/internal/storage/config"
	"spoke7-go/internal/storage/controllers"
	"spoke7-go/internal/storage/pb"
	"spoke7-go/internal/storage/services"
	"spoke7-go/internal/storage/storage_interface"
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
	"google.golang.org/protobuf/proto"
)

var once sync.Once

type StorageServer struct {
	version  string
	logger   logger.Logger
	e        *echo.Echo
	repo     storage_interface.StorageInterface
	grpcConn *grpc.ClientConn

	grpcMetadataService grpc_client.GrpcMetadataClient
	// services
	storedFileService services.StoredFileService
	// controllers
	storedFileController pb.StoredFileServiceServer
}

func NewStorageServer(e *echo.Echo, repo storage_interface.StorageInterface, grpcMetadataString string, version string, logger logger.Logger, organizationName string) *StorageServer {
	grpcMetadataService, err := grpc_client.NewGrpcMetadataClient(grpcMetadataString)
	if err != nil {
		panic(fmt.Errorf("error while connecting to data microservice"))
	}

	// Initialize services
	storedFileService := services.NewStoredFileService(repo, *grpcMetadataService, organizationName)

	// Initialize controllers
	storedFileController := controllers.NewStoredFileController(storedFileService)

	return &StorageServer{
		e:                   e,
		repo:                repo,
		version:             version,
		logger:              logger,
		grpcMetadataService: *grpcMetadataService,

		storedFileService: storedFileService,

		storedFileController: storedFileController,
	}
}

func (s *StorageServer) Start() error {

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

func (s *StorageServer) startGrpc() error {
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
				grpc.ChainUnaryInterceptor(authz.JWTAuthInterceptor(config.AppConfig.Authentication.SkipUrls, jwtAuthConfig, s.logger), request_id.RequestIDUnaryServerInterceptor()),
			)

			// Register reflection service on gRPC server.
			reflection.Register(grpcServer)

			// register grpc services
			pb.RegisterStoredFileServiceServer(grpcServer, s.storedFileController)

			// Serve gRPC server
			grpcServer.Serve(lis)
		}()
	})

	return nil
}

// connectToGrpcServer connects to the gRPC server.
func (s *StorageServer) connectToGrpcServer(ctx context.Context) error {

	url := fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, config.AppConfig.Service.GrpcPort)
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		return err
	}
	s.grpcConn = conn

	return nil
}

func (s *StorageServer) loadMiddlewares() {
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

func (s *StorageServer) LoadRoutes() error {

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
		runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
			if file, ok := resp.(*pb.StoredFileDownloadResponse); ok && len(file.Content) > 0 {
				w.Header().Set("Content-Type", file.FileFormat)
				contentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", file.FileName)
				w.Header().Set("Content-Disposition", contentDisposition)
				w.Header().Set("Content-Length", fmt.Sprintf("%d", len(file.Content)))
				_, err := w.Write(file.Content)
				return err
			}
			return nil
		}),
	)

	err := pb.RegisterStoredFileServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	s.e.Any("/*", echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gwmux.ServeHTTP(w, r)
	})))

	return err
}
