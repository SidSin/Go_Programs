package graph

import (
	"fmt"
	"testing"
)

func TestALDFS(t *testing.T) {
	gptr1 := MakeALGraph(6, true)

	g1 := *gptr1

	//graph on p605 CLS
	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(0, 3))
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 1)) //e5
	g1.InsertEdge(CreateEdge(4, 2))
	g1.InsertEdge(CreateEdge(4, 5))
	g1.InsertEdge(CreateEdge(5, 5)) //e8

	g1.PrintGraphAL()
	//note in this case two trees in the forest

	ALDFS(gptr1)

	PrintALGraphNodeProperties(gptr1)
	s := 0
	v := 4
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	s = 3
	v = 4
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	s = 0
	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}
