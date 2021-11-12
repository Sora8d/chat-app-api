package app

import (
	"github.com/flydevs/chat-app-api/rest-api/src/config"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	//mapUrls()
	router.Run(config.Config["ADDRESS"] + config.Config["PORT"])
}
