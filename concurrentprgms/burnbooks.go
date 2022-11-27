package simpleprogs

import (
	"fmt"
	"time"
)

type Bookshelf struct {
	capacity int
}

func newBookshelf(capacity int) *Bookshelf {
	bs := Bookshelf{capacity: capacity}
	return &bs
}

func (b *Bookshelf) unloadbook() {

	if b.capacity >= 1 {
		b.capacity = b.capacity - 1
	} else {
		fmt.Println("Bookshelf is empty.")
	}
}

func (b *Bookshelf) Isempty() bool {

	if b.capacity >= 1 {
		return false
	} else {
		return true
	}
}

func BookBurner(b *Bookshelf, wokernum int) {

	booksburned := 0

	fmt.Println("    Initial capacity = ", b.capacity, b.Isempty())

	for !b.Isempty() {

		b.unloadbook()
		//sleep for 1 sec to simulate transport time to incinerator
		fmt.Println("")
		fmt.Println("")
		fmt.Println("    (", wokernum, ")  ------------------------------------------------")
		fmt.Println("    (", wokernum, ") Going to Incinerator ... >>")
		time.Sleep(1000000000)

		//sleep for 1 sec to simulate book load on incinerator

		//sleep for 1 sec to simulate book burning
		booksburned = booksburned + 1
		fmt.Println("    (", wokernum, ")  Burning Book No ", booksburned)

		//sleep for 1 sec to simulate transport time from incinerator
		//fmt.Println("    Going back to BookShelf <<... ")
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("(", wokernum, ")    No of books burned = ", booksburned)
}
