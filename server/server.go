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
var Issuers []string
var Clients []string

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
// Send accounts number info
func accountsNumber(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/accountsNumber" {
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

//Send list of issuers
func issuers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/issuers" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
		case "GET":
			err := json.NewEncoder(w).Encode(Issuers)
			if err != nil {
				log.Println("Error encoding", err)
			}
	}
}

func clients(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/clients" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(Clients)
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
	http.HandleFunc("/accountsNumber", accountsNumber)
	http.HandleFunc("/issuers", issuers)
	http.HandleFunc("/clients", clients)


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
