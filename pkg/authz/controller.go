package authz

import (
	"context"
	"fmt"
	"spoke7-go/pkg/logger"
	"strings"

	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ControllerInterceptorABAC(SkipUrls []string, enforcer casbin.IEnforcer, logger logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Warnf("Unauthenticated request: missing metadata")
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		httpPath := ""
		if path := md.Get("path"); len(path) > 0 {
			httpPath = path[0]
		}

		if isSkipUrls(httpPath, SkipUrls) {
			logger.Debugf("Skipping ABAC for path: %s", httpPath)
			return handler(ctx, req)
		}

		user, err := GetUserFromContext(ctx)
		if err != nil {
			logger.Warnf("Unauthenticated request: user context not found for path %s", httpPath)
			return nil, status.Error(codes.Unauthenticated, "missing user context")
		}

		service := info.Server.(AuthResolver)
		object, action, err := service.GetObjectAndActionFromRequest(ctx, req, info)
		if err != nil {
			logger.Errorf("Failed to extract object and action for user %s: %v", user.Username, err)
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		if object == nil {
			logger.Infof("Authorization not requested: user=%+v", user)
			return handler(ctx, req)
		}

		logger.Infof("Authorization request: user=%+v, object=%+v, action=%s", user, object, action)

		allowed, roles, err := enforcer.EnforceEx(user, object, action)
		if err != nil {
			logger.Errorf("Error during ABAC enforcement for user %s: %v", user.Username, err)
			return nil, status.Error(codes.Internal, "authorization check failed")
		}

		if allowed {
			logger.Infof("Access granted: user=%s, action=%s, object=%v, roles=%v", user.Username, action, object, roles)
			logger.Debugf("User groups: %v", user.Groups)
			return handler(ctx, req)
		}

		logger.Warnf("Access denied: user=%s attempted %s on %v", user.Username, action, object)
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("permission denied for %s on %s", action, object.Type))
	}
}

func GetServiceAndMethodFromInfo(fullMethod string) (string, string, error) {
	parts := strings.Split(fullMethod, "/")
	if len(parts) < 3 {
		return "", "", status.Errorf(codes.Internal, "invalid method name: %s", fullMethod)
	}

	return parts[1], parts[2], nil
}
