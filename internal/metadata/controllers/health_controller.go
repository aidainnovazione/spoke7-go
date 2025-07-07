package controllers

import (
	"context"
	"spoke7-go/internal/metadata/pb"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"

	"google.golang.org/protobuf/types/known/emptypb"
)

var HEALTH_MICROSERVICE string = "metadata"
var HEALTH_RESOURCE string = "health"

type healthController struct {
	version string
	logger  logger.Logger
	pb.UnimplementedHealthServiceServer
}

func NewHealthController(version string, logger logger.Logger) pb.HealthServiceServer {
	return &healthController{version: version, logger: logger}
}

func (h *healthController) GetHealthStatus(ctx context.Context, _ *emptypb.Empty) (*pb.Health, error) {
	logText := "Request Health"
	_ = utils.AddControllerRequestLogging(h.logger, ctx, logText, HEALTH_MICROSERVICE, HEALTH_RESOURCE, "GET", "NONE")

	return &pb.Health{Status: "OK", Version: h.version}, nil
}
