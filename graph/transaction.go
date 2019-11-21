package graph

import (
	"github.com/gaspardpeduzzi/spring_block/display"
	"log"
	"reflect"
	"strconv"

	"github.com/gaspardpeduzzi/spring_block/data"
)


func (graph *Graph) ParseTransaction(tx data.Transaction) (newOffers []Offer) {
	//display.DisplayVerbose("====================================================================================")
	//display.DisplayVerbose("Parsing tx", tx.Hash)

	resultingOffers := make([]Offer, 1)

	for _, v := range tx.MetaData.AffectedNodes {
		created := v.CreatedNode.LedgerEntryType
		modified := v.ModifiedNode.LedgerEntryType
		deleted := v.DeletedNode.LedgerEntryType
		c := created != "" && modified == "" && deleted == ""
		m := created == "" && modified != "" && deleted == ""
		d := created == "" && modified == "" && deleted != ""

		if c {

			if v.CreatedNode.LedgerEntryType == "Offer" {
				//display.DisplayVerbose("CREATED NODE", tx.Hash)
				//display.DisplayVerbose("CREATED new offer from", v.CreatedNode.NewFields.Account, "with seq #", v.CreatedNode.NewFields.Sequence)

				test := v.CreatedNode.NewFields.TakerGets
				test1 := v.CreatedNode.NewFields.TakerPays

				//Offering
				currency, amount, issuer := CurrencyAmountAndIssuer(test)
				currency1, amount1, issuer1 := CurrencyAmountAndIssuer(test1)

				//display.DisplayVerbose("I'm paying", amount, currency, "To receive", amount1, currency1)
				var actualIssuer string
				if issuer == "" {
					actualIssuer = issuer1
				} else {
					actualIssuer = issuer
				}
				rate := amount / amount1
				//display.DisplayVerbose("RATE", rate)

				newOffer := &Offer{
					XrpTx:          tx,
					TxHash:         tx.Hash,
					Account:        v.CreatedNode.NewFields.Account,
					SequenceNumber: v.CreatedNode.NewFields.Sequence,
					Rate:           rate,
					Quantity:       amount,
					CreatorWillPay: currency,
					CreatorWillGet: currency1,
					Issuer:         actualIssuer,

				}

				graph.insertNewOffer(newOffer)
			}

		} else if m {


			if v.DeletedNode.LedgerEntryType == "Offer" {
				display.DisplayVerbose("MODIFIED NODE OFFER", v.ModifiedNode.LedgerEntryType)

			}
		} else if d {
			//display.DisplayVerbose("DELETED NODE", v.DeletedNode.LedgerEntryType)
			if v.DeletedNode.LedgerEntryType == "Offer" {

				//graph BTC ETH donne offerCreate pour obtenir ETH en payant BTC
				//[TG][TP]
				account := v.DeletedNode.FinalFields.Account
				seq := v.DeletedNode.FinalFields.Sequence
				tp := v.DeletedNode.FinalFields.TakerPays
				tg := v.DeletedNode.FinalFields.TakerGets

				currencyTP, _, _ := CurrencyAmountAndIssuer(tp)
				currencyTG, _, issuerTG := CurrencyAmountAndIssuer(tg)

				//display.DisplayVerbose("DELETED", account, seq, "ORDERBOOK", currencyTG, "/", currencyTP)

				offer := &Offer{
					XrpTx:          tx,
					TxHash:         tx.Hash,
					Account:        account,
					SequenceNumber: seq,
					Rate:           0,
					Quantity:       0,
					CreatorWillPay: currencyTG,
					CreatorWillGet: currencyTP,
					Issuer:         issuerTG,
				}

				//deletedOffers = append(deletedOffers, offer)
				if graph.Graph[offer.CreatorWillPay][offer.CreatorWillGet] != nil {
					for k, v := range graph.Graph[offer.CreatorWillPay][offer.CreatorWillGet].List {
						if v.Account == account && v.SequenceNumber == seq {
							removeOffer(graph.Graph[offer.CreatorWillPay][offer.CreatorWillGet].List,k)

						}
					}
				}

				//display.DisplayVerbose("Taker gets", currencyTG, amountTG, issuerTG)
				//display.DisplayVerbose("Taker pays", currencyTP, amountTP, issuerTP)

			}

		}
	}
	return resultingOffers

}

//display.DisplayVerbose("====================================================================================")




func typeOfTransaction(currencyTP string, currencyTG string) string {
	if currencyTP == "" && currencyTG != "" {
		return "TG"
	} else if currencyTG == "" && currencyTP != "" {
		return "TP"
	} else if currencyTP != "" && currencyTG != "" {
		return "multi"

	}
	display.DisplayVerbose("TYPES", reflect.TypeOf(currencyTP), reflect.TypeOf(currencyTG))
	return ""
}


/*

func removeAtIndex(a []Order, i int) []Order {
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.

	return a
}
*/

func CurrencyAmountAndIssuer(takerPays interface{}) (string, float64, string) {
	var weWillPayInCurrency string
	var priceToPay string
	var issuer string
	switch object := takerPays.(type) {
	case map[string]interface{}:
		//We need to pay in a given currency
		//log.Print("TAKER PAYS currency ", object["currency"]," value ",object["value"])
		weWillPayInCurrency = object["currency"].(string)
		priceToPay = object["value"].(string)
		issuer = object["issuer"].(string)
	case string:
		//We need to pay with the native currency
		//log.Print("TAKER PAYS value ", DropToXrp(float64(price)), " XRP or ", DropToPriceInUSD(price), " USD" )
		weWillPayInCurrency = "XRP"
		priceToPay = object
	default:
		log.Println("unexpected type %T", object)
	}
	weWillPayAmount, err := strconv.ParseFloat(priceToPay, 64)
	if err != nil {
		log.Println("Error decoding", err)
	}

	return weWillPayInCurrency, weWillPayAmount, issuer
}



/*
	switch typeTx {
	default:
		log.Println("unrecognized type of transaction")
	case "TG":


	case "TP":
		display.DisplayVerbose("TP")
		for _, v := range graph.NGraph[currencyTP]["XRP"].List {
			if v.Account == account && v.SequenceNumber == seq {
				display.DisplayVerbose("DELETE ME")
			}
		}
		display.DisplayVerbose("Taker pays", currencyTP, amountTP, issuerTP)

	case "multi":

		display.DisplayVerbose("MULTI")


		/*
		for _, v := range graph.NGraph[currencyTG][currencyTP].List {
			if v.Account == account && v.SequenceNumber == seq {
				display.DisplayVerbose("DELETE ME")
			}
		}
		display.DisplayVerbose("Taker pays", currencyTP, amountTP, issuerTP)
*/