package graph

import "testing"

func TestAdjGraphCreation01(t *testing.T) {

	gptr1 := MakeALGraph(3, false)

	g1 := *gptr1

	weight := 0

	//its a triangle
	g1.InsertEdge(CreateWeightedEdge(0, 1, weight))
	g1.InsertEdge(CreateWeightedEdge(1, 2, weight))
	g1.InsertEdge(CreateWeightedEdge(2, 0, weight))

	gptr1.PrintEdgeSet()

}
