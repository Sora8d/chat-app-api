package services

import (
	"context"

	"github.com/Sora8d/common/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(accroles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{accessibleRoles: accroles}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.Info("----> unary interceptor: " + info.FullMethod)
		rctx, _ := interceptor.authorize(ctx, info.FullMethod)
		return handler(rctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	_, ok := interceptor.accessibleRoles[method]
	if !ok {
		return ctx, nil
	}
	// This all would be different when i need to return an error whenever auth is not provided
	inc_md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := inc_md.Get("authorization")
		if len(values) == 0 {
			return ctx, nil
		}
		//TODO: add verify function
		m := map[string]string{"user_uuid": "0", "admin": "true"}
		new_md := metadata.New(m)
		new_ctx := metadata.NewIncomingContext(ctx, new_md)
		return new_ctx, nil
	}
	return ctx, nil
}
