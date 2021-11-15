package app

import (
	"github.com/flydevs/chat-app-api/rest-api/src/config"
	"github.com/flydevs/chat-app-api/rest-api/src/controller"
	"github.com/flydevs/chat-app-api/rest-api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router               = gin.Default()
	messaging_controller controller.MessagingControllerInterface
)

func StartApplication() {
	messaging_controller = controller.NewMessagingController(services.NewMessagingService())
	mapUrls()
	router.Run(config.Config["ADDRESS"] + config.Config["PORT"])
}
