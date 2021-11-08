package message

import (
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
)

func (ms *Message) Poblate(direction_out bool, mpb *pb.Message) {
	if direction_out {
		Uuid := pb.Uuid{Uuid: ms.Uuid}
		mpb.Uuid = &Uuid
		AuthorUuid := pb.Uuid{Uuid: ms.AuthorUuid}
		mpb.AuthorUuid = &AuthorUuid
		ConversationUuid := pb.Uuid{Uuid: ms.ConversationUuid}
		mpb.ConversationUuid = &ConversationUuid
		mpb.Text = ms.Text
		mpb.CreatedAt = ms.CreatedAt
		mpb.UpdatedAt = ms.UpdatedAt
	} else {
		if mpb.Uuid != nil {
			ms.Uuid = mpb.Uuid.Uuid
		}
		if mpb.AuthorUuid != nil {
			ms.AuthorUuid = mpb.AuthorUuid.Uuid
		}
		if mpb.ConversationUuid != nil {
			ms.ConversationUuid = mpb.ConversationUuid.Uuid
		}
		ms.Text = mpb.Text
		ms.CreatedAt = mpb.CreatedAt
		ms.UpdatedAt = mpb.UpdatedAt
	}
}
