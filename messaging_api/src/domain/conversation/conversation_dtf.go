package conversation

import (
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
)

func (c *Conversation) Poblate(direction_out bool, cpb *pb.Conversation) {
	if direction_out {
		Uuid := pb.Uuid{Uuid: c.Uuid}
		cpb.Uuid = &Uuid
		cpb.CreatedAt = c.CreatedAt
		lastmsg := pb.Message{}
		c.LastMessage.Poblate(direction_out, &lastmsg)
		cpb.LastMsg = &lastmsg
		cpb.Type = c.Type
		cpb.Description = c.Description
		cpb.AvatarUrl = c.AvatarUrl
		cpb.Name = c.Name
	} else {
		if cpb.Uuid != nil {
			c.Uuid = cpb.Uuid.Uuid
		}
		if cpb.LastMsg != nil && cpb.LastMsg.Uuid != nil {
			c.LastMessage.Poblate(direction_out, cpb.LastMsg)
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
		Uuid := pb.Uuid{Uuid: uc.Uuid}
		cpb.Uuid = &Uuid
		userUuid := pb.Uuid{Uuid: uc.UserUuid}
		cpb.UserUuid = &userUuid
		last_accessUuid := pb.Uuid{Uuid: uc.LastAccessUuid}
		cpb.LastAccessUuid = &last_accessUuid
		cpb.CreatedAt = uc.CreatedAt
	} else {
		if cpb.Uuid != nil {
			uc.Uuid = cpb.Uuid.Uuid
		}
		uc.UserUuid = cpb.UserUuid.Uuid
		if cpb.LastAccessUuid != nil {
			uc.LastAccessUuid = cpb.LastAccessUuid.Uuid
		}
		uc.CreatedAt = cpb.CreatedAt
	}
}

//These return slices because slices are already pointers, and having a pointer to a pointer seems messy.
func (ucs *UserConversationSlice) Poblate(direction_out bool, cpb []*pb.UserConversation) []*pb.UserConversation {
	if direction_out {
		var object_to_return []*pb.UserConversation
		for _, participant := range *ucs {
			pb_participant := pb.UserConversation{}
			participant.Poblate(true, &pb_participant)
			object_to_return = append(object_to_return, &pb_participant)
		}
		return object_to_return
	} else {
		for _, pb_participant := range cpb {
			participant := UserConversation{}
			participant.Poblate(false, pb_participant)
			*ucs = append(*ucs, participant)
		}
		return nil
	}
}
func (cps *ConversationAndParticipantsSlice) Poblate(pbcps []*pb.ConversationAndParticipants) []*pb.ConversationAndParticipants {
	var object_to_return []*pb.ConversationAndParticipants
	for _, content := range *cps {
		var new_pb_convo pb.ConversationAndParticipants
		content.Poblate(true, &new_pb_convo)
		object_to_return = append(object_to_return, &new_pb_convo)
	}
	return object_to_return
}

func (cr *ConversationAndParticipants) Poblate(direction_out bool, pbacr *pb.ConversationAndParticipants) {
	if direction_out {
		pb_convo := pb.Conversation{}
		pb_user_conversation := pb.UserConversation{}

		cr.Conversation.Poblate(direction_out, &pb_convo)
		cr.UserConversation.Poblate(direction_out, &pb_user_conversation)
		pbacr.Conversation = &pb_convo
		pbacr.UserConversation = &pb_user_conversation
		pbacr.Participants = cr.Participants.Poblate(true, nil)
		return
	} else {
		cr.Conversation.Poblate(direction_out, pbacr.Conversation)
		if pbacr.UserConversation != nil {
			cr.UserConversation.Poblate(direction_out, pbacr.UserConversation)
		}
		cr.Participants.Poblate(false, pbacr.Participants)
		return
	}
}

func (cucr *CreateUserConversationRequest) Poblate(cpb []*pb.UserConversation) {
	cucr.UserConversationSlice.Poblate(false, cpb)
}
