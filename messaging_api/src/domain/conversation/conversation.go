package conversation

type Conversation struct {
	Id               int64   `json:"id"`
	Uuid             string  `json:"uuid"`
	Type             int32   `json:"type"`
	CreatedAt        float32 `json:"created_at"`
	LastMessageUuid  string  `json:"last_message_uuid"`
	ConversationInfo `json:"info"`
}

type ConversationInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
}

type UserConversation struct {
	Id               int64   `json:"id"`
	Uuid             string  `json:"uuid"`
	UserId           int64   `json:"user_id"`
	UserUuid         string  `json:"user_uuid"`
	ConversationId   int64   `json:"conversation_id"`
	ConversationUuid string  `json:"conversation_uuid"`
	LastAccessUuid   string  `json:"last_access_uuid"`
	CreatedAt        float32 `json:"created_at"`
}

//Later youll have to do validations
