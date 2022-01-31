package graph

import (
	"fmt"
	"testing"
)

//2 vertices one edge, undirected
func TestPrintGraph01(t *testing.T) {

	graph01 := MakeALGraph(2, false)

	graph01.InsertEdge(CreateEdge(0, 1))
	graph01.PrintGraphAL()
	fmt.Println("")
}

// A triangle with 3 edges, undirected
func TestPrintGraph02(t *testing.T) {

	graph01 := MakeALGraph(3, false)

	graph01.InsertEdge(CreateEdge(0, 1))
	graph01.InsertEdge(CreateEdge(1, 2))
	graph01.InsertEdge(CreateEdge(2, 0))

	graph01.PrintGraphAL()
	fmt.Println("")
}

// A triangle with 3 edges, directed
// 0 --->--- 1 --->--- 2
// |---------<---------|
func TestPrintGraph03(t *testing.T) {

	graph01 := MakeALGraph(3, true)

	graph01.InsertEdge(CreateEdge(0, 1))
	graph01.InsertEdge(CreateEdge(1, 2))
	graph01.InsertEdge(CreateEdge(2, 0))

	graph01.PrintGraphAL()
	fmt.Println("")

}

func TestPrintGraph04(t *testing.T) {

	graph01 := MakeALGraph(6, true)

	graph01.InsertEdge(CreateEdge(0, 1))
	graph01.InsertEdge(CreateEdge(0, 3))
	graph01.InsertEdge(CreateEdge(1, 4))
	graph01.InsertEdge(CreateEdge(2, 4))
	graph01.InsertEdge(CreateEdge(2, 5))
	graph01.InsertEdge(CreateEdge(3, 1))
	graph01.InsertEdge(CreateEdge(4, 3))
	graph01.InsertEdge(CreateEdge(5, 5))

	graph01.PrintGraphAL()
	fmt.Println("")

}
