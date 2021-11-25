package app

import (
	"github.com/flydevs/chat-app-api/rest-api/src/config"
	"github.com/flydevs/chat-app-api/rest-api/src/controller"
	"github.com/flydevs/chat-app-api/rest-api/src/repository"
	"github.com/flydevs/chat-app-api/rest-api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router               = gin.Default()
	messaging_controller controller.MessagingControllerInterface
	users_controller     controller.UsersControllerInterface
	oauth_controller     controller.OauthControllerInterface
)

func StartApplication() {
	messaging_controller = controller.NewMessagingController(services.NewMessagingService(repository.GetMessagingRepository()))
	users_controller = controller.NewUsersController(services.NewUsersService(repository.GetUsersRepository()))
	oauth_controller = controller.NewOauthController(services.NewOauthService(repository.GetOauthRepository()))
	mapUrls()
	router.Run(config.Config["ADDRESS"] + config.Config["PORT"])
}
