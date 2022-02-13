package graph

import (
	"fmt"
	"testing"
)

func TestALDFS01(t *testing.T) {
	gptr1 := MakeALGraph(6, true)

	g1 := *gptr1

	//graph on p605 CLS
	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(0, 3))
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 1)) //e5
	g1.InsertEdge(CreateEdge(4, 2))
	g1.InsertEdge(CreateEdge(4, 5))
	g1.InsertEdge(CreateEdge(5, 5)) //e8

	g1.PrintGraphAL()
	//note in this case two trees in the forest

	ALDFS(gptr1)

	PrintALGraphNodeProperties(gptr1)
	s := 0
	v := 4
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	s = 3
	v = 4
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	s = 0
	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}

func TestALDFS02(t *testing.T) {

	//its a triangle
	gptr1 := MakeALGraph(3, false)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(0, 2))
	g1.InsertEdge(CreateEdge(1, 2))

	g1.PrintGraphAL()

	ALDFS(gptr1)

	PrintALGraphNodeProperties(gptr1)
	s := 0
	v := 1
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	s = 0
	v = 2
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	s = 0
	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}

func TestALDFS03(t *testing.T) {

	//its a square with one diagonal
	//top left to bottom right
	gptr1 := MakeALGraph(4, false)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(0, 2))
	g1.InsertEdge(CreateEdge(0, 3))
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))

	g1.PrintGraphAL()

	ALDFS(gptr1)

	fmt.Println("")
	PrintALGraphNodeProperties(gptr1)
	s := 0
	v := 3
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	s = 0
	v = 2
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	s = 0
	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}
