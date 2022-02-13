package graph

import (
	"fmt"
	"math"
	"sort"
)

type Edge struct {
	u      *GraphNode
	v      *GraphNode
	weight int
}

func (e Edge) GetStartVertexValue() int {

	return e.u.value
}

func (e Edge) GetEndVertexValue() int {

	return e.v.value
}

func CreateEdge(u, v int) Edge {

	ptr1 := &GraphNode{value: u, colour: "WHITE", distance: math.MaxInt, parent: nil, finishtime: math.MaxInt}
	ptr2 := &GraphNode{value: v, colour: "WHITE", distance: math.MaxInt, parent: nil, finishtime: math.MaxInt}

	return Edge{ptr1, ptr2, math.MinInt}
}

func CreateWeightedEdge(u, v, w int) Edge {

	ptr1 := &GraphNode{value: u, colour: "WHITE", distance: math.MaxInt, parent: nil, finishtime: math.MaxInt}
	ptr2 := &GraphNode{value: v, colour: "WHITE", distance: math.MaxInt, parent: nil, finishtime: math.MaxInt}

	return Edge{ptr1, ptr2, w}
}

func (e *Edge) SetEdgeWeight(w int) {
	e.weight = w
}

func (e *Edge) GetEdgeWeight() int {
	return e.weight
}

type byEdgeWeight []Edge

//implementing Len,Less and Swap from sort.Interface
func (es byEdgeWeight) Len() int {

	return len(es)
}

//compare weights of two edges
func (es byEdgeWeight) Less(i, j int) bool {

	return es[i].weight < es[j].weight
}

func (es byEdgeWeight) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]

}

//Returns a pointer to a slice of edges
//where slice of edges = byEdgeweight
func SortEdgesbyWeight(algptr *ADJListGraph) *byEdgeWeight {

	//slice of edges
	es := make([]Edge, len(algptr.edgeset))
	i := 0

	//populate slice of edges in any order
	for e := range algptr.edgeset {
		es[i] = e
		i++
	}

	//convert []Edge to byEdgeWeight
	bew := byEdgeWeight(es)

	//sort the slice
	sort.Sort(bew)

	return &bew
}

//isdg equals true for directed graphs
func PrintEdge(e Edge, isdg bool) {

	if isdg {
		fmt.Println(e.u.value, "-->", e.v.value, " , weight = ", e.weight)
	} else {
		fmt.Println(e.u.value, "--", e.v.value, " , weight = ", e.weight)
	}
}
