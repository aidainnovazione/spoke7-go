package data

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"spoke7-go/internal/data/config"
	"spoke7-go/internal/data/controllers"
	"spoke7-go/internal/data/pb"
	"spoke7-go/internal/data/repository"
	"spoke7-go/internal/data/services"

	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/grpc_client"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"
	"strings"
	"sync"

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

	"spoke7-go/pkg/request_id"
)

var once sync.Once

type DataServer struct {
	version             string
	logger              logger.Logger
	e                   *echo.Echo
	repo                repository.DBClient
	grpcConn            *grpc.ClientConn
	grpcMetadataService grpc_client.GrpcMetadataClient
	grpcStorageService  grpc_client.GrpcStorageClient

	fileUploadService services.UploadFileService

	// services
	realTimeTrafficDataByDetectionPointByLaneService services.RealTimeTrafficDataByDetectionPointByLaneService
	realTimeTrafficDataByDetectionSectionService     services.RealTimeTrafficDataByDetectionSectionService

	currentTrafficDataByDetectionPointService       services.CurrentTrafficDataByDetectionPointService
	currentTrafficDataByDetectionPointByLaneService services.CurrentTrafficDataByDetectionPointByLaneService
	currentTrafficDataByDetectionSectionService     services.CurrentTrafficDataByDetectionSectionService

	historyDayTrafficDataByDetectionPointService       services.HistoryDayTrafficDataByDetectionPointService
	historyDayTrafficDataByDetectionPointByLaneService services.HistoryDayTrafficDataByDetectionPointByLaneService
	historyDayTrafficDataByDetectionSectionService     services.HistoryDayTrafficDataByDetectionSectionService

	historyHourTrafficDataByDetectionPointService       services.HistoryHourTrafficDataByDetectionPointService
	historyHourTrafficDataByDetectionPointByLaneService services.HistoryHourTrafficDataByDetectionPointByLaneService
	historyHourTrafficDataByDetectionSectionService     services.HistoryHourTrafficDataByDetectionSectionService

	// controllers
	realTimeTrafficDataByDetectionPointByLaneController pb.RealTimeTrafficDataByDetectionPointByLaneServiceServer
	realTimeTrafficDataByDetectionSectionController     pb.RealTimeTrafficDataByDetectionSectionServiceServer

	currentTrafficDataByDetectionPointController       pb.CurrentTrafficDataByDetectionPointServiceServer
	currentTrafficDataByDetectionPointByLaneController pb.CurrentTrafficDataByDetectionPointByLaneServiceServer
	currentTrafficDataByDetectionSectionController     pb.CurrentTrafficDataByDetectionSectionServiceServer

	historyDayTrafficDataByDetectionPointController       pb.HistoryDayTrafficDataByDetectionPointServiceServer
	historyDayTrafficDataByDetectionPointByLaneController pb.HistoryDayTrafficDataByDetectionPointByLaneServiceServer
	historyDayTrafficDataByDetectionSectionController     pb.HistoryDayTrafficDataByDetectionSectionServiceServer

	historyHourTrafficDataByDetectionPointController       pb.HistoryHourTrafficDataByDetectionPointServiceServer
	historyHourTrafficDataByDetectionPointByLaneController pb.HistoryHourTrafficDataByDetectionPointByLaneServiceServer
	historyHourTrafficDataByDetectionSectionController     pb.HistoryHourTrafficDataByDetectionSectionServiceServer
}

