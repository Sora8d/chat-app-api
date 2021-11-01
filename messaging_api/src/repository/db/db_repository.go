package db

import (
	"os/user"

	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/client/postgresql"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/conversation"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
)

const (
	queryInsertConversation     = "INSERT INTO conversation(type, name, description, avatar_url) VALUES ($1,$2,$3,$4) RETURNING uuid;"
	queryCreateUserConversation = "INSERT INTO user_conversation(user_id, user_uuid, conversation_id, last_access_at) VALUES($1, $2, $3, $4);"
	queryCreateMessage          = "INSERT INTO message(conversation_id, conversation_uuid, author_id, author_uuid, text) VALUES($1, $2, $3, $4, $5);"
	//	queryGetConversationById         = "SELECT id, uuid, type, created_at, updated_at, name, description, avatar_url FROM conversation WHERE uuid=$1"
	queryGetConversationsFromUser    = "SELECT c.id, c.uuid, c.type, c.created_at, c.updated_at, c.name, c.description, c.avatar_url FROM conversation JOIN user_conversation uc ON c.id = uc.conversation_id WHERE uc.user_id =$1;"
	queryGetConversationByUuid       = "SELECT id, uuid, type, created_at, updated_at, name, description, avatar_url FROM conversation WHERE id=$1;"
	queryGetMessagesByConversationId = "SELECT m.id, m.uuid, m.conversation_id, m.conversation_uuid, m.author_id, m.author_uuid, m.text, m.created_at, m.updated_at FROM message m JOIN conversation c ON m.conversation_id = c.id WHERE c.id=$1;"
	queryGetMessageByAuthorId        = "SELECT m.id, m.uuid, conversation_id, conversation_uuid, author_id, author_uuid, text, created_at, updated_at FROM message WHERE author_id=$1;"
	queryGetUserConversationForUser  = "SELECT "
)

type MessagingDBRepository interface {
	CreateConversation(conversation.Conversation) (*string, server_message.Svr_message)
	CreateUserConversation(user.User, int64, string) server_message.Svr_message
	CreateMessage(message.Message) server_message.Svr_message
	GetConversationsByUser(int64) ([]conversation.Conversation, server_message.Svr_message)
	GetConversationByUuid(string) (*conversation.Conversation, server_message.Svr_message)
	GetMessagesByConversationId(int64) ([]message.Message, server_message.Svr_message)
	GetMessagesByAuthorId(int64) ([]message.Message, server_message.Svr_message)
	//	GetUserConversationsForUser(int64) ([]conversation.UserConversation, server_message.Svr_message)
}

type messagingDBRepository struct {
}

func GetMessagingDBRepository() MessagingDBRepository {
	return &messagingDBRepository{}
}

func (dbr *messagingDBRepository) CreateConversation(convo conversation.Conversation) (*string, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryInsertConversation, convo.Type, convo.Name, convo.Description, convo.AvatarUrl)
	var uuid string
	if err := row.Scan(&uuid); err != nil {
		logger.Error("error in CreateConversation function", err)
		return nil, server_message.NewInternalError()
	}
	return &uuid, nil
}
func (dbr *messagingDBRepository) GetConversationsByUser(user_id int64) ([]conversation.Conversation, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	rows, err := dbclient.Query(queryGetConversationsFromUser, user_id)
	if err != nil {
		logger.Error("error in GetConversationsByUser function, in the query execution", err)
		return nil, server_message.NewInternalError()
	}
	var convos []conversation.Conversation
	for rows.Next() {
		convo := conversation.Conversation{}
		if err := rows.Scan(&convo.Id, &convo.Uuid, &convo.Type, &convo.CreatedAt, &convo.UpdatedAt, &convo.Name, &convo.Description, &convo.AvatarUrl); err != nil {
			logger.Error("error in GetConversationsByUser function, scanning rows", err)
			return nil, server_message.NewInternalError()
		}
		convos = append(convos, convo)
	}
	if len(convos) == 0 {
		return nil, server_message.NewNotFoundError("no conversations where found in which this user partcipates")
	}
	return convos, nil
}
func (dbr *messagingDBRepository) GetConversationByUuid(uuid string) (*conversation.Conversation, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryGetConversationByUuid, uuid)
	convo := conversation.Conversation{}
	if err := row.Scan(&convo.Id, &convo.Uuid, &convo.Type, &convo.CreatedAt, &convo.UpdatedAt, &convo.Name, &convo.Description, &convo.AvatarUrl); err != nil {
		logger.Error("error in GetConversationByUuid function", err)
		return nil, server_message.NewInternalError()
	}
	return &convo, nil
}
func (dbr *messagingDBRepository) CreateMessage(msg message.Message) server_message.Svr_message {
	dbclient := postgresql.GetSession()
	err := dbclient.Execute(queryCreateMessage, msg.ConversationId, msg.ConversationUuid, msg.AuthorId, msg.AuthorUuid, msg.Text)
	if err != nil {
		logger.Error("error in CreateMessage function", err)
		return server_message.NewInternalError()
	}
	return nil
}
func (dbr *messagingDBRepository) GetMessagesByConversationId(id int64) ([]message.Message, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	rows, err := dbclient.Query(queryGetMessagesByConversationId, id)
	if err != nil {
		logger.Error("error in getmessagebyconversationid function, getting rows", err)
		return nil, server_message.NewInternalError()
	}
	msgs := []message.Message{}
	for rows.Next() {
		msg := message.Message{}
		if err := rows.Scan(&msg.Id, &msg.Uuid, &msg.ConversationId, &msg.ConversationUuid, &msg.AuthorId, &msg.AuthorUuid, &msg.Text, &msg.CreatedAt, &msg.UpdatedAt); err != nil {
			logger.Error("error in getmessagebyconversationid function, scanning rows", err)
			return nil, server_message.NewInternalError()
		}
		msgs = append(msgs, msg)
	}
	if len(msgs) == 0 {
		return nil, server_message.NewNotFoundError("no messages where found")
	}
	return msgs, nil

}
func (dbr *messagingDBRepository) GetMessagesByAuthorId(id int64) ([]message.Message, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	rows, err := dbclient.Query(queryGetMessageByAuthorId, id)
	if err != nil {
		logger.Error("error in getmessagebyauthorid function, getting rows", err)
		return nil, server_message.NewInternalError()
	}
	msgs := []message.Message{}
	for rows.Next() {
		msg := message.Message{}
		if err := rows.Scan(&msg.Id, &msg.Uuid, &msg.ConversationId, &msg.ConversationUuid, &msg.AuthorId, &msg.AuthorUuid, &msg.Text, &msg.CreatedAt, &msg.UpdatedAt); err != nil {
			logger.Error("error in getmessagebyathorid function, scanning rows", err)
			return nil, server_message.NewInternalError()
		}
		msgs = append(msgs, msg)
	}
	if len(msgs) == 0 {
		return nil, server_message.NewNotFoundError("no messages where found")
	}
	return msgs, nil
}

func (dbr *messagingDBRepository) CreateUserConversation(user user.User, conversation_id int64, last_msg_uuid string) server_message.Svr_message {
	return nil
}
