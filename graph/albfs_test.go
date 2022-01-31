package graph

import (
	"fmt"
	"testing"
)

func TestALBFS01(t *testing.T) {

	gptr1 := MakeALGraph(8, false)

	g1 := *gptr1

	//graph on p596 CLS
	g1.InsertEdge(CreateEdge(0, 1))
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 4))
	g1.InsertEdge(CreateEdge(3, 5))

	g1.InsertEdge(CreateEdge(4, 5))
	g1.InsertEdge(CreateEdge(5, 6))
	g1.InsertEdge(CreateEdge(6, 7))
	g1.InsertEdge(CreateEdge(4, 7))
	g1.InsertEdge(CreateEdge(4, 6))

	g1.PrintGraphAL()

	fmt.Println("Before")
	PrintALGraphNodeProperties(gptr1)

	ALBFS(gptr1, 3)

	fmt.Println("After")
	PrintALGraphNodeProperties(gptr1)

}

func TestPrintPath01(t *testing.T) {

	//graph on p596 CLS
	gptr1 := MakeALGraph(8, false)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1))
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 4))
	g1.InsertEdge(CreateEdge(3, 5))

	g1.InsertEdge(CreateEdge(4, 5))
	g1.InsertEdge(CreateEdge(5, 6))
	g1.InsertEdge(CreateEdge(6, 7))
	g1.InsertEdge(CreateEdge(4, 7))
	g1.InsertEdge(CreateEdge(4, 6))

	//ALBFS(gptr1, 0) //note the starting point is zero
	ALBFS(gptr1, 3)

	PrintALGraphNodeProperties(gptr1)
	fmt.Println("-----------------------")
	PrintPath(gptr1, 3, 5)
}
