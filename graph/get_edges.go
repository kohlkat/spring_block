package graph

import (
  "math"
)

func CheckProfitable(edges map[int][]Offer) bool {
  product := 1.0
  for _, v := range edges {
    product = product * v[len(v)-1].Rate
  }
  return product > 1
}

func (graph *Graph) Bellm(assets []string) (map[int][]Offer) {

  // Make a copy of the graph
  res := make(map[int][]Offer)
  quantities := make(map[int]float64)
  keys_count := len(graph.Graph)

  for i, _ := range assets {
    // Get best edge
    edge := graph.Graph[assets[i]][assets[(i+1)%keys_count]].List[0]
    // Update total quantities for that edge
    quantities[i] = edge.Volume
    // Remove used edge from graph
    graph.Graph[assets[i]][assets[(i+1)%keys_count]].List = graph.Graph[assets[i]][assets[(i+1)%keys_count]].List[1:]
    // Update selected edges
    res[i] = make([]Offer, 1000000)
    res[i][0] = *edge
  }

  if !CheckProfitable(res) {
    panic("Positive cycle doesn't exist.")
  }

  for true {
    minQuantity := math.MaxFloat64
    for _, v := range quantities {
      if (v < minQuantity) {
        minQuantity = v
      }
    }

    next_edges := make([]*Offer, 100)

    var bottleneck_edge *int

    for i, v := range quantities {
      if (v == minQuantity && len(graph.Graph[assets[i]][assets[(i+1)%keys_count]].List) > 0) {
        *bottleneck_edge = i
      }
    }

    if bottleneck_edge == nil {
      return res
    } else {
      // Getting next edges
      copy(next_edges, graph.Graph[assets[*bottleneck_edge]][assets[(*bottleneck_edge+1)%keys_count]].List)
      // Removing first one if existing
      graph.Graph[assets[*bottleneck_edge]][assets[(*bottleneck_edge+1)%keys_count]].List = graph.Graph[assets[*bottleneck_edge]][assets[(*bottleneck_edge+1)%keys_count]].List[1:]
    }

    next_edge := next_edges[0]

    product := 1.0
    for i, edges := range res {
      if *bottleneck_edge == i {
          product = product * next_edge.Rate
      } else {
          product = product * edges[len(edges)-1].Rate
      }
    }

    if product > 1 {
      quantities[*bottleneck_edge] += next_edge.Volume
      res[*bottleneck_edge] = append(res[*bottleneck_edge], *next_edge)
    } else {
      return res
    }

  }

  return nil
}
