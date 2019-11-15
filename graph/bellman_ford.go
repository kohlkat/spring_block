package graph

import (
	"fmt"
	"math"
)

// BellmanFord : Run Bellman-Ford Algorithm on the Simpler Graph with the best rates
func (graph SimplerGraph) BellmanFord() (string, map[string]string) {

	v := len(graph.Currencies)

	// Initiate distances and predecessors array
	dist := make([]float64, v)
	predecessors := make(map[string]string)

	for i := 0; i < v; i++ {
		dist[i] = math.MaxFloat64
		predecessors[graph.Currencies[i]] = ""
	}

	// initialize distance of source as 0
	dist[0] = 0

	// Relax all edges |V| - 1 times
	for i := 0; i < v-1; i++ {
		for j := 0; j < v; j++ {
			for w := 0; w < v; w++ {
				if dist[w] > dist[j]+graph.Graph[graph.Currencies[j]][graph.Currencies[w]] {
					dist[w] = dist[j] + graph.Graph[graph.Currencies[j]][graph.Currencies[w]]
					predecessors[graph.Currencies[w]] = graph.Currencies[j] // indice to check
				}
			}
		}
	}

	// Find negative cycle
	for i := 0; i < v-1; i++ {
		for j := 0; j < v; j++ {
			if dist[j] > dist[i]+graph.Graph[graph.Currencies[i]][graph.Currencies[j]] {
				fmt.Println("Arbitrage opportunity")
				return graph.Currencies[j], predecessors
				// INSERT CYCLE OUTPUT
			}
		}

	}

	return "", nil

}
