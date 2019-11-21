package main

import (
	"flag"

	display "github.com/gaspardpeduzzi/spring_block/display"
	server "github.com/gaspardpeduzzi/spring_block/server"
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
	liquidOptimizer.NConstructTxGraph()

	server.LaunchServer()

	// Search for arbitrage opportunities and store them
	for {
		<-c
		allOffers, cycle := liquidOptimizer.Graph.GetProfitableOffers()
		server.ArbitrageOffersDB = append(server.ArbitrageOffersDB, &server.ArbitrageOpportunities{Pair: cycle, Offers: make([]*server.OfferSummary, 0)})

		if allOffers != nil {
			for i, offers := range allOffers {
				for _, offer := range offers {
					server.ArbitrageOffersDB[len(server.ArbitrageOffersDB)-1].Offers = append(server.ArbitrageOffersDB[len(server.ArbitrageOffersDB)-1].Offers, &server.OfferSummary{From: cycle[i], To: cycle[(i+1)%len(cycle)], Rate: offer.Rate, Hash: offer.TxHash, Volume: offer.Quantity})
				}
			}
		}
	}

	/*
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
					}
				}
				log.Println("====================================================================================")
				//return
			}
		}*/

}
