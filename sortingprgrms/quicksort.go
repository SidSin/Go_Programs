package sortingprgrms

import "fmt"

func QuickSort(slcptr *[]int, p, r int) *[]int {

	if p < r {
		q := partition(slcptr, p, r)
		fmt.Println(" q = ", q)
		QuickSort(slcptr, p, q-1)
		QuickSort(slcptr, q+1, r)
	} else {
		fmt.Println("***", "p = ", p, " r= ", r, "***")
	}

	return slcptr
}

func partition(slcptr *[]int, p, r int) int {

	slc := *slcptr
	//slclen := len(slc)
	i := p - 1
	x := slc[r]
	temp := 0

	fmt.Println("--------------------------------------------")
	fmt.Println(" p = ", p, " r = ", r, " i = ", i, " x = ", x)

	//Note: j starts from p
	for j := p; j < r; j++ {
		fmt.Println("")
		fmt.Println(fourspaces, "j = ", j)
		if slc[j] <= x {
			i++
			fmt.Println(fourspaces, "i = ", i)
			temp = slc[i]
			slc[i] = slc[j]
			slc[j] = temp
			//fmt.Println(fourspaces, slc)
		}
	}

	fmt.Println("Exchanging elements on indicies ", (i + 1), " and r = ", r)
	temp = slc[i+1]
	slc[i+1] = slc[r]
	slc[r] = temp

	fmt.Println(slc)

	return i + 1

}
