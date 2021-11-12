package conversation

type Conversation struct {
	Id        int64   `json:"id"`
	Uuid      string  `json:"uuid"`
	TwilioSid string  `json:"twilio_sid"`
	Type      int32   `json:"type"`
	CreatedAt float32 `json:"created_at"`
	//DeletedAt
	LastMessageUuid  string `json:"last_message_uuid"` //
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

type UserConverstionSlice []UserConversation

func (uc *UserConversation) SetUserId(id int64) {
	uc.UserId = id
}

type ConversationAndParticipants struct {
	Conversation `json:"conversation"`
	UserConversation
	Participants UserConverstionSlice `json:"participants"`
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
	UserConversationSlice UserConverstionSlice `json:"user_conversation_slice"`
	Conversation          Conversation         `json:"conversation"`
}

func (cucr *CreateUserConversationRequest) SetUserConversationSlice(ucs []UserConversation) {
	cucr.UserConversationSlice = ucs
}
func (cucr *CreateUserConversationRequest) SetConversation(convo Conversation) {
	cucr.Conversation = convo
}

//Later youll have to do validations
