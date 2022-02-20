package sortingprgrms

import (
	"fmt"
	"testing"
)

//array in decrasing order
//worst case scenario for BS
func TestBSV101(t *testing.T) {

	nums := []int{99, 44, 6, 2, 1}
	fmt.Println("Before Sort")
	fmt.Println(nums)

	fmt.Println("")
	fmt.Println("After Sort")
	sortedslcptr := FastBubbleSort(&nums)
	fmt.Println(*sortedslcptr)
}

//partially sorted array
func TestBSV102(t *testing.T) {

	nums := []int{1, 2, 3, 99, 44}
	fmt.Println("Before Sort")
	fmt.Println(nums)

	sortedslcptr := FastBubbleSort(&nums)
	fmt.Println("")
	fmt.Println("After Sort")
	fmt.Println(*sortedslcptr)
}

//already sorted array as input
func TestBSV103(t *testing.T) {

	nums := []int{1, 2, 3, 44, 99}
	fmt.Println("Before Sort")
	fmt.Println(nums)

	sortedslcptr := FastBubbleSort(&nums)
	fmt.Println("")
	fmt.Println("After Sort")
	fmt.Println(*sortedslcptr)
}
