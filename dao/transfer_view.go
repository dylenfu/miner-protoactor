package dao

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/dylenfu/miner-protoactor/messages"
)

type Transfer struct {
	From        string
	To          string
	TxHash      string
	BlockHash   string
	BlockNumber int
	Value       []byte
	Failed      bool
}

func NewTransferViewActor() actor.Actor {
	return &TransferViewActor{}
}

type TransferViewActor struct{}

func (*TransferViewActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *messages.Transfer:
		println(msg.TxHash)
	}
}