func NewDataServer(e *echo.Echo, repo repository.DBClient, grpcMetadataString string, grpcStorageString string, version string, logger logger.Logger, organizationName string) *DataServer {
	// Grpc services
	grpcMetadataService, err := grpc_client.NewGrpcMetadataClient(grpcMetadataString)
	if err != nil {
		panic(fmt.Errorf("error while connecting to data microservice"))
	}

	grpcStorageService, err := grpc_client.NewGrpcStorageClient(grpcStorageString)
	if err != nil {
		panic(fmt.Errorf("error while connecting to data microservice"))
	}

	// Initialize services

	realTimeTrafficDataByDetectionPointByLaneService := services.NewRealTimeTrafficDataByDetectionPointByLaneService(repo, *grpcMetadataService, organizationName)
	realTimeTrafficDataByDetectionSectionService := services.NewRealTimeTrafficDataByDetectionSectionService(repo, *grpcMetadataService, organizationName)

	currentTrafficDataByDetectionPointService := services.NewCurrentTrafficDataByDetectionPointService(repo, *grpcMetadataService, organizationName)
	currentTrafficDataByDetectionPointByLaneService := services.NewCurrentTrafficDataByDetectionPointByLaneService(repo, *grpcMetadataService, organizationName)
	currentTrafficDataByDetectionSectionService := services.NewCurrentTrafficDataByDetectionSectionService(repo, *grpcMetadataService, organizationName)

	historyDayTrafficDataByDetectionPointService := services.NewHistoryDayTrafficDataByDetectionPointService(repo, *grpcMetadataService, organizationName)
	historyDayTrafficDataByDetectionPointByLaneService := services.NewHistoryDayTrafficDataByDetectionPointByLaneService(repo, *grpcMetadataService, organizationName)
	historyDayTrafficDataByDetectionSectionService := services.NewHistoryDayTrafficDataByDetectionSectionService(repo, *grpcMetadataService, organizationName)

	historyHourTrafficDataByDetectionPointService := services.NewHistoryHourTrafficDataByDetectionPointService(repo, *grpcMetadataService, organizationName)
	historyHourTrafficDataByDetectionPointByLaneService := services.NewHistoryHourTrafficDataByDetectionPointByLaneService(repo, *grpcMetadataService, organizationName)
	historyHourTrafficDataByDetectionSectionService := services.NewHistoryHourTrafficDataByDetectionSectionService(repo, *grpcMetadataService, organizationName)

	fileUploadService := services.NewUploadFileService(*grpcStorageService)

	// Initialize controllers

	realTimeTrafficDataByDetectionPointByLaneController := controllers.NewRealTimeTrafficDataByDetectionPointByLaneController(realTimeTrafficDataByDetectionPointByLaneService, fileUploadService, logger)
	realTimeTrafficDataByDetectionSectionController := controllers.NewRealTimeTrafficDataByDetectionSectionController(realTimeTrafficDataByDetectionSectionService, fileUploadService, logger)

	currentTrafficDataByDetectionPointController := controllers.NewCurrentTrafficDataByDetectionPointController(currentTrafficDataByDetectionPointService, fileUploadService, logger)
	currentTrafficDataByDetectionPointByLaneController := controllers.NewCurrentTrafficDataByDetectionPointByLaneController(currentTrafficDataByDetectionPointByLaneService, fileUploadService, logger)
	currentTrafficDataByDetectionSectionController := controllers.NewCurrentTrafficDataByDetectionSectionController(currentTrafficDataByDetectionSectionService, fileUploadService, logger)

	historyDayTrafficDataByDetectionPointController := controllers.NewHistoryDayTrafficDataByDetectionPointController(historyDayTrafficDataByDetectionPointService, fileUploadService, logger)
	historyDayTrafficDataByDetectionPointByLaneController := controllers.NewHistoryDayTrafficDataByDetectionPointByLaneController(historyDayTrafficDataByDetectionPointByLaneService, fileUploadService, logger)
	historyDayTrafficDataByDetectionSectionController := controllers.NewHistoryDayTrafficDataByDetectionSectionController(historyDayTrafficDataByDetectionSectionService, fileUploadService, logger)

	historyHourTrafficDataByDetectionPointController := controllers.NewHistoryHourTrafficDataByDetectionPointController(historyHourTrafficDataByDetectionPointService, fileUploadService, logger)
	historyHourTrafficDataByDetectionPointByLaneController := controllers.NewHistoryHourTrafficDataByDetectionPointByLaneController(historyHourTrafficDataByDetectionPointByLaneService, fileUploadService, logger)
	historyHourTrafficDataByDetectionSectionController := controllers.NewHistoryHourTrafficDataByDetectionSectionController(historyHourTrafficDataByDetectionSectionService, fileUploadService, logger)

	return &DataServer{
		e:       e,
		repo:    repo,
		version: version,
		logger:  logger,

		grpcMetadataService: *grpcMetadataService,
		grpcStorageService:  *grpcStorageService,

		fileUploadService: fileUploadService,

		realTimeTrafficDataByDetectionPointByLaneService: realTimeTrafficDataByDetectionPointByLaneService,
		realTimeTrafficDataByDetectionSectionService:     realTimeTrafficDataByDetectionSectionService,

		currentTrafficDataByDetectionPointService:       currentTrafficDataByDetectionPointService,
		currentTrafficDataByDetectionPointByLaneService: currentTrafficDataByDetectionPointByLaneService,
		currentTrafficDataByDetectionSectionService:     currentTrafficDataByDetectionSectionService,

		historyDayTrafficDataByDetectionPointService:       historyDayTrafficDataByDetectionPointService,
		historyDayTrafficDataByDetectionPointByLaneService: historyDayTrafficDataByDetectionPointByLaneService,
		historyDayTrafficDataByDetectionSectionService:     historyDayTrafficDataByDetectionSectionService,

		historyHourTrafficDataByDetectionPointService:       historyHourTrafficDataByDetectionPointService,
		historyHourTrafficDataByDetectionPointByLaneService: historyHourTrafficDataByDetectionPointByLaneService,
		historyHourTrafficDataByDetectionSectionService:     historyHourTrafficDataByDetectionSectionService,

		realTimeTrafficDataByDetectionSectionController:     realTimeTrafficDataByDetectionSectionController,
		realTimeTrafficDataByDetectionPointByLaneController: realTimeTrafficDataByDetectionPointByLaneController,

		currentTrafficDataByDetectionPointController:       currentTrafficDataByDetectionPointController,
		currentTrafficDataByDetectionPointByLaneController: currentTrafficDataByDetectionPointByLaneController,
		currentTrafficDataByDetectionSectionController:     currentTrafficDataByDetectionSectionController,

		historyDayTrafficDataByDetectionPointController:       historyDayTrafficDataByDetectionPointController,
		historyDayTrafficDataByDetectionPointByLaneController: historyDayTrafficDataByDetectionPointByLaneController,
		historyDayTrafficDataByDetectionSectionController:     historyDayTrafficDataByDetectionSectionController,

		historyHourTrafficDataByDetectionPointController:       historyHourTrafficDataByDetectionPointController,
		historyHourTrafficDataByDetectionPointByLaneController: historyHourTrafficDataByDetectionPointByLaneController,
		historyHourTrafficDataByDetectionSectionController:     historyHourTrafficDataByDetectionSectionController,
	}
}

