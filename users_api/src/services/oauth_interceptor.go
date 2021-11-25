package services

import (
	"context"

	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/users-api/src/repository/oauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	accessibleRoles map[string]int
	oauthRepo       oauth.OauthRepositoryInterface
}

func NewAuthInterceptor(accroles map[string]int, oauthRepo oauth.OauthRepositoryInterface) *AuthInterceptor {
	return &AuthInterceptor{accessibleRoles: accroles, oauthRepo: oauthRepo}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.Info("----> unary interceptor: " + info.FullMethod)
		rctx, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(rctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	permission, ok := interceptor.accessibleRoles[method]
	if !ok {
		return ctx, nil
	}
	inc_md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		//This error handling is awful, change it after you update the oauth service
		values := inc_md.Get("access-token")
		if len(values) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "access_token not provided")
		}
		data, err := interceptor.oauthRepo.Verify(values[0])
		if err != nil {
			return nil, err
		}
		switch data.Response.Status {
		case 401:
			return nil, status.Errorf(codes.Unauthenticated, "verification invalid")
		case 500:
			return nil, status.Errorf(codes.Internal, "Oops... Something whent wrong")

		}

		if int32(permission) > data.Permissions {
			return nil, status.Errorf(codes.PermissionDenied, "not enough permissions for this request")

		}

		m := map[string]string{"uuid": data.Uuid.Uuid}
		new_md := metadata.New(m)
		new_ctx := metadata.NewIncomingContext(ctx, new_md)
		return new_ctx, nil
	}
	return ctx, nil
}
