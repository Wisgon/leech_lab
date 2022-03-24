package graph_structure

type Graph struct {
	Nodes []*Node
}

type SuperGraph struct {
	Graphs []*Graph
}
