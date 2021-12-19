package maxsubarray

import (
	"fmt"
	"math"
)

func MaxCrossingSubarray(A []int, low, mid, high int) (int, int, int) {

	leftsum := math.MinInt32
	maxleft := 0

	maxleft, leftsum = findmaxLmaxLsum(mid, low, A)

	rightsum := math.MinInt32
	maxright := 0

	maxright, rightsum = findmaxRmaxRsum(mid, high, A)

	return maxleft, maxright, (leftsum + rightsum)
}

func findmaxRmaxRsum(mid int, high int, A []int) (int, int) {

	rightsum := 0
	maxrightindex := math.MinInt32
	sum := 0

	for i := (mid + 1); i <= high; i++ {
		sum = sum + A[i]
		if sum > rightsum {
			rightsum = sum
			maxrightindex = i
		}
		fmt.Println("(2) sum = ", sum, ",  A[", i, "] = ", A[i], " , maxrightindex = ", maxrightindex, ", rightsum = ", rightsum)
	}
	return maxrightindex, rightsum
}

func findmaxLmaxLsum(mid int, low int, A []int) (int, int) {

	sum := 0
	leftsum := math.MinInt32
	maxleftindex := 0

	for i := mid; i >= low; i-- {
		sum = sum + A[i]

		if sum > leftsum {
			leftsum = sum
			maxleftindex = i
			fmt.Println("(1) sum = ", sum, ",  A[", i, "] = ", A[i], " , maxleftindex = ", maxleftindex, ", leftsum = ", leftsum)
		}
	}
	return maxleftindex, leftsum
}
