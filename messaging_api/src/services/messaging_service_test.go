package services

/*
var (
	mess_test_service MessagingService
	mockproto         = true
)

type mockprotoclient struct {
}

func (mpc mockprotoclient) GetUser(mockuuid string) ([]*users.User, server_message.Svr_message) {
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
		mess_test_service = NewMessagingService(db.GetMessagingDBRepository(), mockprotoclient{}, twilio.NewTwilioRepository())
	} else {
		mess_test_service = NewMessagingService(db.GetMessagingDBRepository(), users_client.GetUsersProtoClient(), twilio.NewTwilioRepository())
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var (
	test1mock = "71bcb82f-65ff-4623-bf9b-7d158ed746e6"
	test2mock = "e765f42d-28cb-4a2a-b67f-1e5441d0a6fd"
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
	UC2 = conversation.UserConversation{
		UserUuid: test2mock,
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
	//	defer db_ctrl.DBClient.Flush()
	test1, msg := mess_test_service.CreateConversation(CreateConversationC1)
	if msg.GetStatus() != 200 {
		t.Error(msg)
		return
	}
	UCS := conversation.CreateUserConversationRequest{UserConversationSlice: []conversation.UserConversation{UC1}}
	UCS.Conversation.Uuid = test1.Uuid
	msg = mess_test_service.CreateUserConversation(UCS)
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
	UCS.Conversation.Uuid = test2.Uuid
	msg = mess_test_service.CreateUserConversation(UCS)
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
*/
/*
func TestCreateGetMessage(t *testing.T) {
	//	defer db_ctrl.DBClient.Flush()
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
	UC2.ConversationUuid = convo_uuid.Uuid
	_, msg = mess_test_service.CreateUserConversation(UC2)
	if msg.GetStatus() != 200 {
		t.Error(msg)
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
		t.Error(fmt.Sprintf("userConversation not updated\n %v, \n%+v", convos[0].UserConversation.LastAccessUuid, result[1].Uuid))
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
*/
