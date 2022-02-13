package dsgraph

import (
	"fmt"
	"testing"
)

func TestMakeSet01(t *testing.T) {

	v := 3
	//make two sets each with element v
	x1 := MakeSet(v)
	x2 := MakeSet(v)

	m1 := *x1
	b1 := m1[v]
	b2 := m1[4]

	fmt.Println("b1 = ", b1, "b2 = ", b2)

	m2 := *x2
	b1 = m2[v]
	b2 = m2[4]

	fmt.Println("b1 = ", b1, "b2 = ", b2)

}

func TestUnion01(t *testing.T) {

	x1 := MakeSet(1)
	x2 := MakeSet(2)

	Union(x1, x2)

	x2 = nil

	PrintSet(x1)
	PrintSet(x2)
}

//check if duplicate keys are allowed
func TestUnion02(t *testing.T) {

	x1 := MakeSet(1)
	x2 := MakeSet(1)

	Union(x1, x2)
	x2 = nil

	PrintSet(x1)
	PrintSet(x2)

}

func TestFindSet01(t *testing.T) {

	dsgptr := MakeDSGraph(3, false)
	dsg := *dsgptr
	x1 := dsg.vertexset[0]
	x2 := dsg.vertexset[1]

	Union(x1, x2)
	x2 = nil

	PrintSet(x1)
	PrintSet(x2)

	x3 := FindSet(1, dsgptr)
	PrintSet(x3)

	x3 = FindSet(2, dsgptr)
	PrintSet(x3)

	x3 = FindSet(3, dsgptr)
	PrintSet(x3)
}
