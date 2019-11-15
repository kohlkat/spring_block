package graph

import (
	"github.com/gaspardpeduzzi/spring_block/data"
	"log"
	"strconv"
)

func (graph *Graph) AddOffers(tx data.Transaction) {
	mapTx := make(map[string]*Offer)
	var weWillPay string
	var weWillGet string


	for index, _ := range tx.MetaData.AffectedNodes {

		priceToPay := ""
		priceWillGet := ""

		takerGets := tx.TakerGets
		takerPays := tx.TakerPays

		switch object := takerPays.(type) {

		case map[string]interface{}:
			//We need to pay in a given currency
			//log.Print("TAKER PAYS currency ", object["currency"]," value ",object["value"])
			weWillPay = object["currency"].(string)
			priceToPay = object["value"].(string)
		case string:
			//We need to pay with the native currency
			//log.Print("TAKER PAYS value ", DropToXrp(float64(price)), " XRP or ", DropToPriceInUSD(price), " USD" )
			weWillPay = "XRP"
			priceToPay = object
		default:
			log.Println("unexpected type %T", object)
			log.Println(index)

		}

		switch objectTG := takerGets.(type) {

		case map[string]interface{}:
			//We will get a given currency
			weWillGet = objectTG["currency"].(string)
			//log.Print("TAKER GETS currency ", objectTG["currency"], " value ", objectTG["value"])
			priceWillGet = objectTG["value"].(string)
		case string:
			//We will get the native currency
			weWillGet = "XRP"
			priceWillGet = objectTG
			//log.Print("TAKER GETS value ", DropToXrp(float64(price)), " XRP or ", DropToPriceInUSD(price), " USD" )
		default:
			log.Println("unexpected type %T", objectTG)
			log.Println(index)

		}

		WillGet, err := strconv.ParseFloat(priceWillGet,64)
		if err != nil {
			log.Println("Error decoding",err)

		}
		WillPay, err := strconv.ParseFloat(priceToPay, 64)
		if err != nil {
			log.Println("Error decoding",err)

		}



		rate := WillGet/WillPay
		vol := rate*WillPay

		offer := Offer{
			XrpTx:    tx,
			Hash:     tx.Hash,
			Rate:     rate,
			Volume:   vol,
			Pay: weWillPay,
			Get: weWillGet,
		}

		mapTx[tx.Hash] = &offer

	}

	for _, v := range mapTx {
		// log.Println("ADDING", tx.Hash)
		graph.addNewOffer(v.Pay, v.Get, v)
	}
}
