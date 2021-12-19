package maxsubarray

import (
	"fmt"
	"testing"
)

func TestMCS01(t *testing.T) {

	A := []int{-4, 5, -2, 3, 6, 1, -1}
	mid := (0 + len(A)) / 2
	low := 0
	high := len(A) - 1

	lowlim, highlim, maxsum := MaxCrossingSubarray(A, low, mid, high)

	fmt.Println(lowlim, highlim, maxsum)

	if lowlim != 1 {
		t.Errorf("Low Limit is not 2.")
	} else if highlim != 5 {
		t.Errorf("High Limit is not 5.")
	} else if maxsum != 13 {
		t.Errorf("Max Sum is not 13.")
	}
}

func TestMCS02(t *testing.T) {

	A := []int{-1, 2, -3, 4, -5, 6, -7}
	mid := (0 + len(A)) / 2
	low := 0
	high := len(A) - 1

	lowlim, highlim, maxsum := MaxCrossingSubarray(A, low, mid, high)

	fmt.Println(lowlim, highlim, maxsum)

	if lowlim != 1 {
		t.Errorf("Low Limit is not 2.")
	} else if highlim != 5 {
		t.Errorf("High Limit is not 5.")
	} else if maxsum != 13 {
		t.Errorf("Max Sum is not 13.")
	}
}

func TestMCS03(t *testing.T) {

	A := []int{1, -2, 3, -4, 5, -6, 7}
	mid := (0 + len(A)) / 2
	low := 0
	high := len(A) - 1

	lowlim, highlim, maxsum := MaxCrossingSubarray(A, low, mid, high)

	fmt.Println(lowlim, highlim, maxsum)

	if lowlim != 1 {
		t.Errorf("Low Limit is not 2.")
	} else if highlim != 5 {
		t.Errorf("High Limit is not 5.")
	} else if maxsum != 13 {
		t.Errorf("Max Sum is not 13.")
	}
}
