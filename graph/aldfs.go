package graph

var time int = 0

//performs DFS on the given Adj List graph
func ALDFS(algptr *ADJListGraph) {

	//mark all vertices as WHITE and with no parent
	for i := 0; i < algptr.ALVertexCount(); i++ {
		algptr.vertexset[i].colour = "WHITE"
		algptr.vertexset[i].parent = nil
	}

	//call DFSVisit for all WHITE vertices in the graph
	for i := 0; i < algptr.ALVertexCount(); i++ {
		if algptr.vertexset[i].colour == "WHITE" {
			DFSVisit(algptr, i)
		}
	}
}

func DFSVisit(algptr *ADJListGraph, u int) {
	time = time + 1
	algptr.vertexset[u].distance = time
	algptr.vertexset[u].colour = "GRAY"

	nodeptr := algptr.adjlist[u].GetListHead()

	for nodeptr != nil {
		v := nodeptr.GetNodeValue()

		if algptr.vertexset[v].colour == "WHITE" {
			algptr.vertexset[v].parent = &algptr.vertexset[u]
			DFSVisit(algptr, v)
		}
		nodeptr = nodeptr.GetNextNode()
	}

	algptr.vertexset[u].colour = "BLACK"
	time = time + 1
	algptr.vertexset[u].finishtime = time

}
