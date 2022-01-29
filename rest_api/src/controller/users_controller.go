package controller

import (
	"context"
	"io/ioutil"

	"github.com/Sora8d/common/server_message"
	"github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/users"
	"github.com/flydevs/chat-app-api/rest-api/src/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type usersController struct {
	us_svs services.UsersServiceInterface
}

type UsersControllerInterface interface {
	CreateUser(*gin.Context)
	GetUserProfileByUuid(*gin.Context)
	UpdateUser(*gin.Context)
	SearchContact(*gin.Context)
}

func NewUsersController(svs services.UsersServiceInterface) UsersControllerInterface {
	return &usersController{us_svs: svs}
}

func (uctrl usersController) CreateUser(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := users.RegisterUser{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	result_response_object := uctrl.us_svs.CreateUser(context.Background(), &new_request)
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}
func (uctrl usersController) GetUserProfileByUuid(c *gin.Context) {
	uuids := c.QueryArray("uuid")
	new_request := users.MultipleUuids{}
	for _, uuid := range uuids {
		new_request.Uuids = append(new_request.Uuids, &users.Uuid{Uuid: uuid})
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := uctrl.us_svs.GetUserProfileByUuid(ctx, &new_request)
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}
func (uctrl usersController) UpdateUser(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := users.UpdateUserRequest{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := uctrl.us_svs.UpdateUser(ctx, &new_request)
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (uctrl usersController) SearchContact(c *gin.Context) {
	query := c.Query("query")
	excludes := c.QueryArray("ex_uuid")
	new_request := users.SearchContactRequest{}
	new_request.Query = query
	if len(excludes) > 0 {
		new_request.ExcludeUuids = excludes
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := uctrl.us_svs.SearchContact(ctx, &new_request)
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}
