package sortingprgrms

import (
	"fmt"
	"testing"
)

//array in decrasing order
//worst case scenario for BS
func TestBSV201(t *testing.T) {

	nums := []int{99, 44, 6, 2, 1}
	fmt.Println("Before Sort")
	fmt.Println(nums)

	fmt.Println("")
	fmt.Println("After Sort V2")
	sortedslcptr := BubbleSortV2(&nums)
	fmt.Println(*sortedslcptr)
}

//partially sorted array
func TestBSV202(t *testing.T) {

	nums := []int{1, 2, 3, 99, 44}
	fmt.Println("Before Sort")
	fmt.Println(nums)

	fmt.Println("")
	fmt.Println("After Sort V2")
	sortedslcptr := BubbleSortV2(&nums)
	fmt.Println(*sortedslcptr)
}

//already sorted array as input
func TestBSV203(t *testing.T) {

	nums := []int{1, 2, 3, 44, 99}
	fmt.Println("Before Sort")
	fmt.Println(nums)

	fmt.Println("")
	fmt.Println("After Sort V2")
	sortedslcptr := BubbleSortV2(&nums)
	fmt.Println(*sortedslcptr)
}
