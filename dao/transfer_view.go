package dao

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/dylenfu/miner-protoactor/messages"
)

type Transfer struct {
	ID          int    `gorm:"column:id;primary_key""`
	From        string `gorm:"column:tx_from";type:varchar(42)`
	To          string `gorm:"column:tx_to";type:varchar(42)`
	TxHash      string `gorm:"column:tx_hash";type:varchar(82)`
	BlockHash   string `gorm:"column:block_hash";type:varchar(82)`
	BlockNumber int    `gorm:"column:block_number"`
	Value       []byte `gorm:"column:tx_value";type:varchar(30)`
	Failed      bool   `gorm:"column:tx_failed"`
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
