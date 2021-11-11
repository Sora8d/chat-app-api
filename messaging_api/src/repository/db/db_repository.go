package db

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/clients/postgresql"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/conversation"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/uuids"
)

const (
	queryInsertConversation     = "INSERT INTO conversation(type, name, description, avatar_url, twilio_sid) VALUES ($1,$2,$3,$4,$5) RETURNING uuid;"
	queryCreateUserConversation = "INSERT INTO user_conversation(user_id, user_uuid, conversation_id, conversation_uuid, twilio_sid) VALUES($1, $2, $3, $4, $5) RETURNING uuid;"
	queryCreateMessage          = "INSERT INTO message_table(conversation_id, conversation_uuid, author_id, author_uuid, body, twilio_sid) VALUES($1, $2, $3, $4, $5, $6) RETURNING uuid;"

	queryGetConversationsFromUser = `SELECT c.id, c.uuid, c.twilio_sid, c.type, date_part('epoch',c.created_at), c.last_message_uuid, c.name, c.description, c.avatar_url,
	uc.id, uc.uuid, uc.twilio_sid, uc.user_id, uc.user_uuid, uc.conversation_id, uc.conversation_uuid, uc.last_access_uuid, date_part('epoch',uc.created_at)
	FROM conversation c JOIN user_conversation uc ON c.id = uc.conversation_id 
	WHERE uc.user_uuid =$1
	ORDER BY c.created_at;`
	queryGetConversationByUuid     = "SELECT id, uuid, twilio_sid, type, date_part('epoch',created_at), last_message_uuid, name, description, avatar_url FROM conversation WHERE uuid=$1;"
	queryConversationUpdateInfo    = "UPDATE conversation SET name=$2, description=$3, avatar_url=$4 WHERE uuid=$1 RETURNING uuid, name, description, avatar_url;"
	queryConversationUpdateMsgUuid = "UPDATE conversation SET last_message_uuid=$2 WHERE uuid=$1 RETURNING uuid, last_message_uuid;"

	queryGetMessagesByConversationUuid = "SELECT m.id, m.uuid, m.twilio_sid, m.conversation_id, m.conversation_uuid, m.author_id, m.author_uuid, m.body, date_part('epoch',m.created_at), date_part('epoch',m.updated_at) FROM message_table m JOIN conversation c ON m.conversation_id = c.id WHERE c.uuid=$1 ORDER BY m.created_at;"
	queryUpdateMessage                 = "UPDATE message_table SET body=$2, updated_at=timezone('utc'::text, now()) WHERE uuid=$1 RETURNING uuid, conversation_uuid, twilio_sid, body, date_part('epoch',updated_at);"

	queryGetUserConversationForUser         = "SELECT id, uuid, twilio_sid, user_id, user_uuid, conversation_id, conversation_uuid, last_access_uuid, date_part('epoch',created_at) FROM user_conversation WHERE user_id=$1;"
	queryGetUserConversationByUuid          = "SELECT id, uuid, twilio_sid, user_id, user_uuid, conversation_id, conversation_uuid, last_access_uuid, date_part('epoch',created_at) FROM user_conversation WHERE uuid=$1;"
	queryGetUserConversationForConversation = "SELECT uc.id, uc.uuid, uc.twilio_sid, uc.user_id, uc.user_uuid, uc.conversation_id, uc.conversation_uuid, uc.last_access_uuid, date_part('epoch',uc.created_at) FROM user_conversation uc JOIN conversation c ON uc.conversation_id = c.id WHERE c.uuid=$1;"
	queryUserConversationUpdateLastAccess   = "UPDATE user_conversation SET last_access_uuid=$2 WHERE uuid=$1 RETURNING uuid, last_access_uuid;"
)

var GoquDialect goqu.DialectWrapper

func init() {
	GoquDialect = goqu.Dialect("postgres")
}

type MessagingDBRepository interface {
	CreateConversation(conversation.Conversation) (*uuids.Uuid, server_message.Svr_message)
	CreateUserConversation(conversation.CreateUserConversationRequest) server_message.Svr_message
	CreateMessage(message.Message) (*uuids.Uuid, server_message.Svr_message)

	GetConversationsByUser(string) ([]conversation.ConversationAndParticipants, server_message.Svr_message)
	UpdateConversationInfo(string, conversation.ConversationInfo) (*conversation.Conversation, server_message.Svr_message)
	UpdateConversationLastMsg(string, string) (*conversation.Conversation, server_message.Svr_message)

	GetMessagesByConversation(string) ([]message.Message, server_message.Svr_message)
	UpdateMessage(string, string) (*message.Message, server_message.Svr_message)

	UserConversationUpdateLastAccess(string, string) (*conversation.UserConversation, server_message.Svr_message)
	GetConversationByUuid(string) (*conversation.Conversation, server_message.Svr_message)

	GetUserConversationByUuid(string) (*conversation.UserConversation, server_message.Svr_message)
}

