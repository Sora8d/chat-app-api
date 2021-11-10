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

type ConversationResponse struct {
	Conversation `json:"conversation"`
	UserConversation
	Participants []UserConversation `json:"participants"`
}

//Later youll have to do validations
