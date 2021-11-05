package services

import (
	"fmt"
	"net/http"
	"os"
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
	mockproto         = false
)

type mockprotoclient struct {
}

func (mpc mockprotoclient) GetUser(mockuuid string) (*users.User, server_message.Svr_message) {
	switch mockuuid {
	case test2mock:
		mock_user := users.User{
			Uuid: mockuuid,
			Id:   2,
		}
		return &mock_user, server_message.NewCustomMessage(http.StatusOK, "user retrieved")
	case test1mock:
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

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var (
	test1mock = "f812a359-20d4-41b2-a841-2cc003a14594"
	test2mock = "0bf8c42d-3356-44fa-abd5-fb3fa7fa357b"
)

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
	UC1 = conversation.UserConversation{
		UserUuid: test1mock,
	}
	CreateMessageM1 = message.Message{
		Text:       "test1",
		AuthorUuid: test1mock,
	}
	test1text = "test1"
)

func TestCreateGetConversationInternal(t *testing.T) {
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

func TestCreateGetConversationExternal(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	test1, msg := mess_test_service.CreateConversation(CreateConversationC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	UC1.ConversationUuid = test1.Uuid
	_, msg = mess_test_service.CreateUserConversation(UC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	result, msg := mess_test_service.GetConversationsByUser(test1mock)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	if len(result) != 1 || result[0].Conversation.Uuid != test1.Uuid || result[0].Participants[0] != result[0].UserConversation {
		fmt.Println(len(result) != 1, result[0].Conversation.Uuid != test1.Uuid, result[0].Participants[0] == result[0].UserConversation)
		t.Error(fmt.Sprintf("received data is not right, %+v", result))
		return
	}
	test2, msg := mess_test_service.CreateConversation(CreateConversationC2)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	UC1.ConversationUuid = test2.Uuid
	_, msg = mess_test_service.CreateUserConversation(UC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	result, msg = mess_test_service.GetConversationsByUser(test1mock)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	if len(result) != 2 || result[0].Conversation.Uuid != test1.Uuid || result[1].Conversation.Uuid != test2.Uuid || result[0].Participants[0] != result[0].UserConversation || result[1].Participants[0].UserUuid != result[0].Participants[0].UserUuid {
		fmt.Println(len(result) != 2, result[0].Conversation.Uuid != test1.Uuid, result[1].Conversation.Uuid != test2.Uuid, result[0].Participants[0] != result[0].UserConversation, result[1].Participants[0].UserUuid != result[0].Participants[0].UserUuid)
		t.Error(fmt.Sprintf("received data is not right, %+v \n %+v", result[0], result[1]))
		return
	}

}

func TestCreateGetMessage(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	convo_uuid, response_msg := mess_test_service.CreateConversation(CreateConversationC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
		return
	}
	UC1.ConversationUuid = convo_uuid.Uuid
	uc_uuid, msg := mess_test_service.CreateUserConversation(UC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	req_msg := CreateMessageM1
	req_msg.ConversationUuid = convo_uuid.Uuid

	_, response_msg = mess_test_service.CreateMessage(req_msg)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
		return
	}

	req_msg.AuthorUuid = test2mock
	_, response_msg = mess_test_service.CreateMessage(req_msg)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
		return
	}

	result, response_msg := mess_test_service.GetMessagesByConversation(uc_uuid.Uuid, convo_uuid.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
		return
	}
	if len(result) != 2 || result[0].Uuid == "" || result[0].Text != test1text || result[1].AuthorUuid != test2mock {
		fmt.Println(len(result) != 1, result[0].Uuid == "", result[0].Text != test1text)
		t.Error(fmt.Sprintf("received data is not right, %+v", result))
		return
	}

	convos, msg := mess_test_service.GetConversationsByUser(test1mock)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	if convos[0].UserConversation.LastAccessUuid != result[1].Uuid {
		t.Error("userConversation not updated")
		return
	}

}

func TestUpdateMessage(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	convo_uuid, response_msg := mess_test_service.CreateConversation(CreateConversationC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	UC1.ConversationUuid = convo_uuid.Uuid
	uc_uuid, msg := mess_test_service.CreateUserConversation(UC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	req_msg := CreateMessageM1
	req_msg.ConversationUuid = convo_uuid.Uuid

	_, response_msg = mess_test_service.CreateMessage(req_msg)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
		return
	}

	req_msg.AuthorUuid = test2mock
	_, response_msg = mess_test_service.CreateMessage(req_msg)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
		return
	}

	result, response_msg := mess_test_service.GetMessagesByConversation(uc_uuid.Uuid, convo_uuid.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
		return
	}
	if result[0].Text != test1text {
		t.Error(fmt.Sprintf("received data is not right, %+v", result))
		return
	}

	changetext := "test2help"

	msg_updated, response_msg := mess_test_service.UpdateMessage(result[0].Uuid, changetext)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
		return
	}
	if msg_updated.Text != changetext {
		t.Error(fmt.Sprintf("error updating msg: %+v", result))
		return
	}

	result, response_msg = mess_test_service.GetMessagesByConversation(uc_uuid.Uuid, convo_uuid.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if result[0].Text == test1text || result[0].Text != changetext {
		fmt.Println(result[1].Text == test1text, result[1].Text != changetext)
		t.Error(fmt.Sprintf("received data has bad body, %+v\n%+v", result[0], result[1]))
		return
	}

	uc_result, msg := mess_test_service.GetConversationsByUser(test1mock)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	if uc_result[0].UserConversation.LastAccessUuid != result[1].Uuid {
		fmt.Printf("\n%+v\n%+v\n", result[0], result[1])
		fmt.Printf("%+v", uc_result[0].UserConversation.LastAccessUuid)
		t.Error("this shouldnt be the last_message read")
	}
}

var (
	convo_info = conversation.ConversationInfo{
		Name:        "change1",
		Description: "Difference1",
		AvatarUrl:   "goodavatar.jpg",
	}
)

func TestConversationUpdateInfo(t *testing.T) {
	defer db_ctrl.DBClient.Flush()
	test1, msg := mess_test_service.CreateConversation(CreateConversationC2)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	UC1.ConversationUuid = test1.Uuid
	_, msg = mess_test_service.CreateUserConversation(UC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	test2, msg := mess_test_service.CreateConversation(CreateConversationC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	UC1.ConversationUuid = test2.Uuid
	_, msg = mess_test_service.CreateUserConversation(UC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	result, msg := mess_test_service.GetConversationsByUser(test1mock)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	if len(result) != 2 || result[0].Conversation.Uuid != test1.Uuid || result[0].Participants[0] != result[0].UserConversation || result[0].Conversation.ConversationInfo == convo_info {
		fmt.Println(len(result) != 1, result[0].Conversation.Uuid != test1.Uuid, result[0].Participants[0] == result[0].UserConversation)
		t.Error(fmt.Sprintf("received data is not right, %+v", result))
		return
	}
	_, msg = mess_test_service.UpdateConversationInfo(result[0].Conversation.Uuid, convo_info)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	result, msg = mess_test_service.GetConversationsByUser(test1mock)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}

	if result[0].Conversation.ConversationInfo != convo_info {
		t.Error(fmt.Printf("\nconversation_info should have changed\n%+v", result[0].Conversation.ConversationInfo))
	}

}

/*
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

	if result.Uuid == "" {
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
		UserUuid: test1mock,
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
		AuthorUuid: test1mock,
	}
	GetMsgM2 = message.Message{
		Text:       "test2",
		AuthorUuid: test1mock,
	}
	GetMsgM3 = message.Message{
		Text:       "test3",
		AuthorUuid: test2mock,
	}
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
	if len(result) != 2 || result[0].Text != "test1" || result[1].Text != "test2" || result[0].ConversationId == result[1].ConversationId {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}

	result, response_msg = mess_test_service.GetMessagesByAuthor(test2mock)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if len(result) != 1 || result[0].Text != "test3" {
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
	if len(result) != 2 || result[0].Text != "test1" || result[1].Text != "test3" || result[0].ConversationId != result[1].ConversationId {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}

	result, response_msg = mess_test_service.GetMessagesByConversation(uuid2.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}
	if len(result) != 1 || result[0].Text != "test2" {
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
		UserUuid: test1mock,
	}
	GetUCUC2 = conversation.UserConversation{
		UserUuid: test1mock,
	}
	GetUCUC3 = conversation.UserConversation{
		UserUuid: test2mock,
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

	if len(result) != 1 || result[0].ConversationUuid != uuid1.Uuid {
		t.Error(fmt.Sprintf("%+v: there are differences in the results and the expected values: %+v", uuid2, result))
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

	if len(result) != 1 || result[0].ConversationUuid != uuid2.Uuid {
		t.Error(fmt.Sprintf("there are differences in the results and the expected values: %+v", result))
	}
}

var (
	UpdateConvoC1 = conversation.Conversation{
		Type: 1,
	}
	UpdateConvoM1 = message.Message{
		Text:       "test1",
		AuthorUuid: test1mock,
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


var (
	UpdateUC1 = conversation.UserConversation{
		UserUuid: test1mock,
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

var (
	change1Info = conversation.ConversationInfo{
		Name:        "test1",
		Description: "desct1",
		AvatarUrl:   "avatar.jpg",
	}
	change2Info = conversation.ConversationInfo{
		Name:        "test2",
		Description: "desct2",
		AvatarUrl:   "icon.jpg",
	}
)

func TestConversationUpdateInfo(t *testing.T) {
	defer db_ctrl.DBClient.Flush()

	uuid, response_msg := mess_test_service.CreateConversation(GetUCC1)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	result, response_msg := mess_test_service.UpdateConversationInfo(uuid.Uuid, change1Info)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	convo, response_msg := mess_test_service.GetConversationByUuid(result.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if convo.Uuid != uuid.Uuid || convo.ConversationInfo != change1Info {
		t.Error(fmt.Sprintf("something differs %+v", convo))
	}

	result, response_msg = mess_test_service.UpdateConversationInfo(uuid.Uuid, change2Info)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	convo, response_msg = mess_test_service.GetConversationByUuid(result.Uuid)
	if response_msg.GetStatus() != 200 {
		t.Error(response_msg)
	}

	if convo.Uuid != uuid.Uuid || convo.ConversationInfo != change2Info {
		t.Error(fmt.Sprintf("something differs %+v", convo))
	}
}
*/
