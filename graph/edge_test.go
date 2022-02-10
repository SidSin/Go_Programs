package graph

import (
	"fmt"
	"testing"
)

func TestEdgeSort01(t *testing.T) {

	gptr1 := MakeALGraph(3, false)

	g1 := *gptr1

	//its a triangle
	g1.InsertEdge(CreateWeightedEdge(0, 1, 10))
	g1.InsertEdge(CreateWeightedEdge(1, 2, -12))
	g1.InsertEdge(CreateWeightedEdge(2, 0, 11))

	gptr1.PrintEdgeSet()

	sortedbywt := SortEdgesbyWeight(gptr1)
	fmt.Println("--- --- --- --- --- --- ")
	fmt.Println(sortedbywt)

	s := *sortedbywt

	fmt.Println("--- --- --- Using PrintEdge--- --- --- ")
	for _, k := range s {
		PrintEdge(k, g1.isdirected)
	}
}
