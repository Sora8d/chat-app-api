package controllers

import (
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/messaging-api/src/services"
)

type messagingController struct {
	pb.UnimplementedMessagingProtoInterfaceServer

	svc services.MessagingService
}

type MessagingController interface {
	CreateConversation()
	CreateMessage()
	CreateUserConversation()

	GetConversationsByUser()
	GetConversationByUuid()
	UpdateConversationLastMsg()

	GetMessagesBuAuthor()
	GetMessagesByConversation()
	GetMessageByUuid()
	UpdateMessage()

	GetUserConversationForUser()
	GetUserConversationForConversation()
	GetUserConversationByUuid()
	UpdateUCLastAccess()
}

func (mc *messagingController) CreateConversation() {

}
