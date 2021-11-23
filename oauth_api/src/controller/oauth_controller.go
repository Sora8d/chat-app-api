package controller

import (
	"context"

	"github.com/flydevs/chat-app-api/common/server_message"
	proto_oauth "github.com/flydevs/chat-app-api/oauth-api/src/clients/rpc/oauth"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/client"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/users"
	"github.com/flydevs/chat-app-api/oauth-api/src/services"
)

type oauthController struct {
	oauthsvs services.OauthServiceInterface
}

type OauthControllerInterface interface {
	GenerateUser(context.Context, *proto_oauth.LoginRequest) (*proto_oauth.JWTAndUuidResponse, server_message.Svr_message)
	GenerateService(context.Context, *proto_oauth.ServiceKey) (*proto_oauth.JWTResponse, server_message.Svr_message)
	Verify(context.Context, *proto_oauth.JWT) (*proto_oauth.EntityResponse, server_message.Svr_message)
}

func GetNewOauthController(oauthsvs services.OauthServiceInterface) OauthControllerInterface {
	return &oauthController{oauthsvs: oauthsvs}
}

func (oauthctrl oauthController) GenerateUser(ctx context.Context, proto_request *proto_oauth.LoginRequest) (*proto_oauth.JWTAndUuidResponse, server_message.Svr_message) {
	var request users.LoginRequest
	request.Poblate(false, proto_request)
	uuid, jwt, aErr := oauthctrl.oauthsvs.GenerateUser(request)
	if aErr != nil {
		return nil, aErr
	}
	proto_uuid := proto_oauth.Uuid{Uuid: *uuid}
	proto_result := proto_oauth.JWTAndUuidResponse{Uuid: &proto_uuid, Jwt: *jwt}
	return &proto_result, nil
}
func (oauthctrl oauthController) GenerateService(ctx context.Context, proto_request *proto_oauth.ServiceKey) (*proto_oauth.JWTResponse, server_message.Svr_message) {
	var request client.ServiceKey
	request.Set(proto_request.Key)
	result, aErr := oauthctrl.oauthsvs.GenerateService(request)
	if aErr != nil {
		return nil, aErr
	}
	proto_result := proto_oauth.JWTResponse{Jwt: *result}
	return &proto_result, nil
}
func (oauthctrl oauthController) Verify(ctx context.Context, proto_request *proto_oauth.JWT) (*proto_oauth.EntityResponse, server_message.Svr_message) {
	result, aErr := oauthctrl.oauthsvs.Verify(proto_request.GetJwt())
	if aErr != nil {
		return nil, aErr
	}
	proto_uuid := proto_oauth.Uuid{Uuid: result.Uuid}
	proto_result := proto_oauth.EntityResponse{Uuid: &proto_uuid, Permissions: int32(result.Permissions)}
	return &proto_result, nil
}
