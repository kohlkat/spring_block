package main

import (
	"github.com/gaspardpeduzzi/spring_block/data"
	"log"
)


var maxCap = 10000000
var oldestIndex = 0

type Optimizer struct {
	Endpoint     string
	Transactions []data.Transaction
	CreateTxs    []data.Transaction
	CancelTxs    []data.Transaction



}

func NewOptimizer(endpoint string) *Optimizer {
	txs := make([]data.Transaction, maxCap)
	txsOC := make([]data.Transaction, maxCap)
	txsCancel := make([]data.Transaction, maxCap)

	return &Optimizer{endpoint,txs, txsOC, txsCancel}
}


func (lo *Optimizer) ConstructTxGraph(){

	lastIndex := data.GetLastLedgerSeq(&lo.Endpoint)

	if lastIndex > oldestIndex {
		oldestIndex = lastIndex

		txs := data.GetLedgerData(&lo.Endpoint, lastIndex)
		for _,v := range txs {
			if v.TransactionType == "OfferCreate" {
				lo.CreateTxs = append(lo.CreateTxs, v)
				log.Println(v.Hash, v.TransactionType)
			} else
			if v.TransactionType == "OfferCancel" {
				lo.CancelTxs = append(lo.CancelTxs, v)
				log.Println(v.Hash, v.TransactionType)
			}
			v.
		}
		lo.parseTransactions()
		//lo.ConstructTxGraph()
	} else {
		lo.ConstructTxGraph()
	}
}

func (lo *Optimizer) parseTransactions() {
	for _, tx := range lo.CreateTxs {
		for _, v := range tx.MetaData.AffectedNodes{
			//not sure if needed
			atomicAddress := v.CreatedNode.NewFields.Account
			tg := v.CreatedNode.NewFields.TakerGets
			tp := v.CreatedNode.NewFields.TakerPays
			log.Println(tx.Hash, tx.TakerGets, tx.TakerPays,atomicAddress, tg, tp)
		}
	}
}




