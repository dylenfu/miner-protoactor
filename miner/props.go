package miner

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/dylenfu/miner-protoactor/config"
	. "github.com/dylenfu/miner-protoactor/dao"
)

func init() {
	conf := config.LoadConfig("")

	NewDb(&conf.Mysql)

	remote.Register("transaction-view", actor.FromProducer(NewTransactionViewActor))
	remote.Register("transfer-view", actor.FromProducer(NewTransferViewActor))
}
