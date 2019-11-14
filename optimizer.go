package main

import (
	"github.com/gaspardpeduzzi/spring_block/data"
	"log"
	"strconv"
)


var maxCap = 1000000
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
		}

		lo.parseTransactions()
		//lo.ConstructTxGraph()
	} else {
		log.Println("SIZE TRANSACTIONS", len(lo.CancelTxs)+len(lo.CreateTxs))
		lo.ConstructTxGraph()
	}
}




func (lo *Optimizer) parseTransactions() {
	log.Println("parsing..")

	for _, tx := range lo.CreateTxs {

		for index, _ := range tx.MetaData.AffectedNodes {

			takerGets := tx.TakerGets
			takerPays := tx.TakerPays

			//log.Println(tx.TakerPays, tx.TakerGets)
			log.Println("TX at", tx.Hash)

			switch object := takerPays.(type) {
			default:
				log.Println("unexpected type %T", object)
				log.Println(index)
			case map[string]interface{}:
				//log.Print("MAP object ")
				log.Print("TAKER PAYS currency ", object["currency"]," value ",object["value"])
			case string:
				//log.Print("STRING object")
				price, err := strconv.Atoi(object)
				if err != nil {
					log.Println(err)
				}
				log.Print("TAKER PAYS value ", dropToXrp(float64(price)), " XRP or ", dropToPriceInUSD(price) )
				//log.Print("TAKER PAYS value ", price, " XRP or ", price*1/4, " USD")
			}

			switch objectTG := takerGets.(type) {
			default:
				log.Println("unexpected type %T", objectTG)
				log.Println(index)
			case map[string]interface{}:
				//log.Print("MAP object ")
				log.Print("TAKER GETS currency ", objectTG["currency"], " value ", objectTG["value"])
			case string:
				//log.Print("STRING object")
				price, err := strconv.Atoi(objectTG)
				if err != nil {
					log.Println(err)
				}
				log.Print("TAKER GETS value ", dropToXrp(float64(price)), " XRP or ", dropToPriceInUSD(price), " USD" )
			}
			//log.Print("\n")
			log.Println("===========================================================")


		}
	}
}




