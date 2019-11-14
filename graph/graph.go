package graph

import (
	"sort"
	"sync"
	"github.com/gaspardpeduzzi/spring_block/data"
)

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
