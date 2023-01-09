package simpleprogs

import (
	"fmt"
	"time"
)

// V2
// added Book struct

type Book struct {
	title  string
	isbn   int
	status string
}

func NewBook(title string, isbn int) *Book {
	b := Book{title: title, isbn: isbn, status: "OnBookshelf"}
	return &b
}

type Bookshelf struct {
	capacity int
	initcap  int
	slot     []*Book
}

func NewBookshelf(capacity int) *Bookshelf {
	bs := Bookshelf{capacity: capacity, initcap: capacity}

	bs.slot = make([]*Book, capacity)

	for i := 0; i < bs.capacity; i++ {
		bs.slot[i] = NewBook(string(rune(i)), i)
	}
	return &bs
}

func (b *Bookshelf) PrintBookShelf() {
	fmt.Println("###########################################")
	for i := 0; i < b.initcap; i++ {

		fmt.Println(b.slot[i].isbn, b.slot[i].status)

	}
	fmt.Println("###########################################")
}

func (b *Bookshelf) unloadbook() int {

	fmt.Println(" --- ", timenow(), " --- ")
	pickedbook := -1

	if b.capacity >= 1 {

		for i := 0; i < b.initcap; i++ {
			if b.slot[i].status == "OnBookshelf" {
				b.slot[i].status = "Picked"
				b.capacity = b.capacity - 1
				fmt.Println("New capacity = ", b.capacity)
				pickedbook = b.slot[i].isbn
				break
			}
		}
	} else {
		fmt.Println("Bookshelf is empty.")
	}

	fmt.Println("Book Unloaded = ", pickedbook)
	return pickedbook
}

func (b *Bookshelf) Isempty() bool {

	if b.capacity >= 1 {
		return false
	} else {
		return true
	}
}

//Added channel for co-ordination
//Divided workers as bookpicker and bookburner

func burningtime() {
	time.Sleep(time.Second)
}

func traveltime() {
	time.Sleep(time.Second)
}

func timenow() string {
	return time.Now().Format("2006-01-02 15:04:05.000000000")
}

func BookPicker(b *Bookshelf, wokernum int, outchannel chan int) {
	fmt.Println("")
	fmt.Println("(", wokernum, ") *** Initial capacity = ", b.capacity, b.Isempty(), timenow(), " *** ")
	fmt.Println("")

	for !b.Isempty() {

		booknum := b.unloadbook()

		if booknum < 0 {

			fmt.Println("(", wokernum, ")", "No more books to pick.")
			break

		} else {
			//sleep for 1 sec to simulate transport time to staging area
			fmt.Println("")
			fmt.Println("")
			fmt.Println("    (", wokernum, ") ------------------------------------------------")
			fmt.Println("    (", wokernum, ")", timenow(), "Picked Book Num ", booknum)
			fmt.Println("    (", wokernum, ")", timenow(), "Going to staging area ... >>")
			traveltime()

			outchannel <- booknum

			fmt.Println("    (", wokernum, ")", timenow(), "Book dropped to staging,Going back to bookshelf ... <<")
			traveltime()

		}

	}
	close(outchannel)
}

func BookBurner(b *Bookshelf, wokernum int, inchannel chan int) {

	booksburned := 0

	for booknum := range inchannel {

		//booknum := <-inchannel
		fmt.Println("    (", wokernum, ")", timenow(), "Value received from channel = ", booknum)

		//sleep for 1 sec to simulate going to incinerator
		traveltime()

		//sleep for 1 sec to simulate book burning
		fmt.Println("    (", wokernum, ")", timenow(), "Burning Book No ", booknum)
		burningtime()
		booksburned = booksburned + 1
		b.slot[booknum].status = "Burned"

		//sleep for 1 sec to simulate transport time from incinerator
		fmt.Println("    (", wokernum, ")", timenow(), "Going back to staging area <<... ")
		traveltime()
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("(", wokernum, ") No of books burned = ", booksburned)
}
