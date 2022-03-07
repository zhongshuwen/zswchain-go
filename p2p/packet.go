package p2p

import (
	"github.com/zhongshuwen/zswchain-go"
)

type Envelope struct {
	Sender   *Peer
	Receiver *Peer
	Packet   *zsw.Packet `json:"envelope"`
}

func NewEnvelope(sender *Peer, receiver *Peer, packet *zsw.Packet) *Envelope {
	return &Envelope{
		Sender:   sender,
		Receiver: receiver,
		Packet:   packet,
	}
}
