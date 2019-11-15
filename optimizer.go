package main

import (
	// "log"
	"sync"

	"github.com/gaspardpeduzzi/spring_block/data"
	"github.com/gaspardpeduzzi/spring_block/graph"
)

var smallCap = 1000
var maxCap = 1000000
var oldestIndex = 0

type Optimizer struct {
	Endpoint     string
	Transactions []data.Transaction
	CreateTxs    []data.Transaction
	CancelTxs    []data.Transaction
	Graph        graph.Graph
	Channel      chan int
}

func NewOptimizer(endpoint string, c chan int) *Optimizer {
	txs := make([]data.Transaction, maxCap)
	txsOC := make([]data.Transaction, maxCap)
	txsCancel := make([]data.Transaction, maxCap)
	graph := graph.Graph{
		Graph: make(map[string]map[string]*graph.TxList),
		Lock:  sync.RWMutex{},
	}
	return &Optimizer{endpoint, txs, txsOC, txsCancel, graph, c}
}

func (lo *Optimizer) ConstructTxGraph() {

	lastIndex := data.GetLastLedgerSeq(&lo.Endpoint)

	if lastIndex > oldestIndex {
		oldestIndex = lastIndex
		txs := data.GetLedgerData(&lo.Endpoint, lastIndex)
		tmp := make([]data.Transaction, smallCap)

		for _, v := range txs {
			if v.TransactionType == "OfferCreate" {
				lo.CreateTxs = append(lo.CreateTxs, v)
				tmp = append(tmp, v)
				// log.Println(v.Hash, v.TransactionType)
			} else if v.TransactionType == "OfferCancel" {
				lo.CancelTxs = append(lo.CancelTxs, v)
				// log.Println(v.Hash, v.TransactionType)
			}
		}
		lo.parseTransactions(tmp)
		lo.Channel <- 1
		lo.ConstructTxGraph()
	} else {
		lo.ConstructTxGraph()
	}
}

func (lo *Optimizer) parseTransactions(transactions []data.Transaction) {
	//log.Println("============================================================")
	for _, tx := range transactions {
		lo.Graph.AddOffers(tx)
	}
	lo.Graph.SortGraphWithTxs()
	//lo.checkSorting()

}

// func (lo *Optimizer) checkSorting() {
// 	log.Println("CHECK SORT")
// 	for i, v := range lo.Graph.Graph {
// 		for k, _ := range v {
// 			listOfTxs := lo.Graph.Graph[i][k].List
// 			for _, b := range listOfTxs {
// 				log.Println(b.Rate)
// 			}
// 		}

// 	}
//}
