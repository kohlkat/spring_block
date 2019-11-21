package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type OK struct {
	Welcome string
}

// ArbitrageOffersDB : Datastructure to hold arbitrage opportunities
var ArbitrageOffersDB []*ArbitrageOpportunities





func connect(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/connect" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {

	case "GET":
		ack := OK{Welcome: "Connect to liquidOptimizer"}
		err := json.NewEncoder(w).Encode(ack)
		if err != nil {
			fmt.Println("error encoding peers", err)
		}
	}
}

func getArbitrageOpportunities(w http.ResponseWriter, r *http.Request) {


	if r.URL.Path != "/arbitrage" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(ArbitrageOffersDB)
		if err != nil {
			log.Println("Error encoding", err)
		}
	}
}

func LaunchServer() {
	fmt.Println("GUI Server up and running")
	ArbitrageOffersDB = make([]*ArbitrageOpportunities, 0)
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/connect", connect)
	http.HandleFunc("/arbitrage", getArbitrageOpportunities)

}

type ArbitrageOpportunities struct {
	Pair   []string
	Offers []*OfferSummary
}

type OfferSummary struct {
	From   string
	To     string
	Rate   float64
	Hash   string
	Volume float64
}
