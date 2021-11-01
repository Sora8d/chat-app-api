package services

import "github.com/flydevs/chat-app-api/messaging-api/src/repository/db"

type MessagingService interface{}

type messagingService struct {
	dbrepo db.MessagingDBRepository
}

func NewMessagingService(dbrepo db.MessagingDBRepository) MessagingService {
	return &messagingService{dbrepo: dbrepo}
}

func (ms *messagingService) CreateConversation()          {}
func (ms *messagingService) CreateMessage()               {}
func (ms *messagingService) CreateUserConversation()      {}
func (ms *messagingService) GetConversationsByUser()      {}
func (ms *messagingService) GetConversationByUuid()       {}
func (ms *messagingService) GetMessagesByAuthorId()       {}
func (ms *messagingService) GetMessagesByConversationId() {}
func (ms *messagingService) GetUserConversationsForUser() {}