type messagingDBRepository struct {
}

func GetMessagingDBRepository() MessagingDBRepository {
	return &messagingDBRepository{}
}

func (dbr *messagingDBRepository) CreateConversation(convo conversation.Conversation) (*uuids.Uuid, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryInsertConversation, convo.Type, convo.ConversationInfo.Name, convo.ConversationInfo.Description, convo.ConversationInfo.AvatarUrl, convo.TwilioSid)
	uuid := uuids.Uuid{}
	if err := row.Scan(&uuid.Uuid); err != nil {
		logger.Error("error in CreateConversation function", err)
		return nil, server_message.NewInternalError()
	}
	return &uuid, nil
}
func (dbr *messagingDBRepository) GetConversationsByUser(user_uuid string) ([]conversation.ConversationAndParticipants, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	rows, err := dbclient.Query(queryGetConversationsFromUser, user_uuid)
	if err != nil {
		logger.Error("error in GetConversationsByUser function, in the query execution", err)
		return nil, server_message.NewInternalError()
	}
	var array_convos_response []conversation.ConversationAndParticipants
	for rows.Next() {
		convo_response := conversation.ConversationAndParticipants{}
		convo := conversation.Conversation{}
		uc := conversation.UserConversation{}
		if err := rows.Scan(&convo.Id, &convo.Uuid, &convo.TwilioSid, &convo.Type, &convo.CreatedAt, &convo.LastMessageUuid, &convo.ConversationInfo.Name, &convo.ConversationInfo.Description, &convo.ConversationInfo.AvatarUrl,
			&uc.Id, &uc.Uuid, &uc.TwilioSid, &uc.UserId, &uc.UserUuid, &uc.ConversationId, &uc.ConversationUuid, &uc.LastAccessUuid, &uc.CreatedAt,
		); err != nil {
			logger.Error("error in GetConversationsByUser function, scanning rows", err)
			return nil, server_message.NewInternalError()
		}
		convo_response.Conversation = convo
		convo_response.UserConversation = uc
		array_convos_response = append(array_convos_response, convo_response)
	}
	for index, convo_response := range array_convos_response {
		ucs, aErr := dbr.getUserConversationsForConversation(convo_response.Conversation.Uuid)
		if aErr != nil {
			return nil, aErr
		}
		array_convos_response[index].Participants = ucs
	}
	if len(array_convos_response) == 0 {
		return nil, server_message.NewNotFoundError("no conversations where found in which this user partcipates")
	}
	return array_convos_response, nil
}

func (dbr *messagingDBRepository) UpdateConversationInfo(uuid string, conv_info conversation.ConversationInfo) (*conversation.Conversation, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryConversationUpdateInfo, uuid, conv_info.Name, conv_info.Description, conv_info.AvatarUrl)
	convo := conversation.Conversation{}
	if err := row.Scan(&convo.Uuid, &convo.ConversationInfo.Name, &convo.ConversationInfo.Description, &convo.ConversationInfo.AvatarUrl); err != nil {
		logger.Error("error in UpdateConversationInfo function", err)
		return nil, server_message.NewInternalError()
	}
	return &convo, nil
}
func (dbr *messagingDBRepository) UpdateConversationLastMsg(uuid string, last_message_uuid string) (*conversation.Conversation, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryConversationUpdateMsgUuid, uuid, last_message_uuid)
	convo := conversation.Conversation{}
	if err := row.Scan(&convo.Uuid, &convo.LastMessageUuid); err != nil {
		logger.Error("error in ConversationUpdateMsgUuid function", err)
		return nil, server_message.NewInternalError()
	}
	return &convo, nil
}

