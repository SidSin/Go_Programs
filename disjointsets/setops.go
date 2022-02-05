package disjointsets

import "fmt"

//implementing makeset with a map
func MakeSet(x int) *map[int]bool {
	m := make(map[int]bool)
	m[x] = true
	return &m
}

func PrintSet(p *map[int]bool) {

	if p != nil {
		m := *p
		fmt.Println(m)
	} else {
		fmt.Println("Given set is empty.")
	}
}

//assumes that both x and y are not nil
//contents of y are added to x
//preserves the condition that keys are not repated
//after union even if x and y have the same element
func Union(x, y *map[int]bool) *map[int]bool {

	m1 := *x
	m2 := *y

	//add all elements of y to x
	for k := range m2 {
		m1[k] = true
	}

	return x
}

// func FindSet(x int) *map[int]bool {

// }
