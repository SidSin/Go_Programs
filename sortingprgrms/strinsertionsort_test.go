package sortingprgrms

import (
	"fmt"
	"testing"
)

func TestInsSort01(t *testing.T) {

	slc := []int{99, 44, 6, 2, 1}
	fmt.Println(" --- Original Array ---")
	fmt.Println(slc)

	sortedslcptr := StraightInsertionSort(&slc)

	fmt.Println(*sortedslcptr)
}

func TestInsSort02(t *testing.T) {

	slc := []int{1, 2, 6, 44, 99}
	fmt.Println(" --- Original Array ---")
	fmt.Println(slc)

	sortedslcptr := StraightInsertionSort(&slc)

	fmt.Println(*sortedslcptr)
}

func TestInsSort03(t *testing.T) {

	slc := []int{6, 10, 2}
	fmt.Println(" --- Original Array ---")
	fmt.Println(slc)

	sortedslcptr := StraightInsertionSort(&slc)

	fmt.Println("")
	fmt.Println(*sortedslcptr)
}

func TestInsSort04(t *testing.T) {

	slc := []int{6, 10, 2, 23, 1, 37, 3, 48}
	fmt.Println(" --- Original Array ---")
	fmt.Println(slc)

	sortedslcptr := StraightInsertionSort(&slc)

	fmt.Println("")
	fmt.Println(*sortedslcptr)
}
