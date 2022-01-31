package graph

import "testing"

func TestALDFS(t *testing.T) {
	gptr1 := MakeALGraph(6, true)

	g1 := *gptr1

	//graph on p605 CLS
	g1.InsertEdge(CreateEdge(0, 1))
	g1.InsertEdge(CreateEdge(0, 3))
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 1))
	g1.InsertEdge(CreateEdge(4, 2))
	g1.InsertEdge(CreateEdge(4, 5))
	g1.InsertEdge(CreateEdge(5, 5))

	g1.PrintGraphAL()
	//note in this case two trees in the forest

	ALDFS(gptr1)

	PrintALGraphNodeProperties(gptr1)

}
