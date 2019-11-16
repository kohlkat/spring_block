package main

import (
	"github.com/gaspardpeduzzi/spring_block/display"
	"sync"

	"github.com/gaspardpeduzzi/spring_block/data"
	"github.com/gaspardpeduzzi/spring_block/graph"
)

var smallCap = 1000
var maxCap = 1000000
var oldestIndex = 0

type Optimizer struct {
	Endpoint     string
	Graph        graph.Graph
	Channel      chan int
}

func NewOptimizer(endpoint string, c chan int) *Optimizer {
	graph := graph.Graph{
		Graph: make(map[string]map[string]*graph.TxList),
		Lock:  sync.RWMutex{},
	}
	return &Optimizer{endpoint, graph, c}
}

func (lo *Optimizer) ConstructTxGraph() {

	lastIndex := data.GetLastLedgerSeq(&lo.Endpoint)

	if lastIndex > oldestIndex {
		oldestIndex = lastIndex
		txs := data.GetLedgerData(&lo.Endpoint, lastIndex)
		var tmp []data.Transaction

		for _, v := range txs {
			if v.TransactionType == "OfferCreate" {
				//display.DisplayVerbose(v.Hash, v.TransactionType)
				tmp = append(tmp, v)
			}
		}
		lo.parseTransactions(tmp)
		lo.Channel <- 1
		lo.ConstructTxGraph()
	}
	lo.ConstructTxGraph()

}

func (lo *Optimizer) parseTransactions(transactions []data.Transaction) {
	display.DisplayVerbose("ADDED", len(transactions), "new transactions")
	for _, tx := range transactions {
		lo.Graph.AddOffers(tx)
	}
	lo.Graph.SortGraphWithTxs()
}


