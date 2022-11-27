package simpleprogs

import (
	"fmt"
	"testing"
	"time"
)

func TestBookBurning(t *testing.T) {

	bs1 := newBookshelf(7)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("    bs1 capcity = ", bs1.capacity, bs1.Isempty())

	go BookBurner(bs1, 1)
	go BookBurner(bs1, 2)

	time.Sleep(1000000000 * 20)
	fmt.Println("    Ending TestBookBurning")

}
