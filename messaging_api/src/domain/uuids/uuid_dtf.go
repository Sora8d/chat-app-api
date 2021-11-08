package uuids

import (
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
)

func (uuid *Uuid) Poblate(direction_out bool, pbuuid *pb.Uuid) {
	if direction_out {
		pbuuid.Uuid = uuid.Uuid
	} else {
		uuid.Uuid = pbuuid.Uuid
	}
}
