package conversation

import (
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
)

func (c *Conversation) Poblate(direction_out bool, cpb *pb.Conversation) {
	if direction_out {
		Uuid := pb.Uuid{Uuid: c.Uuid}
		cpb.Uuid = &Uuid
		cpb.CreatedAt = c.CreatedAt
		lastmsgUuid := pb.Uuid{Uuid: c.LastMessageUuid}
		cpb.LastMsgUuid = &lastmsgUuid
		cpb.Type = c.Type
		cpb.Description = c.Description
		cpb.AvatarUrl = c.AvatarUrl
		cpb.Name = c.Name
	} else {
		if cpb.Uuid != nil {
			c.Uuid = cpb.Uuid.Uuid
		}
		if cpb.LastMsgUuid != nil {
			c.LastMessageUuid = cpb.LastMsgUuid.Uuid
		}
		c.CreatedAt = cpb.CreatedAt
		c.Type = cpb.Type
		c.Description = cpb.Description
		c.AvatarUrl = cpb.AvatarUrl
		c.Name = cpb.Name
	}
}

func (uc *UserConversation) Poblate(direction_out bool, cpb *pb.UserConversation) {
	if direction_out {
		uuid := pb.Uuid{Uuid: uc.Uuid}
		cpb.Uuid = &uuid
		userUuid := pb.Uuid{Uuid: uc.UserUuid}
		cpb.UserUuid = &userUuid
		conversationUuid := pb.Uuid{Uuid: uc.ConversationUuid}
		cpb.ConversationUuid = &conversationUuid
		last_accessUuid := pb.Uuid{Uuid: uc.LastAccessUuid}
		cpb.LastAccessUuid = &last_accessUuid
		cpb.CreatedAt = uc.CreatedAt
	} else {
		if cpb.Uuid != nil {
			uc.Uuid = cpb.Uuid.Uuid
		}
		uc.UserUuid = cpb.UserUuid.Uuid
		uc.ConversationUuid = cpb.ConversationUuid.Uuid
		if cpb.LastAccessUuid != nil {
			uc.LastAccessUuid = cpb.LastAccessUuid.Uuid
		}
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
