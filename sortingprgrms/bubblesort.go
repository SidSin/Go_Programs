package sortingprgrms

import "fmt"

const fourspaces = "    "

//bound changes during execution
func FastBubbleSort(sp *[]int) *[]int {

	slc := *sp
	bound := len(slc) - 1 //max allowed index for j
	t := 0
	j := 0

	swapcounter := 0
	compcounter := 0

	for bound > 0 {

		for j < bound {

			compcounter++
			if slc[j] > slc[j+1] {
				temp := slc[j]
				slc[j] = slc[j+1]
				slc[j+1] = temp
				t = j
				swapcounter++
			}
			j++
		}

		j = 0
		bound = t
		fmt.Println(fourspaces, "New bound = ", bound)
		t = 0

	}

	fmt.Println("No of swaps = ", swapcounter, ", No of comparisons = ", compcounter)

	return sp

}

//in this case bound is fixed at slice lenght minus one
func BubbleSortV2(sp *[]int) *[]int {

	slc := *sp
	slclen := len(slc)
	swapcounter := 0
	compcounter := 0

	for i := 0; i < slclen; i++ {
		for j := 0; j < slclen-1; j++ {

			compcounter++
			if slc[j] > slc[j+1] {
				temp := slc[j]
				slc[j] = slc[j+1]
				slc[j+1] = temp
				swapcounter++

			}
		}
	}
	fmt.Println("No of swaps = ", swapcounter, ", No of comparisons = ", compcounter)
	return sp

}
