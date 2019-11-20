package main

import (
	"flag"
	display "github.com/gaspardpeduzzi/spring_block/display"
	"reflect"
	// data "github.com/gaspardpeduzzi/spring_block/data"
	"log"
	"fmt"
	"os/exec"
	"encoding/json"
)

func main() {

	//go s.LaunchServer()

	var addr = flag.String("addr", "s1.ripple.com:51233", "http service address")
	var verb = flag.Bool("verb", false, "Display more information")
	flag.Parse()
	display.VERBOSE = *verb
	display.Init()

	c := make(chan int)
	liquidOptimizer := NewOptimizer(*addr, c)
	go liquidOptimizer.ConstructTxGraph()

		for {
			display.DisplayVerbose("waiting for next block...")
			<-c
			allOffers, cycle := liquidOptimizer.Graph.GetProfitableOffers()
			//returns map[int][]Offer, []string

			if allOffers != nil {
				//Should never be displayed in verbose mode :)
				log.Println("Found profitable cycle:", cycle)
				log.Println("====================================================================================")
				for i, offers := range allOffers {
					for _, offer := range offers {
						log.Println(cycle[i], "->", cycle[(i+1)%len(cycle)], offer.Rate, "OfferCreate Hash:", offer.Hash, "Volume:", offer.Volume)
						display.DisplayVerbose(i, cycle[i], "->", cycle[(i+1)%len(cycle)])

						// for _, node := range offer.XrpTx.MetaData.AffectedNodes {
						//
						//
						// }


					}
				}
				log.Println("====================================================================================")
				//return
			}
		}

}

// func tmp(node interface{}) {
// 	if node.Type().String() == "data.CreatedNode" {
// 		// 	log.Println("node", node)
// 		// }
// 		log.Println("node_type", node)
// 	}
// }

// type CreatedNode data.CreatedNode
func getType(a interface{}) string {
  val := reflect.Indirect(reflect.ValueOf(a))
	return val.Field(0).Type().Name()
}

type TmpOrder struct {
	Hash string
	Rate float64
	Quantity float64
	Taker struct {
		Currency string
		Value float64
		Issuer string
	}
}

func submit_transaction2(takerPays TmpOrder) {
	out, err := json.Marshal(takerPays)
	if err != nil {
			panic(err)
	}
	cmd := fmt.Sprintf("./submit.sh %s %s", takerPays.Hash, string(out))
	log.Println("cmd", cmd)
	out, err = exec.Command(cmd).Output()
	log.Println("out, err", out, err)
}

// func submit_transaction(node interface{}, hash string) {
//
// 	log.Println("node", node)
//
// 	switch getType(node) {
// 	case "CreatedNode":
//
// 		field, ok := node.(data.CreatedNode)
//
// 		log.Println("field, ok", field, ok)
//
// 		// createdNode,_ := reflect.ValueOf(node).Interface().(data.CreatedNode)
// 		// log.Println("CreatedNode", createdNode)
//
// 	case "ModifiedNode":
// 		log.Println("ModifiedNode")
// 	case "DeletedNode":
// 		log.Println("DeletedNode")
// 	default:
// 		log.Println("default", getType(node))
// 		// tmp(node)
// 	}
//
// }
