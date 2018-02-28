package dao

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/dylenfu/miner-protoactor/messages"
)

type Transaction struct {
	ID     int    `gorm:"column:id;primary_key""`
	From   string `gorm:"column:tx_from;type:varchar(42)"`
	To     string `gorm:"column:tx_to;type:varchar(42)"`
	Hash   string `gorm:"column:tx_hash;type:varchar(82)"`
	Value  []byte `gorm:"column:tx_value;type:varchar(30)"`
	Failed bool   `gorm:"column:tx_failed"`
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
