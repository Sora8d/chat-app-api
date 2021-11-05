package message

import (
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
)

func (ms *Message) Poblate(direction_out bool, mpb *pb.Message) {
	if direction_out {
		mpb.Uuid.Uuid = ms.Uuid
		mpb.AuthorUuid.Uuid = ms.AuthorUuid
		mpb.ConversationUuid.Uuid = ms.ConversationUuid
		mpb.Text = ms.Text
		mpb.CreatedAt = ms.CreatedAt
		mpb.UpdatedAt = ms.UpdatedAt
	} else {
		if mpb.Uuid != nil {
			ms.Uuid = mpb.Uuid.Uuid
		}
		ms.AuthorUuid = mpb.AuthorUuid.Uuid
		ms.ConversationUuid = mpb.ConversationUuid.Uuid
		ms.Text = mpb.Text
		ms.CreatedAt = mpb.CreatedAt
		ms.UpdatedAt = mpb.UpdatedAt
	}
}
