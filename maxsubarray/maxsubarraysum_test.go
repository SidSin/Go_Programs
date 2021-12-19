package maxsubarray

import (
	"fmt"
	"testing"
)

func TestMaxSubArry01(t *testing.T) {

	A := []int{1, -2, 3}

	lowindex, highindex, maxsum := FindMaxSubArray(A, 0, len(A)-1)
	fmt.Println(lowindex, highindex, maxsum)
}

func TestMaxSubArry02(t *testing.T) {

	A := []int{-1, 2, -3}

	lowindex, highindex, maxsum := FindMaxSubArray(A, 0, len(A)-1)
	fmt.Println(lowindex, highindex, maxsum)
}

func TestMaxSubArry03(t *testing.T) {

	A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}

	fmt.Println("Array Len = ", len(A))

	lowindex, highindex, maxsum := FindMaxSubArray(A, 0, len(A)-1)
	fmt.Println(lowindex, highindex, maxsum)

	if maxsum != 43 {
		t.Errorf("Maxsum is not correct.")
	}
}
