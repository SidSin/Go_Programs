package graph

import "math"

type Edge struct {
	u      *GraphNode
	v      *GraphNode
	weight int
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
