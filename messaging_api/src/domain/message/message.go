package message

type Message struct {
	Id               int64   `json:"id"`
	Uuid             string  `json:"uuid"`
	TwilioSid        string  `json:"twilio_sid"`
	ConversationId   int64   `json:"conversation_id"`
	ConversationUuid string  `json:"conversation_uuid"`
	AuthorId         int64   `json:"author_id"` //Not needed anymore since messages will be searched from conversations
	AuthorUuid       string  `json:"author_uuid"`
	Text             string  `json:"text"`
	CreatedAt        float32 `json:"created_at"`
	UpdatedAt        float32 `json:"updated_at"`
}
