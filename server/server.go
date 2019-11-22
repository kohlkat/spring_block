package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

)


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


type OK struct {
	Welcome string
}
//[{step: 1, sent: "139.359 XRP", received: "1501.834 ULT",
// rate: 10.433, hash: "KJDSNFKJSDNF"},  {step: 2 ....} {step: 3 ...} ]
type Opportunity struct {
	Step int
	Sent string
	Received string
	Rate float64
	Hash string

}

type OpportunityInfo struct {
	Pairs string
	CycleSize int
	Volume float64
	Profit float64
}



// ArbitrageOffersDB : Datastructure to hold arbitrage opportunities
var ArbitrageOffersDB []*ArbitrageOpportunities
//The number of accounts active
var AccountsNumber int
//List the issuers
var Issuers []string
//List the clients
var Clients []string
//For a given account give the number of tx it is involved
var AccountOrders map[string]int

var LatestOpportunity []*Opportunity
var RecentOpportunities []*OpportunityInfo


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

func accountOrders(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/accountOrders" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		account := r.URL.Query().Get("account")
		err := json.NewEncoder(w).Encode(AccountOrders[account])
		//Install

		if err != nil {
			log.Println("Error encoding", err)
		}
	}
}

func latestOpportunity(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/latestOpportunity" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(LatestOpportunity)
		//Install

		if err != nil {
			log.Println("Error encoding", err)
		}
	}
}

func recentOpportunities(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/recentOpportunities" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(RecentOpportunities)
		//Install

		if err != nil {
			log.Println("Error encoding", err)
		}
	}
}



func LaunchServer() {
	log.Println("GUI Server up and running")
	ArbitrageOffersDB = make([]*ArbitrageOpportunities, 0)
	AccountOrders = make(map[string]int)
	LatestOpportunity = make([]*Opportunity, 0)
	RecentOpportunities = make([]*OpportunityInfo, 0)

	//fs := http.FileServer(http.Dir(""))
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)

	http.HandleFunc("/connect", connect)
	http.HandleFunc("/arbitrage", arbitrage)
	http.HandleFunc("/accountsNumber", accountsNumber)
	http.HandleFunc("/issuers", issuers)
	http.HandleFunc("/clients", clients)
	http.HandleFunc("/accountOrders", accountOrders)
	http.HandleFunc("/latestOpportunity", latestOpportunity)
	http.HandleFunc("/recentOpportunities", recentOpportunities)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}


}
