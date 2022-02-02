package graph

import (
	"fmt"
	"math"
	"sidsin/srcmodule/list"
)

//performs BFS on the given AL Graph
//and stores the results in each GraphNode/vertex
func ALBFS(ALGptr *ADJListGraph, s int) {

	//initialize graph nodes
	for i := range ALGptr.vertexset {
		if ALGptr.vertexset[i].value != s {
			ALGptr.vertexset[i].colour = "WHITE"
			ALGptr.vertexset[i].distance = math.MaxInt
			ALGptr.vertexset[i].parent = nil
		} else {
			//initialize the graphnode with value = s
			ALGptr.vertexset[i].colour = "GRAY"
			ALGptr.vertexset[i].distance = 0
			ALGptr.vertexset[i].parent = nil
		}
	}

	Q := list.CreateEmptyLinkedList()
	Q.AddatEnd(s)

	for !Q.IsLinkedListEmpty() {
		u := Q.GetListHead().GetNodeValue()
		Q.RemoveFromFront()

		L := ALGptr.adjlist[u]
		currnode := L.GetListHead()

		for currnode != nil {

			v := currnode.GetNodeValue()

			if ALGptr.vertexset[v].colour == "WHITE" {
				ALGptr.vertexset[v].colour = "GRAY"

				//restting distance to zero before incrementing it
				//this avoids setting distance to MinInt when incremented
				//also correctly calculates distances
				if ALGptr.vertexset[v].distance == math.MaxInt {
					ALGptr.vertexset[v].distance = 0
				}
				ALGptr.vertexset[v].distance = ALGptr.vertexset[u].distance + 1
				ALGptr.vertexset[v].parent = &ALGptr.vertexset[u]
				Q.AddatEnd(v)
			}
			currnode = currnode.GetNextNode()
		}
		ALGptr.vertexset[u].colour = "BLACK"
	}

}

func PrintPath(gptr *ADJListGraph, s, v int) {

	adjlistgraph := *gptr

	if v >= adjlistgraph.ALVertexCount() {
		fmt.Println("No such vertex as ", v)
		return
	}

	if v == s {
		fmt.Println(v)
	} else if adjlistgraph.vertexset[v].parent == nil {
		fmt.Println("No path from ", s, " to ", v, " exists.")
	} else {
		PrintPath(gptr, s, adjlistgraph.vertexset[v].parent.value)
		fmt.Println(v)
	}
}
