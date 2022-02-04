package graph

type GraphNode struct {
	value      int
	colour     string
	distance   int
	parent     *GraphNode
	finishtime int //used in DFS
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
//getter functions for graphnode
///////////////////////////////////////////////////////////////
func (GN GraphNode) GetGraphNodeValue() interface{} {
	return GN.value
}

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
//setter functions for graphnode
///////////////////////////////////////////////////////////////
