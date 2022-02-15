package dsgraph

import (
	"sidsin/srcmodule/graph"
)

//Uses Kruskal's algorithm to find the minimum weight spanning tree (MST)of the given graph.
//The graph returned is the MST in the adjacency list representation.
//The arguemnt ALG is not modified by this function.
func MinSpanTreeKK(ALG *graph.ADJListGraph) *graph.ADJListGraph {

	//disjoint set representation of the graph pointed to by the argument ALG
	dsg := MakeDSGraph(ALG.ALVertexCount(), ALG.IsDirected())

	sortedgraph := graph.SortEdgesbyWeight(ALG)

	//MST with no edges
	mstgraph := graph.MakeALGraph(ALG.ALVertexCount(), ALG.IsDirected())

	for _, e := range *sortedgraph {

		x := e.GetStartVertexValue()
		y := e.GetEndVertexValue()

		m1 := FindSet(x, dsg)
		m2 := FindSet(y, dsg)

		if m1 != m2 {
			Union(m1, m2)
			//add edge to the MST
			mstgraph.InsertEdge(e)
		}
	}

	return mstgraph
}
