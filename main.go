package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"os/exec"
	display "github.com/gaspardpeduzzi/spring_block/display"
	server "github.com/gaspardpeduzzi/spring_block/server"
	graph "github.com/gaspardpeduzzi/spring_block/graph"
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

		server.LatestTx := liquidOptimizer.Graph.
		//Latest Opportunity


		if allOffers != nil {

			Print("====================================================================================")
			Print(time.Now().String())
			Print(fmt.Sprintf("Found profitable cycle: %s", cycle))

			hello := make([]*server.OfferSummary, 0)

			profit := 1.0

			for i, offer := range allOffers {
				Print(fmt.Sprintf("%s -> %s, %e OfferCreate Hash: %s, Volume: %f", cycle[i], cycle[(i+1)%len(cycle)], offer.Rate, offer.TxHash, offer.Quantity))
				profit = profit * offer.Rate

				summary := &server.OfferSummary{
					From:   offer.CreatorWillPay,
					To:     offer.CreatorWillGet,
					Rate:   offer.Rate,
					Hash:   offer.TxHash,
					Volume: offer.Quantity,
				}
				hello = append(hello, summary)
			}

			Print(fmt.Sprintf("Profit: %f", profit))
			Submit_Transaction(allOffers)
			save(allOffers)

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

		elapsed := time.Since(startingTime)
		log.Printf("Took %s to parse and analyse opportunities ", elapsed)

	}

}

func Submit_Transaction(cycle []graph.Offer) {

	maxQuantity_tmp := cycle[0].Quantity
	for i, offer := range cycle[:1] {
		if maxQuantity_tmp / offer.Rate > cycle[i+1].Quantity {
			maxQuantity_tmp = cycle[i+1].Quantity
		} else {
			maxQuantity_tmp = maxQuantity_tmp / offer.Rate
		}
	}
	goal := maxQuantity_tmp / cycle[len(cycle)-1].Rate

	if goal < 1000000 {
		Print(fmt.Sprintf("Max Quantity is too small: %v", goal))
		return
	}

	args := fmt.Sprintf("%s %v %v", cycle[0].CreatorWillPay, goal, maxQuantity_tmp)
	for i, offer := range cycle[1:] {
		args = fmt.Sprintf("%s %s %s", args, offer.CreatorWillPay, cycle[i].Issuer)
	}

	args = fmt.Sprintf("%s %s", args, "> output")

	Print(args)

	out, err := exec.Command("./submit.sh", args).Output()
	log.Println("submit out", string(out), err)
}

func Print(message string) {
	log.Println(message)
	out, err := exec.Command("./append2.sh", message).Output()
	log.Println("append2 out", string(out), err)
}

func save(cycle []graph.Offer) {
	res := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	for _, offer := range cycle {
		res = fmt.Sprintf("%s\n%s", res, offer.ToString())
	}
	out, err := exec.Command("./append.sh", res).Output()
	log.Println("append out", string(out), err)
}
