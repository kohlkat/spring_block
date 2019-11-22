package main

import (
	"github.com/gaspardpeduzzi/spring_block/data"
	"github.com/gaspardpeduzzi/spring_block/display"
	"github.com/gaspardpeduzzi/spring_block/graph"
	"sync"
)

var oldestIndex = 0

type Optimizer struct {
	Endpoint string
	Graph    graph.Graph
	Channel  chan int
}

func NewOptimizer(endpoint string, c chan int) *Optimizer {
	graph := graph.Graph{
		Graph: make(map[string]map[string]*graph.OrderBook),
		AccountRoots: make(map[string]map[int]*graph.Offer),
		Issuers: make(map[string]bool),
		Clients: make(map[string]bool),
		AccountRoot: make(map[string][]string),
		Lock:  sync.RWMutex{},
	}
	return &Optimizer{endpoint, graph, c}
}

func (lo *Optimizer) ConstructTxGraph() {
	lastIndex := data.GetLastLedgerSeq(&lo.Endpoint)
	if lastIndex > oldestIndex {
		display.DisplayVerbose("New block index: ", lastIndex)
		oldestIndex = lastIndex
		txs := data.GetLedgerData(&lo.Endpoint, lastIndex)
		var tmpCreate []data.Transaction
		var tmpPayment []data.Transaction
		for _, tx := range txs {
			if tx.TransactionType == "OfferCreate" || tx.TransactionType == "OfferCancel"{
				tmpCreate = append(tmpCreate, tx)
			} else if tx.TransactionType == "Payment" {
				tmpPayment = append(tmpPayment, tx)
			}
		}
		lo.ParseOfferCreateTransactions(tmpCreate)
		lo.ParsePaymentTransactions(tmpPayment)
		lo.Channel <- 1
		lo.ConstructTxGraph()
	}
	lo.ConstructTxGraph()
}

func (lo *Optimizer) ParseOfferCreateTransactions(transactions []data.Transaction) {
	display.DisplayVerbose("ADDED", len(transactions), "OfferCreate and OfferCancel transaction(s)")
	for _, tx := range transactions {
		lo.Graph.ParseTransaction(tx)
	}
	lo.Graph.SortGraphWithTxs()
	// lo.Channel <- 1
}

func (lo *Optimizer) ParsePaymentTransactions(transactions []data.Transaction) {
		display.DisplayVerbose("ADDED", len(transactions), "new payment transaction(s)")
		for _, tx := range transactions {
			lo.Graph.PaymentTransactionParse(tx)
		}
		lo.Graph.SortGraphWithTxs()
// lo.Channel <- 1
}
