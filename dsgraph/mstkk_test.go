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

	mincost := 0

	fmt.Println("Edges in the MST")
	for k := range *es {
		graph.PrintEdge(k, false)
		mincost = mincost + k.GetEdgeWeight()
	}
	fmt.Println("Total Edge weight of the MST = ", mincost)

}

func TestMSTKK04(t *testing.T) {

	gptr := graph.MakeALGraph(9, false)

	// graph as on P632 CLS
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 1, 4))
	gptr.InsertEdge(graph.CreateWeightedEdge(0, 7, 8))
	gptr.InsertEdge(graph.CreateWeightedEdge(1, 7, 11))

	// edge b-c in graph
	//e4
	gptr.InsertEdge(graph.CreateWeightedEdge(1, 2, 8))

	// edge c-d in graph
	//e5
	gptr.InsertEdge(graph.CreateWeightedEdge(2, 3, 7))

	// edge d-e in graph
	//e6
	gptr.InsertEdge(graph.CreateWeightedEdge(3, 4, 9))

	// edge e-f in graph
	//e7
	gptr.InsertEdge(graph.CreateWeightedEdge(4, 5, 10))

	// edge f-g in graph
	//e8
	gptr.InsertEdge(graph.CreateWeightedEdge(5, 6, 2))

	// edge g-h in graph
	//e9
	gptr.InsertEdge(graph.CreateWeightedEdge(6, 7, 1))

	// edge c-i in graph
	//e10
	gptr.InsertEdge(graph.CreateWeightedEdge(2, 8, 2))

	// edge i-h in graph
	//e11
	gptr.InsertEdge(graph.CreateWeightedEdge(8, 7, 7))

	// edge i-g in graph
	//e12
	gptr.InsertEdge(graph.CreateWeightedEdge(8, 6, 6))

	// edge c-f in graph
	//e13
	gptr.InsertEdge(graph.CreateWeightedEdge(2, 5, 4))

	// edge d-f in graph
	//e14
	gptr.InsertEdge(graph.CreateWeightedEdge(3, 5, 14))

	mst := MinSpanTreeKK(gptr)
	es := mst.GetEdgeSet()

	mincost := 0
	mstedgecount := 0

	fmt.Println("Edges in the MST")
	for k := range *es {
		graph.PrintEdge(k, false)
		mincost = mincost + k.GetEdgeWeight()
		mstedgecount = mstedgecount + 1
	}
	fmt.Println("")
	fmt.Println("Total Edge weight of the MST = ", mincost)
	fmt.Println("Number of edges in the MST   = ", mstedgecount)
	fmt.Println("")
}
