package dao

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/dylenfu/miner-protoactor/messages"
)

type Transaction struct {
	From   string
	To     string
	Hash   string
	Value  []byte
	Failed bool
}

func NewTransactionViewActor() actor.Actor {
	return &TransactionViewActor{}
}

type TransactionViewActor struct{}

func (*TransactionViewActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *messages.Transaction:
		println(msg.Hash)
	}
}
