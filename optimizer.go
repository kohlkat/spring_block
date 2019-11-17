package main

import (
	"github.com/gaspardpeduzzi/spring_block/display"
	"log"
	"sync"

	"github.com/gaspardpeduzzi/spring_block/data"
	"github.com/gaspardpeduzzi/spring_block/graph"
)

var smallCap = 1000
var maxCap = 1000000
var oldestIndex = 0

type Optimizer struct {
	Endpoint string
	Graph    graph.Graph
	Channel  chan int
}

func NewOptimizer(endpoint string, c chan int) *Optimizer {
	graph := graph.Graph{
		Graph: make(map[string]map[string]*graph.TxList),
		ActiveOffers: make(map[string]*graph.Offer),
		Lock:  sync.RWMutex{},
	}
	return &Optimizer{endpoint, graph, c}
}

func (lo *Optimizer) ConstructTxGraph() {

	lastIndex := data.GetLastLedgerSeq(&lo.Endpoint)

	if lastIndex > oldestIndex {
		oldestIndex = lastIndex
		txs := data.GetLedgerData(&lo.Endpoint, lastIndex)
		var tmpCreate []data.Transaction
		var tmpCancel []data.Transaction
		var tmpPay []data.Transaction

		for _, v := range txs {
			if v.TransactionType == "OfferCreate" {
				//display.DisplayVerbose(v.Hash, v.TransactionType)
				tmpCreate = append(tmpCreate, v)
			} else if v.TransactionType == "OfferCancel" {
				tmpCancel = append(tmpCancel, v)
			} else if v.TransactionType == "Payment" {
				tmpPay = append(tmpPay, v)
			}
		}
		lo.parseTransactions(tmpCreate)
		if len(tmpCancel) != 0 {
			lo.parseTransactionsCancel(tmpCancel)
		}
		lo.parsePay(tmpPay)

		lo.Channel <- 1
		lo.ConstructTxGraph()
	}
	lo.ConstructTxGraph()

}


func (lo *Optimizer) parsePay(transactions []data.Transaction){
	log.Println("ADDED", len(transactions), " Payment transaction(s)")

}

func (lo *Optimizer) parseTransactionsCancel(transactions []data.Transaction) {
	display.DisplayVerbose("ADDED", len(transactions), " OfferCancel transaction(s)")

	deleted := make([]string,1)


	for _, tx := range transactions {
		for _, v := range tx.MetaData.AffectedNodes {
			//log.Println("====================================================================================")
			//display.DisplayVerbose("MODIFIED", v.ModifiedNode.LedgerIndex, "CREATE", v.CreatedNode.LedgerIndex, "DELETED", v.DeletedNode.LedgerIndex)
			//display.DisplayVerbose(v.DeletedNode.FinalFields.PreviousTxnID)
			if v.DeletedNode.FinalFields.PreviousTxnID != "" {
				deleted = append(deleted, v.DeletedNode.FinalFields.PreviousTxnID)
			}
			//log.Println("====================================================================================")
		}
	}
	if deleted[0] != "" {
		display.DisplayVerbose("DELETING OFFERS")
		lo.Graph.DeleteOffers(deleted)
	}

}

func (lo *Optimizer) parseTransactions(transactions []data.Transaction) {
	display.DisplayVerbose("ADDED", len(transactions), " OfferCreate transaction(s)")
	for _, tx := range transactions {
		lo.Graph.AddOffers(tx)
	}
	lo.Graph.SortGraphWithTxs()
}
