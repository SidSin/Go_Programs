package sortingprgrms

func StraightInsertionSort(slptr *[]int) *[]int {

	slc := *slptr
	slclen := len(slc)

	//start from 2nd element of the slice
	for j := 1; j < slclen; j++ {

		i := j - 1
		k := slc[j]

		compareandmove(i, k, slptr)

	}
	return slptr
}

func compareandmove(i int, k int, slptr *[]int) {

	slc := *slptr

	// fmt.Println("")
	// fmt.Println(fourspaces, "Key = ", k)

	for i >= 0 {
		if k < slc[i] {
			slc[i+1] = slc[i]
			// fmt.Println(fourspaces, "--- Array Changed ---")
			// fmt.Println(fourspaces, slc)
			i--
		} else {
			break
		}
	}

	// fmt.Println(fourspaces, "--- Old Array ---")
	// fmt.Println(fourspaces, slc)
	slc[i+1] = k
	// fmt.Println(fourspaces, "--- New Array ---")
	// fmt.Println(fourspaces, slc)

}
