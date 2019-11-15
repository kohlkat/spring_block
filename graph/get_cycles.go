package graph
func GetCycle(asset string, predecessors map[string]string) []string {

  cycle := make([]string, len(predecessors))
  next_asset := asset

  // Going backward on best predecessors until duplicate is found
  for !stringInSlice(next_asset, cycle) {
    cycle = append(cycle, next_asset)
    next_asset = predecessors[next_asset]
  }

  // Removing all asset before the cycle
  tmp := make([]string, len(cycle))
  copy(tmp, cycle)
  for i, b := range tmp {
    if b != next_asset {
      cycle = remove(cycle, i)
    }
  }

  // Remove first occurrence of next_asset
  remove(cycle, 0)

  return cycle
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
