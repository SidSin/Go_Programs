package graph

import (
	"fmt"
	"testing"
)

//case where DFS is started from root node of a tree
func TestPrintDFSTreePath01(t *testing.T) {

	//graph as a tree
	gptr1 := MakeALGraph(8, true)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(1, 4))

	g1.InsertEdge(CreateEdge(0, 2))
	g1.InsertEdge(CreateEdge(2, 5))
	g1.InsertEdge(CreateEdge(2, 6)) //e5

	g1.InsertEdge(CreateEdge(0, 3)) //e6
	g1.InsertEdge(CreateEdge(3, 7)) //e7

	g1.PrintGraphAL()

	s := 0

	ALDFS(gptr1)

	PrintALGraphNodeProperties(gptr1)

	//note here s is constrainted to be zero
	v := 4
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 5
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 7
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 8
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}
