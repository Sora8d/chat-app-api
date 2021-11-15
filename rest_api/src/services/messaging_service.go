package services

import (
	"context"

	"github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/rest-api/src/domain"
	"github.com/flydevs/chat-app-api/rest-api/src/repository"
)

var Response domain.Response

type messagingService struct {
}

type MessagingServiceInterface interface {
	CreateMessage(*messaging.CreateMessageRequest) domain.Response
	CreateConversation(*messaging.Conversation) domain.Response
	CreateUserConversation(*messaging.CreateUserConversationRequest) domain.Response
	GetConversationsByUser(*messaging.Uuid) domain.Response
	GetMessagesByConversation(*messaging.MessageRequest) domain.Response
	UpdateMessage(*messaging.Message) domain.Response
	UpdateConversationInfo(*messaging.Conversation) domain.Response
}

func NewMessagingService() MessagingServiceInterface {
	return &messagingService{}
}

//TODO: later create context

func (ms messagingService) CreateMessage(request *messaging.CreateMessageRequest) domain.Response {
	repository := repository.GetMessagingRepository()
	ctx := context.Background()
	return Response.CreateResponse(repository.CreateMessage(ctx, request))

}

func (ms messagingService) CreateConversation(request *messaging.Conversation) domain.Response {
	repository := repository.GetMessagingRepository()
	ctx := context.Background()
	return Response.CreateResponse(repository.CreateConversation(ctx, request))
}

func (ms messagingService) CreateUserConversation(request *messaging.CreateUserConversationRequest) domain.Response {
	repository := repository.GetMessagingRepository()
	ctx := context.Background()
	return Response.CreateResponse(nil, repository.CreateUserConversation(ctx, request))
}

func (ms messagingService) GetConversationsByUser(request *messaging.Uuid) domain.Response {
	repository := repository.GetMessagingRepository()
	ctx := context.Background()
	return Response.CreateResponse(repository.GetConversationsByUser(ctx, request))
}

func (ms messagingService) GetMessagesByConversation(request *messaging.MessageRequest) domain.Response {
	repository := repository.GetMessagingRepository()
	ctx := context.Background()
	return Response.CreateResponse(repository.GetMessagesByConversation(ctx, request))
}

func (ms messagingService) UpdateMessage(request *messaging.Message) domain.Response {
	repository := repository.GetMessagingRepository()
	ctx := context.Background()
	return Response.CreateResponse(repository.UpdateMessage(ctx, request))
}

func (ms messagingService) UpdateConversationInfo(request *messaging.Conversation) domain.Response {
	repository := repository.GetMessagingRepository()
	ctx := context.Background()
	return Response.CreateResponse(repository.UpdateConversationInfo(ctx, request))
}
