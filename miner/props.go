package miner

import (
	. "github.com/dylenfu/miner-protoactor/dao"
)

func init() {
	RegisterTransactionViewActor("transaction-view")
}
