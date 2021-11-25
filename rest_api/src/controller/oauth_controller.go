package controller

import (
	"io/ioutil"

	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/oauth"
	"github.com/flydevs/chat-app-api/rest-api/src/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type oauthController struct {
	oauthsvs services.OauthServiceInterface
}

type OauthControllerInterface interface {
	LoginUser(*gin.Context)
}

func NewOauthController(oauthsvs services.OauthServiceInterface) OauthControllerInterface {
	return &oauthController{oauthsvs: oauthsvs}
}

func (oauthctrl oauthController) LoginUser(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := oauth.LoginRequest{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	result_response_object, header := oauthctrl.oauthsvs.LoginUser(&new_request)
	if header != nil {
		c.Header("access-token", *header)
	}
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}