func (dbr *messagingDBRepository) CreateMessage(msg message.Message) (*uuids.Uuid, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryCreateMessage, msg.ConversationId, msg.ConversationUuid, msg.AuthorId, msg.AuthorUuid, msg.Text, msg.TwilioSid)
	uuid := uuids.Uuid{}
	if err := row.Scan(&uuid.Uuid); err != nil {
		logger.Error("error in CreateMessage function", err)
		return nil, server_message.NewInternalError()
	}
	return &uuid, nil
}
func (dbr *messagingDBRepository) GetMessagesByConversation(uuid string) ([]message.Message, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	rows, err := dbclient.Query(queryGetMessagesByConversationUuid, uuid)
	if err != nil {
		logger.Error("error in getmessagebyconversationid function, getting rows", err)
		return nil, server_message.NewInternalError()
	}
	msgs := []message.Message{}
	for rows.Next() {
		msg := message.Message{}
		if err := rows.Scan(&msg.Id, &msg.Uuid, &msg.TwilioSid, &msg.ConversationId, &msg.ConversationUuid, &msg.AuthorId, &msg.AuthorUuid, &msg.Text, &msg.CreatedAt, &msg.UpdatedAt); err != nil {
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

func (dbr *messagingDBRepository) UpdateMessage(uuid string, text string) (*message.Message, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryUpdateMessage, uuid, text)
	result := message.Message{}
	if err := row.Scan(&result.Uuid, &result.ConversationUuid, &result.TwilioSid, &result.Text, &result.UpdatedAt); err != nil {
		logger.Error("error at UpdateMessage", err)
		return nil, server_message.NewInternalError()
	}
	return &result, nil
}

func (dbr *messagingDBRepository) CreateUserConversation(convo conversation.CreateUserConversationRequest) server_message.Svr_message {
	dbclient := postgresql.GetSession()
	var rows [][]interface{}
	for _, uc := range convo.Ucs {
		rows = append(rows, goqu.Vals{uc.UserId, uc.UserUuid, convo.Conversation.Id, convo.Conversation.Uuid, uc.TwilioSid})
	}
	queryvals := GoquDialect.Insert("user_conversation").Cols("user_id", "user_uuid", "conversation_id", "conversation_uuid", "twilio_sid").Vals(rows...)
	query, _, err := queryvals.ToSQL()
	if err != nil {
		logger.Error("error at CreateUserConversation", err)
		return server_message.NewInternalError()
	}
	if err := dbclient.Execute(query); err != nil {
		logger.Error("error at CreateUserConversation", err)
		return server_message.NewInternalError()
	}
	return nil
}

func (dbr *messagingDBRepository) getUserConversationsForConversation(uuid string) ([]conversation.UserConversation, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	rows, err := dbclient.Query(queryGetUserConversationForConversation, uuid)
	if err != nil {
		logger.Error("error in GetUserConversationForConversation function, getting rows", err)
		return nil, server_message.NewInternalError()
	}
	ucs := []conversation.UserConversation{}
	for rows.Next() {
		uc := conversation.UserConversation{}
		if err := rows.Scan(&uc.Id, &uc.Uuid, &uc.TwilioSid, &uc.UserId, &uc.UserUuid, &uc.ConversationId, &uc.ConversationUuid, &uc.LastAccessUuid, &uc.CreatedAt); err != nil {
			logger.Error("error in GetUserConversationForConversation function, scanning rows", err)
			return nil, server_message.NewInternalError()
		}
		ucs = append(ucs, uc)
	}
	if len(ucs) == 0 {
		return nil, server_message.NewNotFoundError("no user_conversations where found")
	}
	return ucs, nil
}
func (dbr *messagingDBRepository) GetUserConversationByUuid(uuid string) (*conversation.UserConversation, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryGetUserConversationByUuid, uuid)
	uc := conversation.UserConversation{}
	if err := row.Scan(&uc.Id, &uc.Uuid, &uc.TwilioSid, &uc.UserId, &uc.UserUuid, &uc.ConversationId, &uc.ConversationUuid, &uc.LastAccessUuid, &uc.CreatedAt); err != nil {
		logger.Error("error in GetUserConversationByUuid function", err)
		return nil, server_message.NewInternalError()
	}
	return &uc, nil
}

func (dbr *messagingDBRepository) UserConversationUpdateLastAccess(uuid string, msg_uuid string) (*conversation.UserConversation, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryUserConversationUpdateLastAccess, uuid, msg_uuid)
	result := conversation.UserConversation{}
	if err := row.Scan(&result.Uuid, &result.LastAccessUuid); err != nil {
		logger.Error("error at UserConversationUpdateLastAccess", err)
		return nil, server_message.NewInternalError()
	}
	return &result, nil
}

func (dbr *messagingDBRepository) GetConversationByUuid(uuid string) (*conversation.Conversation, server_message.Svr_message) {
	dbclient := postgresql.GetSession()
	row := dbclient.QueryRow(queryGetConversationByUuid, uuid)
	convo := conversation.Conversation{}
	if err := row.Scan(&convo.Id, &convo.Uuid, &convo.TwilioSid, &convo.Type, &convo.CreatedAt, &convo.LastMessageUuid, &convo.ConversationInfo.Name, &convo.ConversationInfo.Description, &convo.ConversationInfo.AvatarUrl); err != nil {
		logger.Error("error in GetConversationByUuid function", err)
		return nil, server_message.NewInternalError()
	}
	return &convo, nil
}
