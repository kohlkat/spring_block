package graph

import (
	"github.com/gaspardpeduzzi/spring_block/data"
	// "log"
	"math"
	"sort"
	"sync"
)

var capacityList = 0

// SimplerGraph : Data structure for graph of best offers (nxn grid, -log(edge))
type SimplerGraph struct {
	Graph      map[string]map[string]float64
	Currencies []string
	Lock       sync.Mutex
}

// TxList : Data structure for list of offers
type TxList struct {
	List []*Offer
}
// Offer : Data structure for an offer
type Offer struct {
	//Indexing
	XrpTx    		  data.Transaction
	TxHash   		  string
	Account  		  string
	SequenceNumber 	  int

	//Actual transaction data
	Rate   			  float64
	Quantity 		  float64
	CreatorWillPay    string
	CreatorWillGet    string
	Issuer			  string
}

type OrderBook struct {
	List []*Offer
	Pay string
	WillGet string
	//numberOfTransactions int
}

// Graph : Data structure for graph of offers
type Graph struct {
	Graph map[string]map[string]*OrderBook
	AccountRoots map[string]map[int]*Offer
	Lock  sync.RWMutex
}

func (ng *Graph) insertNewOffer(offer *Offer){
	//TODO: check if correct here for the A to B "policy"
	ng.initGraph(offer.CreatorWillPay, offer.CreatorWillGet)
	ng.Lock.Lock()
	ng.Graph[offer.CreatorWillPay][offer.CreatorWillGet].List = append(ng.Graph[offer.CreatorWillPay][offer.CreatorWillGet].List, offer)
	ng.Lock.Unlock()
}

func (graph *Graph) initGraph(pay string, get string) {
	graph.Lock.Lock()

	if graph.Graph[pay] == nil {
		graph.Graph[pay] = make(map[string]*OrderBook)
		if graph.Graph[pay][get] == nil {
			txlist := make([]*Offer, capacityList)
			init := OrderBook{List: txlist, Pay: pay, WillGet: get}
			graph.Graph[pay][get] = &init
		}
	} else if graph.Graph[pay][get] == nil {
		txlist := make([]*Offer, capacityList)
		init := OrderBook{List: txlist, Pay: pay, WillGet: get}
		graph.Graph[pay][get] = &init

	}
	graph.Lock.Unlock()
}



// CreateSimpleGraph : function for creating a SimpleGraph
func (graph *Graph) CreateSimpleGraph() SimplerGraph {

	currencies := graph.getCurrenciesList()

	var simpleGraph = map[string]map[string]float64{}

	for _, i := range currencies {
		simpleGraph[i] = map[string]float64{}
		for _, j := range currencies {
			simpleGraph[i][j] = math.MaxFloat64
		}
	}

	for k1, v1 := range graph.Graph {
		for k2, v2 := range v1 {
			if len(v2.List) > 0 {
				simpleGraph[k1][k2] = -math.Log(v2.List[0].Rate)
			}
		}
	}
	return SimplerGraph{Graph: simpleGraph, Currencies: currencies, Lock: sync.Mutex{}}
}

func (graph *Graph) getCurrenciesList() []string {
	currencies := make([]string, len(graph.Graph))
	i := 0
	for k := range graph.Graph {
		currencies[i] = k
		i++
	}
	return currencies
}

// SortGraphWithTxs : function for creating a new graph with offers sorted by rates
func (graph *Graph) SortGraphWithTxs() {
	sortedGraph := make(map[string]map[string]*OrderBook)
	for k1, v1 := range graph.Graph {
		sortedGraph[k1] = map[string]*OrderBook{}
		for k2, v2 := range v1 {
			list := v2.List
			sort.Slice(list, func(i, j int) bool {
				return v2.List[i].Rate > v2.List[j].Rate
			})
			//sortedGraph[k1][k2] = &TxList{List: list}
			//display.DisplayVerbose("VERIF", list[0], list[1])
			copy(graph.Graph[k1][k2].List, list)
		}
	}
}
