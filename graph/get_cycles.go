package graph

import (
	"log"
)

func GetCycle(asset string, predecessors map[string]string) []string {

  // log.Println("predecessors", asset, predecessors)

  cycle := make([]string, 0)
  next_asset := asset

  // Going backward on best predecessors until duplicate is found
  for !stringInSlice(next_asset, cycle) {
    cycle = append(cycle, next_asset)
    next_asset = predecessors[next_asset]
  }

  // Removing all asset before the cycle
  tmp := make([]string, len(cycle))
  copy(tmp, cycle)
  for _, b := range tmp {
    if b != next_asset {
      cycle = remove(cycle, 0)
    } else {
      break
    }
  }

  // Reverse the list
  for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
    cycle[i], cycle[j] = cycle[j], cycle[i]
  }

  res := adjust_cycle(cycle)

  if res == nil {
    log.Println("Could not find primary asset in cycle", cycle)
    return nil
  }
  return res
}


// Order the cycle so that either USD, XRP, BTC or ETH is first.
func adjust_cycle(cycle []string) []string {
	i := 0
	primary_assets := []string{"XRP"}

  newCycle := make([]string, len(cycle))
  copy(newCycle, cycle)

	for !contains(primary_assets, newCycle[0]) {

    newCycle = make([]string, len(cycle))

    // Offset all element by one
    // cycle[len(cycle)-1] = cycle[0]
		for i, _ := range cycle {
			newCycle[i] = cycle[(i+1)%len(cycle)]
		}

		// At this point we have moved the cycle around entirely
		if i > len(cycle) {
			return nil
		}
		i = i + 1
	}

	return newCycle
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}


func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}
