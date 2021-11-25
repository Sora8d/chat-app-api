package app

func mapUrls() {
	messageUrls()
	usersUrls()
	oauthUrls()
}

func messageUrls() {
	router.POST("/message", messaging_controller.CreateMessage)
	router.POST("/conversation", messaging_controller.CreateConversation)
	router.POST("/user_conversation", messaging_controller.CreateUserConversation)
	router.GET("/conversation", messaging_controller.GetConversationsByUser)
	router.GET("/message", messaging_controller.GetMessagesByConversation)
	router.PUT("/message", messaging_controller.UpdateMessage)
	router.PUT("/conversation/info", messaging_controller.UpdateConversationInfo)
}

func usersUrls() {
	router.POST("/user", users_controller.CreateUser)
	router.GET("/user", users_controller.GetUserProfileByUuid)
	router.PUT("/user", users_controller.UpdateUser)
}

func oauthUrls() {
	router.POST("/login", oauth_controller.LoginUser)
}
