package services

import (
	"context"
	"errors"

	"github.com/flydevs/chat-app-api/common/logger"
	oauth_domain "github.com/flydevs/chat-app-api/messaging-api/src/domain/oauth"
	oauth_repo "github.com/flydevs/chat-app-api/messaging-api/src/repository/oauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	accessibleRoles map[string]int
	oauthRepo       oauth_repo.OauthRepositoryInterface
}

func NewAuthInterceptor(accroles map[string]int, oauthRepo oauth_repo.OauthRepositoryInterface) *AuthInterceptor {
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
	m := make(oauth_domain.OauthDataMap)

	inc_md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		m.SetStatus("3")
		logger.Error("error in function FromIncomingContext inside AuthInterceptor", errors.New("ok false"))
		return makeNewContext(ctx, m), nil
	}

	values := inc_md.Get("access-token")
	if len(values) == 0 {
		m.SetStatus("1")
		return makeNewContext(ctx, m), nil
	}
	data, err := interceptor.oauthRepo.Verify(values[0])
	if err != nil {
		m.SetStatus("3")
		logger.Error("error with oauthRepo Verify function inside AuthInterceptor", err)
		return makeNewContext(ctx, m), nil
	}
	switch {
	case data.Response.Status == 401:
		m.SetStatus("1")
		return makeNewContext(ctx, m), nil
	case data.Response.Status == 500:
		m.SetStatus("3")
		logger.Error("error with oauthRepo Verify function inside AuthInterceptor", err)
		return makeNewContext(ctx, m), nil
	case int32(permission) > data.Permissions:
		m.SetStatus("2")
		return makeNewContext(ctx, m), nil
	}
	m.SetStatus("0")
	m.SetUuid(data.Uuid.Uuid)
	return makeNewContext(ctx, m), nil
}

func makeNewContext(ctx context.Context, mdata oauth_domain.OauthDataMap) context.Context {
	new_md := metadata.New(mdata)
	return metadata.NewIncomingContext(context.Background(), new_md)

}
