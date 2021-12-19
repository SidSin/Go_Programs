package maxsubarray

import "fmt"

func FindMaxSubArray(A []int, low, high int) (int, int, int) {

	if low == high {
		return low, high, A[low]
	} else {

		mid := ((low + high) >> 1)
		fmt.Println("mid = ", mid)

		leftlow, lefthigh, leftsum := FindMaxSubArray(A, low, mid)
		rightlow, righthigh, rightsum := FindMaxSubArray(A, mid+1, high)
		crosslow, crosshigh, crosssum := MaxCrossingSubarray(A, low, mid, high)

		if leftsum >= rightsum && leftsum >= crosssum {
			return leftlow, lefthigh, leftsum
		} else if rightsum >= leftsum && rightsum >= crosssum {
			return rightlow, righthigh, rightsum
		} else {
			return crosslow, crosshigh, crosssum
		}

	}
}
