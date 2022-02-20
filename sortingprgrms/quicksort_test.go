package sortingprgrms

import (
	"fmt"
	"testing"
)

func TestQS01(t *testing.T) {

	slc := []int{2, 1, 3, 4}

	fmt.Println(" --- Original Slice ---")
	fmt.Println(slc)

	sortedslcptr := QuickSort(&slc, 0, 3)

	fmt.Println(*sortedslcptr)
}

func TestQS02(t *testing.T) {

	slc := []int{1, 2, 3, 4}

	fmt.Println(" --- Original Slice ---")
	fmt.Println(slc)

	sortedslcptr := QuickSort(&slc, 0, 3)

	fmt.Println(*sortedslcptr)
}

func TestQS03(t *testing.T) {

	slc := []int{4, 3, 2, 1}
	slc2 := []int{1, 2, 3, 4}

	p := 0
	r := len(slc) - 1

	fmt.Println(" --- Original Slice ---")
	fmt.Println(slc)

	sortedslcptr := QuickSort(&slc, p, r)
	fmt.Println(*sortedslcptr)

	b, s := SliceAreEqual(sortedslcptr, &slc2)
	if b {
		t.Errorf(s)
	}

}
