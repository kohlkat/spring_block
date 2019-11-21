package main

import (
	"github.com/gaspardpeduzzi/spring_block/display"
	"log"
	"sync"
	"github.com/gaspardpeduzzi/spring_block/data"
	"github.com/gaspardpeduzzi/spring_block/graph"
)


var oldestIndex = 0

type Optimizer struct {
	Endpoint string
	Graph    graph.Graph
	Channel  chan int
}

func NewOptimizer(endpoint string, c chan int) *Optimizer {
	graph := graph.Graph{
		NGraph: make(map[string]map[string]*graph.OrderBook),
		AccountRoots: make(map[string]map[int]*graph.Offer),
		Lock:  sync.RWMutex{},
	}
	return &Optimizer{endpoint, graph, c}
}

func (lo *Optimizer) NConstructTxGraph() {
	lastIndex := data.GetLastLedgerSeq(&lo.Endpoint)
	if lastIndex > oldestIndex {
		display.DisplayVerbose("New block index: ", lastIndex)
		oldestIndex = lastIndex
		txs := data.GetLedgerData(&lo.Endpoint, lastIndex)
		var tmpCreate []data.Transaction
		for _, tx := range txs {
			if tx.TransactionType == "OfferCreate" || tx.TransactionType == "OfferCancel" || tx.TransactionType == "Payment" {
				tmpCreate = append(tmpCreate, tx)
			}
		}
		lo.ParseOfferCreateTransactions(tmpCreate)
		//lo.Channel <- 1
		lo.NConstructTxGraph()
	}
	lo.NConstructTxGraph()
}



func (lo *Optimizer) ParseOfferCreateTransactions(transactions []data.Transaction) {
	log.Println("ADDED", len(transactions), " OfferCreate transaction(s)")
	for _, tx := range transactions {
		lo.Graph.ParseTransaction(tx)

	}
	lo.Graph.SortGraphWithTxs()
}



