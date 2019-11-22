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
var AccountsNumber int

func connect(w http.ResponseWriter, r *http.Request) {
	log.Println("RECEIVED REQUEST")
	if r.URL.Path != "/connect" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
		case "GET":
			ack := OK{Welcome: "Connect to Jack The Rippler"}
			err := json.NewEncoder(w).Encode(ack)
			if err != nil {
				fmt.Println("error encoding peers", err)
			}
	}
}

func arbitrage(w http.ResponseWriter, r *http.Request) {
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

func accounts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/accounts" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(AccountsNumber)
		if err != nil {
			log.Println("Error encoding", err)
		}
	}
}

func LaunchServer() {
	log.Println("GUI Server up and running")
	ArbitrageOffersDB = make([]*ArbitrageOpportunities, 0)

	//fs := http.FileServer(http.Dir(""))
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)

	http.HandleFunc("/connect", connect)
	http.HandleFunc("/arbitrage", arbitrage)
	http.HandleFunc("/accounts", accounts)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}


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
