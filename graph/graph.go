package graph

import (
	"github.com/gaspardpeduzzi/spring_block/data"
	"sort"
	"sync"
)

var capacityList = 1000000

// Graph : Data structure for graph of offers
type Graph struct {
	Graph map[string]map[string]*TxList
	Lock  sync.Mutex
}

// SimplerGraph : Data structure for graph of best offers
type SimplerGraph struct {
	Graph map[string]map[string]float64
	Lock  sync.Mutex
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

func (graph *Graph) initGraph(pay string, get string){
	graph.Lock.Lock()
	if graph.Graph[pay] == nil {
		graph.Graph[pay] = make(map[string]*TxList)
		txlist := make([]*Offer, capacityList)
		init := TxList{List:txlist}
		graph.Graph[pay][get] = &init
	} else if graph.Graph[pay][get] == nil {
		txlist := make([]*Offer, capacityList)
		init := TxList{List:txlist}
		graph.Graph[pay][get] = &init
	}
	graph.Lock.Unlock()
}



// CreateSimpleGraph : function for creating a SimpleGraph
func (graph Graph) CreateSimpleGraph() SimplerGraph {
	simpleGraph := make(map[string]map[string]float64)
	for k1, v1 := range graph.Graph {
		for k2, v2 := range v1 {
			simpleGraph[k1][k2] = v2.List[0].Rate
		}
	}
	return SimplerGraph{Graph: simpleGraph, Lock: sync.Mutex{}}
}

// SortGraphWithTxs : function for creating a new graph with offers sorted by rates
func (graph Graph) SortGraphWithTxs() Graph {
	sortedGraph := make(map[string]map[string]*TxList)
	for _, v1 := range graph.Graph {
		for _, v2 := range v1 {
			sort.Slice(v2.List, func(i, j int) bool {
				return v2.List[i].Rate < v2.List[j].Rate
			})
		}
	}
	return Graph{Graph: sortedGraph, Lock: sync.Mutex{}}
}
