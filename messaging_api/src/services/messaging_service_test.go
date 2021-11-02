package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/conversation"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/users"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/db"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/users_client"
)

var (
	mess_test_service MessagingService
	mockproto         = true
)

type mockprotoclient struct {
}

func (mpc mockprotoclient) GetUser(mockuuid string) (*users.User, server_message.Svr_message) {
	if mockuuid == "mock1" {
		mock_user := users.User{
			Uuid: mockuuid,
			Id:   1,
		}
		return &mock_user, server_message.NewCustomMessage(http.StatusOK, "user retrieved")
	}
	return nil, server_message.NewNotFoundError("user not found")

}

func init() {
	if mockproto {
		mess_test_service = NewMessagingService(db.GetMessagingDBRepository(), mockprotoclient{})
	} else {
		mess_test_service = NewMessagingService(db.GetMessagingDBRepository(), users_client.GetUsersProtoClient())
	}
}

var (
	CreateConversationC1 = conversation.Conversation{
		Type: 1,
	}
	CreateConversationC2 = conversation.Conversation{
		Type: 2,
		ConversationInfo: conversation.ConversationInfo{
			Name:        "Convo1",
			Description: "Convo1_Test",
			AvatarUrl:   "test1.jpg",
		},
	}
)

func TestCreateGetConversation(t *testing.T) {
	test1, msg := mess_test_service.CreateConversation(CreateConversationC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
	}
	result1, msg := mess_test_service.GetConversationByUuid(test1.Uuid)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	if result1.Id != 1 || result1.Uuid == "" || result1.Type != 1 {
		t.Error(fmt.Sprintf("received data is not right, %+v", result1))
		return
	}

	test2, msg := mess_test_service.CreateConversation(CreateConversationC2)
	if msg.GetStatus() != 200 {
		t.Error(msg)
	}
	result2, msg := mess_test_service.GetConversationByUuid(test2.Uuid)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	if result2.Id != 2 || result2.Uuid == "" || result2.Type != 2 || result2.ConversationInfo == CreateConversationC2.ConversationInfo {
		t.Error(fmt.Sprintf("received data is not right, %+v", result2))
		return
	}
}

var (
	CreateMessageC1 = conversation.Conversation{
		Type: 1,
	}
	CreateMessageM1 = message.Message{
		Text:       "test1",
		AuthorUuid: "mock1",
	}
	test1text = "test1"
)

func TestCreateGetMessage(t *testing.T) {
	convo_uuid, response_msg := mess_test_service.CreateConversation(CreateConversationC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	req_msg := CreateMessageM1
	req_msg.ConversationUuid = convo_uuid.Uuid

	test1, response_msg := mess_test_service.CreateMessage(req_msg)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.GetMessageByUuid(test1.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if result.Uuid == "" || result.AuthorId != 1 || result.ConversationId != 1 || result.Text == test1text {
		t.Error(fmt.Sprintf("received data is not right, %+v", result))
		return
	}
}

var (
	CreateUserConversationC1 = conversation.Conversation{
		Type: 1,
	}
	CreateUserConversationUC1 = conversation.UserConversation{
		UserUuid: "mock1",
	}
	useruuidmock = "mock1"
)

func TestCreateGetUserConversation(t *testing.T) {
	uuid, response_msg := mess_test_service.CreateConversation(CreateUserConversationC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	req_uc := CreateUserConversationUC1
	req_uc.ConversationUuid = uuid.Uuid

	test, response_msg := mess_test_service.CreateUserConversation(req_uc)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.GetUserConversationByUuid(test.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if result.Uuid == "" || result.UserId != 1 || result.ConversationId != 1 || result.UserUuid == useruuidmock {
		t.Error(fmt.Sprintf("received data is not right, %+v", result))
		return
	}
}
