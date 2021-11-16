package services

import (
	"context"

	"github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/rest-api/src/domain"
	"github.com/flydevs/chat-app-api/rest-api/src/repository"
)

var Response domain.Response

type messagingService struct {
	msg_repo repository.MessagingRepositoryInterface
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

func NewMessagingService(msg_repo repository.MessagingRepositoryInterface) MessagingServiceInterface {
	return &messagingService{msg_repo: msg_repo}
}

//TODO: later create context

func (ms messagingService) CreateMessage(request *messaging.CreateMessageRequest) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(ms.msg_repo.CreateMessage(ctx, request))

}

func (ms messagingService) CreateConversation(request *messaging.Conversation) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(ms.msg_repo.CreateConversation(ctx, request))
}

func (ms messagingService) CreateUserConversation(request *messaging.CreateUserConversationRequest) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(nil, ms.msg_repo.CreateUserConversation(ctx, request))
}

func (ms messagingService) GetConversationsByUser(request *messaging.Uuid) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(ms.msg_repo.GetConversationsByUser(ctx, request))
}

func (ms messagingService) GetMessagesByConversation(request *messaging.MessageRequest) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(ms.msg_repo.GetMessagesByConversation(ctx, request))
}

func (ms messagingService) UpdateMessage(request *messaging.Message) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(ms.msg_repo.UpdateMessage(ctx, request))
}

func (ms messagingService) UpdateConversationInfo(request *messaging.Conversation) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(ms.msg_repo.UpdateConversationInfo(ctx, request))
}
