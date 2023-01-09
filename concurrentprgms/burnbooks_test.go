package simpleprogs

import (
	"fmt"
	"testing"
	"time"
)

func TestBookBurning(t *testing.T) {

	initcapacity := 6

	bs1 := NewBookshelf(initcapacity)

	c1 := make(chan int)

	fmt.Println("")
	fmt.Println("")

	bs1.PrintBookShelf()

	//fmt.Println("    bs1 capcity = ", bs1.capacity, bs1.Isempty())

	go BookPicker(bs1, 1, c1)
	go BookBurner(bs1, 2, c1)

	time.Sleep(time.Second * 20)

	fmt.Println("     Final Capacity = ", bs1.capacity)
	bs1.PrintBookShelf()

	fmt.Println("    Ending TestBookBurning")

}
