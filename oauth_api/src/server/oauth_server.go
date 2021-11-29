package server

import (
	"context"
	"net/http"

	"github.com/Sora8d/common/server_message"
	proto_oauth "github.com/flydevs/chat-app-api/oauth-api/src/clients/rpc/oauth"
	"github.com/flydevs/chat-app-api/oauth-api/src/controller"
)

type OauthServer struct {
	proto_oauth.UnimplementedOauthProtoInterfaceServer
	oauthctrl controller.OauthControllerInterface
}

func GetNewServer(oauthctrl controller.OauthControllerInterface) *OauthServer {
	return &OauthServer{oauthctrl: oauthctrl}
}

func (oauthsvr OauthServer) LoginUser(ctx context.Context, request *proto_oauth.LoginRequest) (*proto_oauth.JWTwRrefreshUuidResponse, error) {
	var proto_response *proto_oauth.JWTwRrefreshUuidResponse
	response, aErr := oauthsvr.oauthctrl.GenerateUser(ctx, request)
	if aErr != nil {
		response_with_error := proto_oauth.JWTwRrefreshUuidResponse{}
		response_with_error.Response = poblateMessage(aErr)
		proto_response = &response_with_error
	} else {
		proto_response = response
		proto_response.Response = poblateMessage(OKMessage("user logged succesfully"))
	}
	return proto_response, nil
}
func (oauthsvr OauthServer) LoginClient(ctx context.Context, request *proto_oauth.ServiceKey) (*proto_oauth.JWTResponse, error) {
	var proto_response *proto_oauth.JWTResponse
	response, aErr := oauthsvr.oauthctrl.GenerateService(ctx, request)
	if aErr != nil {
		response_with_error := proto_oauth.JWTResponse{}
		response_with_error.Response = poblateMessage(aErr)
		proto_response = &response_with_error
	} else {
		proto_response = response
		proto_response.Response = poblateMessage(OKMessage("client logged succesfully"))
	}
	return proto_response, nil
}

func (oauthsvr OauthServer) Verify(ctx context.Context, request *proto_oauth.JWT) (*proto_oauth.EntityResponse, error) {
	var proto_response *proto_oauth.EntityResponse
	response, aErr := oauthsvr.oauthctrl.Verify(ctx, request)
	if aErr != nil {
		response_with_error := proto_oauth.EntityResponse{}
		response_with_error.Response = poblateMessage(aErr)
		proto_response = &response_with_error
	} else {
		proto_response = response
		proto_response.Response = poblateMessage(OKMessage("entity verified succesfully"))
	}
	return proto_response, nil
}

func (oauthsvr OauthServer) ValidateRefreshToken(ctx context.Context, request *proto_oauth.JWT) (*proto_oauth.JWTwRrefreshUuidResponse, error) {
	var proto_response *proto_oauth.JWTwRrefreshUuidResponse
	response, aErr := oauthsvr.oauthctrl.ValidateRefreshToken(ctx, request)
	if aErr != nil {
		response_with_error := proto_oauth.JWTwRrefreshUuidResponse{}
		response_with_error.Response = poblateMessage(aErr)
		proto_response = &response_with_error
	} else {
		proto_response = response
		proto_response.Response = poblateMessage(OKMessage("refresh token accepted succesfully"))
	}
	return proto_response, nil
}

func (oauthsvr OauthServer) RevokeUsersTokens(ctx context.Context, request *proto_oauth.Uuid) (*proto_oauth.SvrMsg, error) {
	var proto_response *proto_oauth.SvrMsg
	if aErr := oauthsvr.oauthctrl.RevokeUsersTokens(ctx, request); aErr != nil {
		proto_response = poblateMessage(aErr)
	} else {
		proto_response = poblateMessage(OKMessage("tokens blocked"))
	}
	return proto_response, nil
}

func poblateMessage(msg server_message.Svr_message) *proto_oauth.SvrMsg {
	pErr := proto_oauth.SvrMsg{}
	pErr.Status = int32(msg.GetStatus())
	pErr.Message = msg.GetMessage()
	return &pErr
}

func OKMessage(message string) server_message.Svr_message {
	return server_message.NewCustomMessage(http.StatusOK, message)
}
