package conversation

import (
	"github.com/Sora8d/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/users"
)

type Conversation struct {
	Id        int64   `json:"id"`
	Uuid      string  `json:"uuid"`
	TwilioSid string  `json:"twilio_sid"`
	Type      int32   `json:"type"`
	CreatedAt float32 `json:"created_at"`
	//DeletedAt
	LastMessage      message.Message `json:"last_message"`
	ConversationInfo `json:"info,omitempty"`
}

func (c *Conversation) SetTwilioSid(twilio_sid string) {
	c.TwilioSid = twilio_sid
}

type ConversationInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
}

//
type UserConversation struct {
	Id               int64   `json:"id"`
	Uuid             string  `json:"uuid"` //This is necessary, you could also just use this instead of the conversation UUID when sendin info to the Front.
	TwilioSid        string  `json:"twilio_sid"`
	UserId           int64   `json:"user_id"`
	UserUuid         string  `json:"user_uuid"` //Later check the necessity of having so many identifiers
	ConversationId   int64   `json:"conversation_id"`
	ConversationUuid string  `json:"conversation_uuid"`
	LastAccessUuid   string  `json:"last_access_uuid"` //Cambio a timestamp
	CreatedAt        float32 `json:"created_at"`
}

func (uc *UserConversation) SetTwilioSid(twilio_sid string) {
	uc.TwilioSid = twilio_sid
}

func (uc *UserConversation) SetUserId(id int64) {
	uc.UserId = id
}

type UserConversationSlice []UserConversation

func (ucs *UserConversationSlice) GetUuidsStringSlice() []string {
	uuids := []string{}
	for _, uc := range *ucs {
		uuids = append(uuids, uc.UserUuid)
	}
	return uuids
}

func (ucs *UserConversationSlice) ParseIds(user_slice []*users.User) server_message.Svr_message {
	if len(*ucs) != len(user_slice) {
		return server_message.NewBadRequestError("one of the users given doesnt exist")
	}
	for index, uc := range *ucs {
		found := false
		for index_users, user := range user_slice {
			if user.Uuid == uc.UserUuid {
				found = true
				uc.UserId = user.Id
				user_slice = append(user_slice[:index_users], user_slice[index_users+1:]...)
				break
			}
		}
		if !found {
			return server_message.NewBadRequestError("one of the users given doesnt exist")
		}
		(*ucs)[index] = uc
	}
	return nil
}

type ConversationAndParticipants struct {
	Conversation `json:"conversation"`
	UserConversation
	Participants UserConversationSlice `json:"participants"`
}

type ConversationAndParticipantsSlice []ConversationAndParticipants

func (cp *ConversationAndParticipants) SetUserConversationSlice(ucs []UserConversation) {
	cp.Participants = ucs
}
func (cp *ConversationAndParticipants) SetUserConversation(uc UserConversation) {
	cp.UserConversation = uc
}
func (cp *ConversationAndParticipants) SetConversation(convo Conversation) {
	cp.Conversation = convo
}

type CreateUserConversationRequest struct {
	UserConversationSlice UserConversationSlice `json:"user_conversation_slice"`
	Conversation          Conversation          `json:"conversation"`
}

func (cucr *CreateUserConversationRequest) SetUserConversationSlice(ucs []UserConversation) {
	cucr.UserConversationSlice = ucs
}
func (cucr *CreateUserConversationRequest) SetConversation(convo Conversation) {
	cucr.Conversation = convo
}

//Later youll have to do validations
