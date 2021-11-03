package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/flydevs/chat-app-api/common/server_message"
	db_ctrl "github.com/flydevs/chat-app-api/messaging-api/src/clients/testing_tools"
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
	switch mockuuid {
	case "123e4568-e89b-12d3-a456-426655440000":
		mock_user := users.User{
			Uuid: mockuuid,
			Id:   2,
		}
		return &mock_user, server_message.NewCustomMessage(http.StatusOK, "user retrieved")
	case "123e4567-e89b-12d3-a456-426655440000":
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
	defer db_ctrl.DBClient.Flush()
	test1, msg := mess_test_service.CreateConversation(CreateConversationC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
	}
	result1, msg := mess_test_service.GetConversationByUuid(test1.Uuid)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	if result1.Uuid == "" || result1.Type != 1 {
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
	if result2.Uuid == "" || result2.Type != 2 || result2.ConversationInfo != CreateConversationC2.ConversationInfo {
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
		AuthorUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
	test1text = "test1"
	test1mock = "123e4567-e89b-12d3-a456-426655440000"
)

func TestCreateGetMessage(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
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
	if result.Uuid == "" || result.AuthorId != 1 || result.Text != test1text || result.AuthorUuid != test1mock {
		t.Error(fmt.Sprintf("received data is not right, %+v", result))
		return
	}
}

var (
	CreateUserConversationC1 = conversation.Conversation{
		Type: 1,
	}
	CreateUserConversationUC1 = conversation.UserConversation{
		UserUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
	useruuidmock = "123e4567-e89b-12d3-a456-426655440000"
)

func TestCreateGetUserConversation(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
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

	if result.Uuid == "" || result.UserId != 1 || result.UserUuid != useruuidmock {
		t.Error(fmt.Sprintf("received data is not right, %+v", result))
		return
	}
}

var (
	GetConvsByUserC1 = conversation.Conversation{
		Type: 1,
	}
	GetConvsByUserC2 = conversation.Conversation{
		Type: 2,
		ConversationInfo: conversation.ConversationInfo{
			Name:        "Convo1",
			Description: "Convo1_Test",
			AvatarUrl:   "test1.jpg",
		},
	}
	GetConvsByUserUC = conversation.UserConversation{
		UserUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
)

func TestGetConversationsByUser(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	uuid1, response_msg := mess_test_service.CreateConversation(GetConvsByUserC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	uc := GetConvsByUserUC
	uc.ConversationUuid = uuid1.Uuid
	_, response_msg = mess_test_service.CreateUserConversation(uc)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	uuid2, response_msg := mess_test_service.CreateConversation(GetConvsByUserC2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	uc = GetConvsByUserUC
	uc.ConversationUuid = uuid2.Uuid
	_, response_msg = mess_test_service.CreateUserConversation(uc)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.GetConversationsByUser(useruuidmock)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if len(result) != 2 || result[0].Uuid != uuid1.Uuid || result[1].Uuid != uuid2.Uuid {
		t.Error(fmt.Sprintf("There is something wrong: %+v", result))
	}
}

var (
	GetMsgC1 = conversation.Conversation{
		Type: 1,
	}
	GetMsgC2 = conversation.Conversation{
		Type: 2,
		ConversationInfo: conversation.ConversationInfo{
			Name:        "Convo1",
			Description: "Convo1_Test",
			AvatarUrl:   "test1.jpg",
		},
	}
	GetMsgM1 = message.Message{
		Text:       "test1",
		AuthorUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
	GetMsgM2 = message.Message{
		Text:       "test2",
		AuthorUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
	GetMsgM3 = message.Message{
		Text:       "test3",
		AuthorUuid: "123e4568-e89b-12d3-a456-426655440000",
	}
	test2mock = "123e4568-e89b-12d3-a456-426655440000"
)

func TestGetMessagesByAuthor(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	uuid, response_msg := mess_test_service.CreateConversation(GetMsgC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	GetMsgM1.ConversationUuid = uuid.Uuid
	GetMsgM3.ConversationUuid = uuid.Uuid

	uuid, response_msg = mess_test_service.CreateConversation(GetMsgC2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	GetMsgM2.ConversationUuid = uuid.Uuid

	_, response_msg = mess_test_service.CreateMessage(GetMsgM1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	_, response_msg = mess_test_service.CreateMessage(GetMsgM2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	_, response_msg = mess_test_service.CreateMessage(GetMsgM3)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.GetMessagesByAuthor(test1mock)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if len(result) != 2 || result[0].AuthorId != 1 || result[1].AuthorId != 1 || result[0].Text != "test1" || result[1].Text != "test2" || result[0].ConversationId == result[1].ConversationId {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}

	result, response_msg = mess_test_service.GetMessagesByAuthor(test2mock)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if len(result) != 1 || result[0].AuthorId != 2 || result[0].Text != "test3" {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}
}

func TestGetMessagesByConversation(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	uuid1, response_msg := mess_test_service.CreateConversation(GetMsgC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	GetMsgM1.ConversationUuid = uuid1.Uuid
	GetMsgM3.ConversationUuid = uuid1.Uuid

	uuid2, response_msg := mess_test_service.CreateConversation(GetMsgC2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	GetMsgM2.ConversationUuid = uuid2.Uuid

	_, response_msg = mess_test_service.CreateMessage(GetMsgM1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	_, response_msg = mess_test_service.CreateMessage(GetMsgM2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	_, response_msg = mess_test_service.CreateMessage(GetMsgM3)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.GetMessagesByConversation(uuid1.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if len(result) != 2 || result[0].AuthorId != 1 || result[1].AuthorId != 2 || result[0].Text != "test1" || result[1].Text != "test3" || result[0].ConversationId != result[1].ConversationId {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}

	result, response_msg = mess_test_service.GetMessagesByConversation(uuid2.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if len(result) != 1 || result[0].AuthorId != 1 || result[0].Text != "test2" {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}
}

var (
	GetUCC1 = conversation.Conversation{
		Type: 1,
	}
	GetUCC2 = conversation.Conversation{
		Type: 2,
		ConversationInfo: conversation.ConversationInfo{
			Name:        "Convo1",
			Description: "Convo1_Test",
			AvatarUrl:   "test1.jpg",
		},
	}
	GetUCUC1 = conversation.UserConversation{
		UserUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
	GetUCUC2 = conversation.UserConversation{
		UserUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
	GetUCUC3 = conversation.UserConversation{
		UserUuid: "123e4568-e89b-12d3-a456-426655440000",
	}
)

func TestGetUserConversationByUser(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	uuid1, response_msg := mess_test_service.CreateConversation(GetUCC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	GetUCUC1.ConversationUuid = uuid1.Uuid
	GetUCUC3.ConversationUuid = uuid1.Uuid

	_, response_msg = mess_test_service.CreateUserConversation(GetUCUC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	_, response_msg = mess_test_service.CreateUserConversation(GetUCUC3)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	uuid2, response_msg := mess_test_service.CreateConversation(GetUCC2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	GetUCUC2.ConversationUuid = uuid2.Uuid
	_, response_msg = mess_test_service.CreateUserConversation(GetUCUC2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.GetUserConversationsForUser(test1mock)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if len(result) != 2 || result[0].ConversationUuid != uuid1.Uuid || result[0].ConversationUuid == result[1].ConversationUuid || result[0].UserId != 1 {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}

	result, response_msg = mess_test_service.GetUserConversationsForUser(test2mock)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if len(result) != 1 || result[0].ConversationUuid != uuid2.Uuid || result[0].UserId != 2 {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}
}

func TestGetUserConversationByConversation(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	uuid1, response_msg := mess_test_service.CreateConversation(GetUCC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	GetUCUC1.ConversationUuid = uuid1.Uuid
	GetUCUC3.ConversationUuid = uuid1.Uuid

	_, response_msg = mess_test_service.CreateUserConversation(GetUCUC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	_, response_msg = mess_test_service.CreateUserConversation(GetUCUC3)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	uuid2, response_msg := mess_test_service.CreateConversation(GetUCC2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	GetUCUC2.ConversationUuid = uuid2.Uuid
	_, response_msg = mess_test_service.CreateUserConversation(GetUCUC2)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.GetUserConversationsForConversation(uuid1.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if len(result) != 2 || result[0].ConversationUuid != uuid1.Uuid || result[0].ConversationUuid != result[1].ConversationUuid || result[0].UserId != 1 {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}

	result, response_msg = mess_test_service.GetUserConversationsForConversation(uuid2.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if len(result) != 1 || result[0].ConversationUuid != uuid2.Uuid || result[0].UserId != 1 {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}
}

var (
	UpdateConvoC1 = conversation.Conversation{
		Type: 1,
	}
	UpdateConvoM1 = message.Message{
		Text:       "test1",
		AuthorUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
)

func TestUpateConversationLastMsg(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	uuid1, response_msg := mess_test_service.CreateConversation(GetUCC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	UpdateConvoM1.ConversationUuid = uuid1.Uuid

	msg, response_msg := mess_test_service.CreateMessage(UpdateConvoM1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.UpdateConversationLastMsg(uuid1.Uuid, msg.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if result.Uuid != uuid1.Uuid || result.LastMessageUuid != msg.Uuid {
		t.Error("somethings wrong i can feel it.", result)
	}

	msg, response_msg = mess_test_service.CreateMessage(UpdateConvoM1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg = mess_test_service.UpdateConversationLastMsg(uuid1.Uuid, msg.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if result.Uuid != uuid1.Uuid || result.LastMessageUuid != msg.Uuid {
		t.Error("somethings wrong i can feel it.", result)
	}
}

var (
	changetext = "change1"
)

func TestUpdateMessage(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	uuid, response_msg := mess_test_service.CreateConversation(GetUCC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	UpdateConvoM1.ConversationUuid = uuid.Uuid

	uuid, response_msg = mess_test_service.CreateMessage(UpdateConvoM1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.UpdateMessage(uuid.Uuid, changetext)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if result.Text != changetext {
		t.Error(fmt.Sprintf("error updating msg: %+v", result))
	}
}

var (
	UpdateUC1 = conversation.UserConversation{
		UserUuid: "123e4567-e89b-12d3-a456-426655440000",
	}
)

func TestUserConversationLastAccess(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	uuid, response_msg := mess_test_service.CreateConversation(GetUCC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	UpdateUC1.ConversationUuid = uuid.Uuid

	UpdateConvoM1.ConversationUuid = uuid.Uuid

	uuidmsg, response_msg := mess_test_service.CreateMessage(UpdateConvoM1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	uuiduc, response_msg := mess_test_service.CreateUserConversation(UpdateUC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.UserConversationUpdateLastAccess(uuiduc.Uuid, uuidmsg.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if result.LastAccessUuid != uuidmsg.Uuid {
		t.Error(result)
	}
}