func (s *DataServer) Start() error {

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

func (s *DataServer) startGrpc() error {
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
			pb.RegisterRealTimeTrafficDataByDetectionPointByLaneServiceServer(grpcServer, s.realTimeTrafficDataByDetectionPointByLaneController)
			pb.RegisterRealTimeTrafficDataByDetectionSectionServiceServer(grpcServer, s.realTimeTrafficDataByDetectionSectionController)

			pb.RegisterCurrentTrafficDataByDetectionPointByLaneServiceServer(grpcServer, s.currentTrafficDataByDetectionPointByLaneController)
			pb.RegisterCurrentTrafficDataByDetectionPointServiceServer(grpcServer, s.currentTrafficDataByDetectionPointController)
			pb.RegisterCurrentTrafficDataByDetectionSectionServiceServer(grpcServer, s.currentTrafficDataByDetectionSectionController)

			pb.RegisterHistoryDayTrafficDataByDetectionPointByLaneServiceServer(grpcServer, s.historyDayTrafficDataByDetectionPointByLaneController)
			pb.RegisterHistoryDayTrafficDataByDetectionPointServiceServer(grpcServer, s.historyDayTrafficDataByDetectionPointController)
			pb.RegisterHistoryDayTrafficDataByDetectionSectionServiceServer(grpcServer, s.historyDayTrafficDataByDetectionSectionController)

			pb.RegisterHistoryHourTrafficDataByDetectionPointByLaneServiceServer(grpcServer, s.historyHourTrafficDataByDetectionPointByLaneController)
			pb.RegisterHistoryHourTrafficDataByDetectionPointServiceServer(grpcServer, s.historyHourTrafficDataByDetectionPointController)
			pb.RegisterHistoryHourTrafficDataByDetectionSectionServiceServer(grpcServer, s.historyHourTrafficDataByDetectionSectionController)

			// Serve gRPC server
			grpcServer.Serve(lis)
		}()
	})

	return nil
}

// connectToGrpcServer connects to the gRPC server.
func (s *DataServer) connectToGrpcServer(ctx context.Context) error {

	url := fmt.Sprintf("%s:%d", config.AppConfig.Service.Host, config.AppConfig.Service.GrpcPort)
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		return err
	}
	s.grpcConn = conn

	return nil
}

func (s *DataServer) loadMiddlewares() {
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

func (s *DataServer) LoadRoutes() error {

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
			if file, ok := resp.(*pb.DownloadResponse); ok && len(file.FileContent) > 0 {
				w.Header().Set("Content-Type", file.ContentType)
				contentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", file.Filename)
				w.Header().Set("Content-Disposition", contentDisposition)
				w.Header().Set("Content-Length", fmt.Sprintf("%d", len(file.FileContent)))
				_, err := w.Write(file.FileContent)
				return err
			}
			return nil
		}),
	)

	err := pb.RegisterCurrentTrafficDataByDetectionPointByLaneServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterCurrentTrafficDataByDetectionPointServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterCurrentTrafficDataByDetectionSectionServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	err = pb.RegisterRealTimeTrafficDataByDetectionSectionServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterRealTimeTrafficDataByDetectionPointByLaneServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	err = pb.RegisterHistoryDayTrafficDataByDetectionPointByLaneServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterHistoryDayTrafficDataByDetectionPointServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterHistoryDayTrafficDataByDetectionSectionServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	err = pb.RegisterHistoryHourTrafficDataByDetectionPointByLaneServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterHistoryHourTrafficDataByDetectionPointServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}
	err = pb.RegisterHistoryHourTrafficDataByDetectionSectionServiceHandler(context.Background(), gwmux, s.grpcConn)
	if err != nil {
		return err
	}

	s.e.Any("/*", echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gwmux.ServeHTTP(w, r)
	})))

	return err
}
