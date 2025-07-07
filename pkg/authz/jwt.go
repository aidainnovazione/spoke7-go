package authz

import (
	"context"
	"spoke7-go/pkg/logger"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JWTAuthInterceptorConfig struct {
	JwksUrl          string
	UsernameClaim    string
	RolesClaim       string
	GroupsClaim      string
	OrganizationName string
}

func JWTAuthInterceptor(SkipUrls []string, config JWTAuthInterceptorConfig, logger logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}
		httpPath := ""
		if path := md.Get("path"); len(path) > 0 {
			httpPath = path[0]
		}
		if isSkipUrls(httpPath, SkipUrls) {
			return handler(ctx, req)
		}

		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
		if tokenStr == authHeader[0] {
			return nil, status.Error(codes.Unauthenticated, "malformed authorization header")
		}

		k, err := keyfunc.NewDefaultCtx(context.Background(), []string{config.JwksUrl})
		if err != nil {
			return nil, err
		}

		token, err := jwt.Parse(tokenStr, k.Keyfunc)
		if err != nil || !token.Valid {
			return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, status.Error(codes.Internal, "invalid claims structure")
		}

		var roles []string
		switch v := claims[config.RolesClaim].(type) {
		case string:
			roles = strings.Split(v, ",")
		case []interface{}:
			for _, r := range v {
				if roleStr, ok := r.(string); ok {
					roles = append(roles, roleStr)
				}
			}
		case []string:
			roles = v
		default:
			return nil, status.Error(codes.Internal, "unexpected role claim type")
		}

		var groups []string
		switch v := claims[config.GroupsClaim].(type) {
		case string:
			groups = strings.Split(v, ",")
		case []interface{}:
			for _, r := range v {
				if groupStr, ok := r.(string); ok {
					groupStr := strings.TrimPrefix(groupStr, config.OrganizationName+"/")
					groups = append(groups, groupStr)
				}
			}
		case []string:
			groups = v
		default:
			return nil, status.Error(codes.Internal, "unexpected role claim type")
		}

		user := User{
			Username: claims[config.UsernameClaim].(string),
			Groups:   groups,
			Roles:    roles,
		}

		ctx = context.WithValue(ctx, UserCtxKey, user)
		ctx = context.WithValue(ctx, "authorization", authHeader)
		logger.Debugf("Authenticated user: %s, roles: %v, groups %v", user.Username, roles, groups)

		return handler(ctx, req)
	}
}

func isSkipUrls(path string, urls []string) bool {
	for _, url := range urls {
		if path == url || path == "" || strings.HasPrefix(path, url) {
			return true
		}
	}
	return false

}
