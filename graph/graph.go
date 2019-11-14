package graph

import (
	"github.com/gaspardpeduzzi/spring_block/data"
	"sync"
)

type Graph struct {
	Graph map[string]map[string]Tx
	Lock sync.Mutex

}

type Tx struct {
	Txs []*data.Transaction

}
