package dsgraph

import (
	"fmt"
	"sidsin/srcmodule/graph"
	"testing"
)

func TestMSTKK01(t *testing.T) {

	gptr := graph.MakeALGraph(4, false)

	//its a square with two diagonals
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 1, 1))
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 2, 5)) //1st diagonal
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 3, 4))
	gptr.InsertEdge(graph.CreateWeightedEdge(1, 2, 2))
	gptr.InsertEdge(graph.CreateWeightedEdge(2, 3, 3))
	gptr.InsertEdge(graph.CreateWeightedEdge(1, 3, 6)) //2nd diagonal

	mst := MinSpanTreeKK(gptr)
	es := mst.GetEdgeSet()

	fmt.Println("Edges in the MST")
	for k := range *es {
		graph.PrintEdge(k, false)
	}
}

func TestMSTKK02(t *testing.T) {

	gptr := graph.MakeALGraph(4, false)

	//its a square with two diagonals
	//with the 2nd diagonal wt reduced to 1
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 1, 1))
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 2, 5)) //1st diagonal
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 3, 4))
	gptr.InsertEdge(graph.CreateWeightedEdge(1, 2, 2))
	gptr.InsertEdge(graph.CreateWeightedEdge(2, 3, 3))
	gptr.InsertEdge(graph.CreateWeightedEdge(1, 3, 1)) //2nd diagonal

	mst := MinSpanTreeKK(gptr)
	es := mst.GetEdgeSet()

	fmt.Println("")
	fmt.Println("Edges in the MST")
	for k := range *es {
		graph.PrintEdge(k, false)
	}
}

func TestMSTKK03(t *testing.T) {

	gptr := graph.MakeALGraph(4, false)

	//its a square with two diagonals
	//with the 1st diagonal wt reduced to 1
	//and with the edge (0,3) wt reduced to 1
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 1, 1))
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 2, 1)) //1st diagonal
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 3, 1))
	gptr.InsertEdge(graph.CreateWeightedEdge(1, 2, 2))
	gptr.InsertEdge(graph.CreateWeightedEdge(2, 3, 3))
	gptr.InsertEdge(graph.CreateWeightedEdge(1, 3, 6)) //2nd diagonal

	mst := MinSpanTreeKK(gptr)
	es := mst.GetEdgeSet()

	fmt.Println("Edges in the MST")
	for k := range *es {
		graph.PrintEdge(k, false)
	}
}
