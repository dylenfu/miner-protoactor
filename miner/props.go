package miner

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	. "github.com/dylenfu/miner-protoactor/dao"
)

func init() {
	remote.Register("transaction-view", actor.FromProducer(NewTransactionViewActor))
}
