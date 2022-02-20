package sortingprgrms

import (
	"testing"
)

func TestSlCompare01(t *testing.T) {

	slc1 := []int{4, 3, 2, 1}
	slc2 := []int{4, 3, 2, 1}

	b, s := SliceAreEqual(&slc1, &slc2)
	if !b {
		t.Errorf(s)
	}
}

//diff lenght
func TestSlCompare02(t *testing.T) {

	slc1 := []int{4, 3, 2, 1}
	slc2 := []int{4, 3, 2}

	b, s := SliceAreEqual(&slc1, &slc2)
	if b {
		t.Errorf(s)
	}
}

//both slices are empty
func TestSlCompare03(t *testing.T) {

	slc1 := []int{}
	slc2 := []int{}

	b, s := SliceAreEqual(&slc1, &slc2)

	if !b {
		t.Errorf(s)
	}
}
