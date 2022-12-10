package simpleprogs

import (
	"fmt"
	"testing"
	"time"
)

func TestBookBurning(t *testing.T) {

	initcapacity := 6

	bs1 := NewBookshelf(initcapacity)

	fmt.Println("")
	fmt.Println("")

	bs1.PrintBookShelf()

	//fmt.Println("    bs1 capcity = ", bs1.capacity, bs1.Isempty())

	go BookBurner(bs1, 1)
	go BookBurner(bs1, 2)
	// go BookBurner(bs1, 3)
	// go BookBurner(bs1, 4)
	// go BookBurner(bs1, 5)
	// go BookBurner(bs1, 6)
	// go BookBurner(bs1, 7)
	// go BookBurner(bs1, 8)
	// go BookBurner(bs1, 9)
	// go BookBurner(bs1, 10)

	time.Sleep(1000000000 * 10)

	fmt.Println("     Final Capacity = ", bs1.capacity)
	bs1.PrintBookShelf()

	fmt.Println("    Ending TestBookBurning")

}
