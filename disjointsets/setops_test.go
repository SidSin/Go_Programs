package disjointsets

import (
	"fmt"
	"testing"
)

func TestMakeSet01(t *testing.T) {

	v := 3
	x1 := MakeSet(v)
	x2 := MakeSet(v)

	if x1 == x2 {
		fmt.Println("Same Map")
	} else {
		fmt.Println("Diff Maps")
	}

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

	PrintSet(x1)
	PrintSet(x2)

}
