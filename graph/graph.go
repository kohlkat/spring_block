package graph

import (
	"sync"

	"github.com/gaspardpeduzzi/spring_block/data"
)

type Graph struct {
	Graph map[string]map[string]TxList
	Lock  sync.Mutex
}

type TxList struct {
	List []*OfferCreate
}

type OfferCreate struct {
	xrpTx    data.Transaction
	rate     float64
	index    string
	volume   float64
	makerCur float64
	takerCur float64
}
