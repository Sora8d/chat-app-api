package conversation

type Conversation struct {
	Id          int64  `json:"id"`
	Uuid        string `json:"uuid"`
	Type        int    `json:"type"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
}

type UserConversation struct {
	Id             int64  `json:"id"`
	UserId         int64  `json:"user_id"`
	UserUuid       string `json:"user_uuid"`
	ConversationId int64  `json:"conversation_id"`
	LastAccessAt   int    `json:"last_access_at"`
}
