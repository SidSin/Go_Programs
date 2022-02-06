package graph

type GraphNode struct {
	value      int
	colour     string
	distance   int
	parent     *GraphNode
	finishtime int //used in DFS
}

func MakeGraphNode(val int, col string, dist int, p *GraphNode, ftime int) *GraphNode {

	gnp := &GraphNode{value: val, colour: col, distance: dist, parent: p, finishtime: ftime}
	return gnp
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
func (GNP *GraphNode) SetGraphNodeValue(v int) {
	GN := *GNP
	GN.value = v
}
