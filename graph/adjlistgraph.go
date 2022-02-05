package graph

import (
	"fmt"
	"sidsin/srcmodule/list"
)

type ADJListGraph struct {
	vertexset  []GraphNode
	adjlist    []*list.LinkedList
	edgeset    map[Edge]bool
	isdirected bool
}

func (ALG ADJListGraph) ALVertexCount() int {
	return len(ALG.vertexset)
}

func (ALG ADJListGraph) ALEdgeCount() int {
	return len(ALG.edgeset)
}

func (ALG ADJListGraph) IsDirected() bool {
	return ALG.isdirected
}

//makes a graph with Adjacency list representation
//No edges are added
//Note: Vertices start from zero
func MakeALGraph(vertexcount int, isdirected bool) *ADJListGraph {

	vertexset := make([]GraphNode, vertexcount)

	//assign value to each vertex
	for i := 0; i < vertexcount; i++ {
		vertexset[i].value = i
	}

	adjlist := make([]*list.LinkedList, vertexcount)

	//initialize adj lists
	for i := 0; i < vertexcount; i++ {
		adjlist[i] = list.CreateEmptyLinkedList()
	}

	edgeset := make(map[Edge]bool)

	adlgraph := ADJListGraph{vertexset: vertexset, adjlist: adjlist, edgeset: edgeset, isdirected: isdirected}

	return &adlgraph
}

func (ALG ADJListGraph) InsertEdge(e Edge) {

	startnodevalue := e.u.value
	endnodevalue := e.v.value

	//add end vertex to the adjacency list of start vertex
	ALG.adjlist[startnodevalue].AddatEnd(endnodevalue)
	if !ALG.isdirected {
		//add start vertex to the adj list of end vertex
		ALG.adjlist[endnodevalue].AddatEnd(startnodevalue)
	}
}

//Prints the adjacency list representation of the graph
//for which it is called
//Note: vertices start from zero
func (ALG ADJListGraph) PrintGraphAL() {

	fmt.Println(" --- Adj List of given graph --- ")

	for i := range ALG.adjlist {
		fmt.Print(i, ": ")

		listheadptr := ALG.adjlist[i]
		currnode := listheadptr.GetListHead()

		for currnode != nil {
			fmt.Print(currnode.GetNodeValue(), " ")
			currnode = currnode.GetNextNode()
		}
		fmt.Println("")
	}
}

// after performing BFS or DFS on a graph, this func
// can be used to print the properties of each node
// in the graph
func PrintALGraphNodeProperties(algptr *ADJListGraph) {

	algraph := *algptr

	for i := 0; i < algraph.ALVertexCount(); i++ {

		parentvalue := -1
		graphnode := algraph.vertexset[i]

		if graphnode.parent != nil {
			parentvalue = graphnode.parent.value
		}

		fmt.Print("Vertex#:", graphnode.value, " -- ", "Colour:", graphnode.colour, " -- ", "D/Ft:", graphnode.distance, "/", graphnode.finishtime, " -- ", "Parent:", parentvalue)
		fmt.Println()
	}

}
