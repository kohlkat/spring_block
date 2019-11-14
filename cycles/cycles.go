package txGraph

type Vertice struct {
	in    string
	out   string
	rate  float64
	index string
}

type Graph struct {
	v []Vertice
}
