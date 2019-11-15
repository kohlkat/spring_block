package main

import (
	"flag"
	"strconv"

	"github.com/gaspardpeduzzi/spring_block/data"
	display "github.com/gaspardpeduzzi/spring_block/display_cli"
)

func main() {
	//go s.LaunchServer()
	var addr = flag.String("addr", "s1.ripple.com:51233", "http service address")
	var verb = flag.Bool("verb", true, "Display more information")
	flag.Parse()
	display.VERBOSE = *verb

	// Init Display
	display.Init()
	display.DisplayVerbose("Checking last sequence number ", strconv.Itoa(data.GetLastLedgerSeq(addr)))

	c := make(chan int)
	liquidOptimizer := NewOptimizer(*addr, c)
	go liquidOptimizer.ConstructTxGraph()

	for {

		<-c
		all_offers, cycle := liquidOptimizer.Graph.GetProfitableOffers()
		if all_offers != nil {
			display.DisplayVerbose("Found profitable cycle:", cycle)
			for i, offers := range all_offers {
				for _, offer := range offers {
					display.DisplayVerbose(cycle[i], "->", cycle[(i+1)%len(cycle)], offer.Rate, offer.Hash)
				}
			}
			return
		}
	}

}
