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
	CreateMessage(context.Context, *messaging.CreateMessageRequest) domain.Response
	CreateConversation(context.Context, *messaging.Conversation) domain.Response
	CreateUserConversation(context.Context, *messaging.CreateUserConversationRequest) domain.Response
	KickUser(context.Context, *messaging.KickUserRequest) domain.Response
	GetConversationsByUser(context.Context, *messaging.Uuid) domain.Response
	GetMessagesByConversation(context.Context, *messaging.GetMessages) domain.Response
	UpdateMessage(context.Context, *messaging.Message) domain.Response
	UpdateConversationInfo(context.Context, *messaging.Conversation) domain.Response
}

func NewMessagingService(msg_repo repository.MessagingRepositoryInterface) MessagingServiceInterface {
	return &messagingService{msg_repo: msg_repo}
}

//TODO: later create context

func (ms messagingService) CreateMessage(ctx context.Context, request *messaging.CreateMessageRequest) domain.Response {
	return Response.CreateResponse(ms.msg_repo.CreateMessage(ctx, request))

}

func (ms messagingService) CreateConversation(ctx context.Context, request *messaging.Conversation) domain.Response {
	return Response.CreateResponse(ms.msg_repo.CreateConversation(ctx, request))
}

func (ms messagingService) CreateUserConversation(ctx context.Context, request *messaging.CreateUserConversationRequest) domain.Response {
	return Response.CreateResponse(nil, ms.msg_repo.CreateUserConversation(ctx, request))
}

func (ms messagingService) KickUser(ctx context.Context, request *messaging.KickUserRequest) domain.Response {
	return Response.CreateResponse(nil, ms.msg_repo.KickUser(ctx, request))
}

func (ms messagingService) GetConversationsByUser(ctx context.Context, request *messaging.Uuid) domain.Response {
	return Response.CreateResponse(ms.msg_repo.GetConversationsByUser(ctx, request))
}

func (ms messagingService) GetMessagesByConversation(ctx context.Context, request *messaging.GetMessages) domain.Response {
	return Response.CreateResponse(ms.msg_repo.GetMessagesByConversation(ctx, request))
}

func (ms messagingService) UpdateMessage(ctx context.Context, request *messaging.Message) domain.Response {
	return Response.CreateResponse(ms.msg_repo.UpdateMessage(ctx, request))
}

func (ms messagingService) UpdateConversationInfo(ctx context.Context, request *messaging.Conversation) domain.Response {
	return Response.CreateResponse(ms.msg_repo.UpdateConversationInfo(ctx, request))
}
