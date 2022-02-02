package graph

import (
	"math"
)

type GraphNode struct {
	value      int
	colour     string
	distance   int
	parent     *GraphNode
	finishtime int
}

type Edge struct {
	u *GraphNode
	v *GraphNode
}

func CreateEdge(u, v int) Edge {

	ptr1 := &GraphNode{value: u, colour: "WHITE", distance: math.MaxInt, parent: nil, finishtime: math.MaxInt}
	ptr2 := &GraphNode{value: v, colour: "WHITE", distance: math.MaxInt, parent: nil, finishtime: math.MaxInt}

	return Edge{ptr1, ptr2}
}

type Graph interface {
	VertexCount() int
	EdgeCount() int
	IsDirected() bool
	InsertEdge(Edge)
	RemoveEdge(Edge)
	EdgeExists(Edge) bool
	GetAdjNodes(GraphNode)
}

///////////////////////////////////////////////////////////////
//getter functions
///////////////////////////////////////////////////////////////
func (GN GraphNode) GetGraphNodeValue() interface{} {
	return GN.value
}

// func (GN GraphNode) GetNextGraphNode() *GraphNode {
// 	return GN.next
// }

func (GN GraphNode) GetGraphNodeColour() string {
	return GN.colour
}

func (GN GraphNode) GetGraphNodeDistance() int {
	return GN.distance
}

func (GN GraphNode) GetGraphNodeParent() *GraphNode {
	return GN.parent
}

///////////////////////////////////////////////////////////////
//setter functions
///////////////////////////////////////////////////////////////
