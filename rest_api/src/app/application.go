package app

import (
	"github.com/flydevs/chat-app-api/rest-api/src/config"
	"github.com/flydevs/chat-app-api/rest-api/src/controller"
	"github.com/flydevs/chat-app-api/rest-api/src/repository"
	"github.com/flydevs/chat-app-api/rest-api/src/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router               = gin.Default()
	messaging_controller controller.MessagingControllerInterface
	users_controller     controller.UsersControllerInterface
	oauth_controller     controller.OauthControllerInterface
)

func StartApplication() {
	new_config := cors.DefaultConfig()
	new_config.AllowOrigins = append(new_config.AllowOrigins, "*")
	new_config.AllowMethods = append(new_config.AllowMethods, "OPTION")
	new_config.AllowHeaders = append(new_config.AllowHeaders, "access-token")
	router.Use(cors.New(new_config))
	messaging_controller = controller.NewMessagingController(services.NewMessagingService(repository.GetMessagingRepository()))
	users_controller = controller.NewUsersController(services.NewUsersService(repository.GetUsersRepository()))
	oauth_controller = controller.NewOauthController(services.NewOauthService(repository.GetOauthRepository()))
	mapUrls()
	router.Run(config.Config["ADDRESS"] + config.Config["PORT"])
}
