package main

import (
	"flag"
	"fmt"
	display "github.com/gaspardpeduzzi/spring_block/display"
	server "github.com/gaspardpeduzzi/spring_block/server"
)

func main() {

	var addr = flag.String("addr", "s1.ripple.com:51233", "http service address")
	var verb = flag.Bool("verb", false, "Display more information")
	var analysis = flag.Bool("analysis",false, "Analyse statistics")

	flag.Parse()
	display.VERBOSE = *verb
	display.ANALYSIS = *analysis
	display.Init()

	c := make(chan int)

	liquidOptimizer := NewOptimizer(*addr, c)

	go liquidOptimizer.ConstructTxGraph()

	go server.LaunchServer()

	// Search for arbitrage opportunities and store them
	for {
		display.DisplayVerbose("waiting for next block...")
		<-c
		allOffers, cycle := liquidOptimizer.Graph.GetProfitableOffers()
		seq_nb := 1
		server.AccountsNumber = len(liquidOptimizer.Graph.Clients)



		if allOffers != nil {
			//Should never be displayed in verbose mode :)
			server.ArbitrageOffersDB = append(server.ArbitrageOffersDB, &server.ArbitrageOpportunities{Pair: cycle, Offers: make([]*server.OfferSummary, 0)})
			fmt.Println("Found profitable cycle:", cycle)
			fmt.Println("====================================================================================")
			for i, offers := range allOffers {
				for _, offer := range offers {
					fmt.Println(cycle[i], "->", cycle[(i+1)%len(cycle)], offer.Rate, "OfferCreate Hash:", offer.TxHash, "Volume:", offer.Quantity)
					//offer.Submit_Transaction(seq_nb)
					seq_nb = seq_nb + 1

				}
			}
			server.ArbitrageOffersDB = append(server.ArbitrageOffersDB, &server.ArbitrageOpportunities{Pair: cycle, Offers: make([]*server.OfferSummary, 0)})
			fmt.Println("====================================================================================")
		}
	}

}
