package dsgraph

import (
	"sidsin/srcmodule/graph"
)

func MinSpanTreeKK(ALG *graph.ADJListGraph) *graph.ADJListGraph {

	dsg := MakeDSGraph(ALG.ALVertexCount(), false)

	sortedgraph := graph.SortEdgesbyWeight(ALG)

	mstgraph := graph.MakeALGraph(ALG.ALVertexCount(), false)

	for _, e := range *sortedgraph {

		x := e.GetStartVertexValue()
		y := e.GetEndVertexValue()

		m1 := FindSet(x, dsg)
		m2 := FindSet(y, dsg)

		if m1 != m2 {
			Union(m1, m2)
			mstgraph.InsertEdge(e)
		}
	}

	return mstgraph
}
