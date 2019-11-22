package main

import (
	"flag"
	"fmt"
	display "github.com/gaspardpeduzzi/spring_block/display"
	server "github.com/gaspardpeduzzi/spring_block/server"
	"reflect"
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
		//update server data
		server.AccountsNumber = len(liquidOptimizer.Graph.Clients)

		allOffers, cycle := liquidOptimizer.Graph.GetProfitableOffers()

		seq_nb := 1

		server.AccountsNumber = len(liquidOptimizer.Graph.Clients)
		//Create array of issuers and clients
		keys := reflect.ValueOf(liquidOptimizer.Graph.Issuers).MapKeys()
		for i := 0; i < len(keys); i++ {
			server.Issuers = append(server.Issuers, keys[i].String())
		}
		keys = reflect.ValueOf(liquidOptimizer.Graph.Clients).MapKeys()
		for i := 0; i < len(keys); i++ {
			server.Clients = append(server.Clients, keys[i].String())
		}

		keys = reflect.ValueOf(liquidOptimizer.Graph.AccountLedger).MapKeys()
		for i := 0; i < len(keys); i++ {
			server.AccountOrders[keys[i].String()] = len(liquidOptimizer.Graph.AccountLedger[keys[i].String()])

		}

		//server.LatestTx := liquidOptimizer.Graph.
		//Latest Opportunity


		if allOffers != nil {
			fmt.Println("Found profitable cycle:", cycle)
			fmt.Println("====================================================================================")
			hello := make([]*server.OfferSummary, 0)
			for i, offers := range allOffers {
				for _, offer := range offers {
					fmt.Println(cycle[i], "->", cycle[(i+1)%len(cycle)], offer.Rate, "OfferCreate Hash:", offer.TxHash, "Volume:", offer.Quantity)
					//offer.Submit_Transaction(seq_nb)
					seq_nb = seq_nb + 1
					summary := &server.OfferSummary{
						From:   offer.CreatorWillPay,
						To:     offer.CreatorWillGet,
						Rate:   offer.Rate,
						Hash:   offer.TxHash,
						Volume: offer.Quantity,

					}
					hello = append(hello, summary)
				}
			}
			server.ArbitrageOffersDB = append(server.ArbitrageOffersDB, &server.ArbitrageOpportunities{Pair: cycle, Offers: hello})


			//Latest Opportunity
			latest := server.ArbitrageOffersDB[len(server.ArbitrageOffersDB)-1]
			var product float64 = 1.0
			for i, offer := range latest.Offers {
				sent := product
				product = product*offer.Rate
				opp := &server.Opportunity{
					Step:     i,
					Sent:      fmt.Sprintf("%f", sent) + " "+ cycle[i] ,
					Received: fmt.Sprintf("%f", product) + " " +cycle[(i+1)%len(cycle)],
					Rate:     offer.Rate,
					Hash:     offer.Hash,
				}
				server.LatestOpportunity = append(server.LatestOpportunity, opp)
			}

			//Recent Opportunities
			recents := server.ArbitrageOffersDB
			if len(server.ArbitrageOffersDB) > 10 {
				recents = server.ArbitrageOffersDB[len(server.ArbitrageOffersDB)-11 : len(server.ArbitrageOffersDB)-1]
			}

			for _, offers := range recents {
				cycleSize := 1
				pairs := ""
				volume := offers.Offers[0].Volume
				var product float64 = 1.0
				for _, offer := range offers.Offers {
					product = product*offer.Rate
					pairs += offer.From
					cycleSize += 1
					if volume > offer.Volume {
						volume = offer.Volume
					}
				}

				opp := &server.OpportunityInfo{
					Pairs:     pairs,
					CycleSize: cycleSize,
					Volume:    volume,
					Profit:    product,
				}
				server.RecentOpportunities = append(server.RecentOpportunities, opp)
			}


			fmt.Println("====================================================================================")
		}
	}

}
