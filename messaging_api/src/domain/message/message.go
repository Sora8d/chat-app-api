package message

type Message struct {
	Id               int64  `json:"id"`
	Uuid             string `json:"uuid"`
	ConversationId   int64  `json:"conversation_id"` //Delete this later
	ConversationUuid string `json:"conversation_uuid"`
	AuthorId         int64  `json:"author_id"`
	AuthorUuid       string `json:"author_uuid"`
	Text             string `json:"text"`
	CreatedAt        int    `json:"created_at"`
	UpdatedAt        int    `json:"updated_at"`
}
