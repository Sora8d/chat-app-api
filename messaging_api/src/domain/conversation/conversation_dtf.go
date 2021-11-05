package conversation

import (
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
)

func (c *Conversation) Poblate(direction_out bool, cpb *pb.Conversation) {
	if direction_out {
		cpb.Uuid.Uuid = c.Uuid
		cpb.CreatedAt = c.CreatedAt
		cpb.LastMsgUuid = c.LastMessageUuid
		cpb.Type = c.Type
		cpb.Description = c.Description
		cpb.AvatarUrl = c.AvatarUrl
		cpb.Name = c.Name
	} else {
		c.Uuid = cpb.Uuid.Uuid
		c.CreatedAt = cpb.CreatedAt
		c.LastMessageUuid = cpb.LastMsgUuid
		c.Type = cpb.Type
		c.Description = cpb.Description
		c.AvatarUrl = cpb.AvatarUrl
		c.Name = cpb.Name
	}
}

func (uc *UserConversation) Poblate(direction_out bool, cpb *pb.UserConversation) {
	if direction_out {
		cpb.Uuid.Uuid = uc.Uuid
		cpb.UserUuid.Uuid = uc.UserUuid
		cpb.ConversationUuid.Uuid = uc.ConversationUuid
		cpb.LastAccessUuid.Uuid = uc.LastAccessUuid
		cpb.CreatedAt = uc.CreatedAt
	} else {
		uc.Uuid = cpb.Uuid.Uuid
		uc.UserUuid = cpb.UserUuid.Uuid
		uc.ConversationUuid = cpb.ConversationUuid.Uuid
		uc.LastAccessUuid = cpb.LastAccessUuid.Uuid
		uc.CreatedAt = cpb.CreatedAt
	}
}

func (cr *ConversationResponse) Poblate(direction_out bool, pbacr *pb.ConversationAndParticipants) {
	if direction_out {
		pb_convo := pb.Conversation{}
		pb_user_conversation := pb.UserConversation{}
		pb_participants := []*pb.UserConversation{}

		cr.Conversation.Poblate(direction_out, &pb_convo)
		cr.UserConversation.Poblate(direction_out, &pb_user_conversation)
		for _, participant := range cr.Participants {
			pb_participant := pb.UserConversation{}
			participant.Poblate(direction_out, &pb_participant)
			pb_participants = append(pb_participants, &pb_participant)
		}
		pbacr.Conversation = &pb_convo
		pbacr.UserConversation = &pb_user_conversation
		pbacr.Participants = pb_participants
		return
	} else {
		cr.Conversation.Poblate(direction_out, pbacr.Conversation)
		cr.UserConversation.Poblate(direction_out, pbacr.UserConversation)
		for _, pb_participant := range pbacr.Participants {
			participant := UserConversation{}
			participant.Poblate(direction_out, pb_participant)
			cr.Participants = append(cr.Participants, participant)
		}
		return
	}
}
