package graph

import (
	"math"
	"sort"
	"sync"
	"log"
	"github.com/gaspardpeduzzi/spring_block/data"

)

var capacityList = 0

// Graph : Data structure for graph of offers
type Graph struct {
	Graph map[string]map[string]*TxList
	Lock  sync.RWMutex
}

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
	XrpTx  data.Transaction
	Rate   float64
	Volume float64
	Hash   string
	Pay    string
	Get    string
}

func (graph *Graph) addNewOffer(pay string, get string, offer *Offer) {
	graph.initGraph(pay, get)
	graph.Lock.Lock()
	graph.Graph[pay][get].List = append(graph.Graph[pay][get].List, offer)
	graph.Lock.Unlock()

}

func (graph *Graph) initGraph(pay string, get string) {
	graph.Lock.Lock()
	if graph.Graph[pay] == nil {
		graph.Graph[pay] = make(map[string]*TxList)
		txlist := make([]*Offer, capacityList)
		init := TxList{List: txlist}
		graph.Graph[pay][get] = &init
	} else if graph.Graph[pay][get] == nil {
		txlist := make([]*Offer, capacityList)
		init := TxList{List: txlist}
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
			if len(v2.List) > 2 {
				log.Println("check", v2.List[0].Rate > v2.List[1].Rate)
			}
			simpleGraph[k1][k2] = -math.Log(v2.List[0].Rate)
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
func (graph *Graph) SortGraphWithTxs() Graph {
	sortedGraph := make(map[string]map[string]*TxList)

	for k1, v1 := range graph.Graph {
		sortedGraph[k1] = map[string]*TxList{}
		for k2, v2 := range v1 {
			list := v2.List
			sort.Slice(list, func(i, j int) bool {
				return v2.List[i].Rate > v2.List[j].Rate
			})
			sortedGraph[k1][k2] = &TxList{List: list}
		}
	}
	return Graph{Graph: sortedGraph, Lock: sync.RWMutex{}}
}
