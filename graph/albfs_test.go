package graph

import (
	"fmt"
	"testing"
)

func TestALBFS01(t *testing.T) {

	gptr1 := MakeALGraph(8, false)

	g1 := *gptr1

	//graph on p596 CLS
	g1.InsertEdge(CreateEdge(0, 1))
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 4))
	g1.InsertEdge(CreateEdge(3, 5))

	g1.InsertEdge(CreateEdge(4, 5))
	g1.InsertEdge(CreateEdge(5, 6))
	g1.InsertEdge(CreateEdge(6, 7))
	g1.InsertEdge(CreateEdge(4, 7))
	g1.InsertEdge(CreateEdge(4, 6))

	g1.PrintGraphAL()

	fmt.Println("Before")
	PrintALGraphNodeProperties(gptr1)

	ALBFS(gptr1, 3)

	fmt.Println("After")
	PrintALGraphNodeProperties(gptr1)

}

//case where BFS is started from node zero
func TestPrintPath00(t *testing.T) {

	//graph on p596 CLS
	gptr1 := MakeALGraph(8, false)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 4))
	g1.InsertEdge(CreateEdge(3, 5)) //e5

	g1.InsertEdge(CreateEdge(4, 5)) //e6
	g1.InsertEdge(CreateEdge(4, 6))
	g1.InsertEdge(CreateEdge(5, 6))
	g1.InsertEdge(CreateEdge(5, 7))
	g1.InsertEdge(CreateEdge(6, 7)) //e10

	g1.PrintGraphAL()

	s := 0

	ALBFS(gptr1, s) //note the starting point is zero

	PrintALGraphNodeProperties(gptr1)

	//note here s is constrainted to be zero
	//because the BFS was done with node zero as source
	v := 5
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 7
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}

//case where BFS is started from node one
func TestPrintPath01(t *testing.T) {

	//graph on p596 CLS
	gptr1 := MakeALGraph(8, false)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 4))
	g1.InsertEdge(CreateEdge(3, 5)) //e5

	g1.InsertEdge(CreateEdge(4, 5)) //e6
	g1.InsertEdge(CreateEdge(4, 6))
	g1.InsertEdge(CreateEdge(5, 6))
	g1.InsertEdge(CreateEdge(5, 7))
	g1.InsertEdge(CreateEdge(6, 7)) //e10

	g1.PrintGraphAL()

	s := 1

	ALBFS(gptr1, s)

	PrintALGraphNodeProperties(gptr1)

	//note here s is constrainted to be one
	//because the BFS was done with node one as source
	v := 5
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 7
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}

//case where BFS is started from node two
func TestPrintPath02(t *testing.T) {

	//graph on p596 CLS
	gptr1 := MakeALGraph(8, false)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 4))
	g1.InsertEdge(CreateEdge(3, 5)) //e5

	g1.InsertEdge(CreateEdge(4, 5)) //e6
	g1.InsertEdge(CreateEdge(4, 6))
	g1.InsertEdge(CreateEdge(5, 6))
	g1.InsertEdge(CreateEdge(5, 7))
	g1.InsertEdge(CreateEdge(6, 7)) //e10

	g1.PrintGraphAL()

	s := 2

	ALBFS(gptr1, s)

	PrintALGraphNodeProperties(gptr1)

	//note here s is constrainted to be two
	//because the BFS was done with node two as source
	v := 5
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 7
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}

//case where BFS is started from node three
func TestPrintPath03(t *testing.T) {

	//graph on p596 CLS
	gptr1 := MakeALGraph(8, false)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 4))
	g1.InsertEdge(CreateEdge(3, 5)) //e5

	g1.InsertEdge(CreateEdge(4, 5)) //e6
	g1.InsertEdge(CreateEdge(4, 6))
	g1.InsertEdge(CreateEdge(5, 6))
	g1.InsertEdge(CreateEdge(5, 7))
	g1.InsertEdge(CreateEdge(6, 7)) //e10

	g1.PrintGraphAL()

	s := 3

	ALBFS(gptr1, s)

	PrintALGraphNodeProperties(gptr1)

	//note here s is constrainted to be three
	//because the BFS was done with node three as source
	v := 6
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 7
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}

//case where BFS is started from node five
func TestPrintPath05(t *testing.T) {

	//graph on p596 CLS
	gptr1 := MakeALGraph(8, false)

	g1 := *gptr1

	g1.InsertEdge(CreateEdge(0, 1)) //e1
	g1.InsertEdge(CreateEdge(1, 2))
	g1.InsertEdge(CreateEdge(2, 3))
	g1.InsertEdge(CreateEdge(3, 4))
	g1.InsertEdge(CreateEdge(3, 5)) //e5

	g1.InsertEdge(CreateEdge(4, 5)) //e6
	g1.InsertEdge(CreateEdge(4, 6))
	g1.InsertEdge(CreateEdge(5, 6))
	g1.InsertEdge(CreateEdge(5, 7))
	g1.InsertEdge(CreateEdge(6, 7)) //e10

	g1.PrintGraphAL()

	s := 5

	ALBFS(gptr1, s)

	PrintALGraphNodeProperties(gptr1)

	//note here s is constrainted to be three
	//because the BFS was done with node three as source
	v := 6
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 0
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 7
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

	v = 8
	fmt.Println("------Printing Path from ", s, " to ", v, "-----------------")
	PrintPath(gptr1, s, v)

}
