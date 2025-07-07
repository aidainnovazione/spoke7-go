package request_id

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestIDUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if requestIDs := md.Get("x-request-id"); len(requestIDs) > 0 {
				ctx = context.WithValue(ctx, RequestIdCtxKey, requestIDs[0])
			}
		}
		return handler(ctx, req)
	}
}
