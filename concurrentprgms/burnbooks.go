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

// func PrintTime() {
// 	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000000"))
// }

func (b *Bookshelf) PrintBookShelf() {
	fmt.Println("###########################################")
	for i := 0; i < b.initcap; i++ {

		fmt.Println(b.slot[i].isbn, b.slot[i].status)

	}
	fmt.Println("###########################################")
}

func (b *Bookshelf) unloadbook() int {

	fmt.Println(" --- ", time.Now().Format("2006-01-02 15:04:05.000000000"), " --- ")
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

//No channel for communication used
// Two workers burn books at the same time (down to nanosecond level)
func BookBurner(b *Bookshelf, wokernum int) {

	booksburned := 0

	fmt.Println("")
	fmt.Println("(", wokernum, ") *** Initial capacity = ", b.capacity, b.Isempty(), " *** ")
	fmt.Println("")

	for !b.Isempty() {

		booknum := b.unloadbook()
		if booknum < 0 {

			fmt.Println("(", wokernum, ")", "No more books to pick.")
			break

		} else {
			//sleep for 1 sec to simulate transport time to incinerator
			fmt.Println("")
			fmt.Println("")
			fmt.Println("    (", wokernum, ") ------------------------------------------------")
			fmt.Println("    (", wokernum, ")", time.Now().Format("2006-01-02 15:04:05.000000000"), "Picked Book Num ", booknum)
			fmt.Println("    (", wokernum, ")", time.Now().Format("2006-01-02 15:04:05.000000000"), "Going to Incinerator ... >>")
			time.Sleep(1000000000)

			//sleep for 1 sec to simulate book load on incinerator

			//sleep for 1 sec to simulate book burning
			booksburned = booksburned + 1
			fmt.Println("    (", wokernum, ")", time.Now().Format("2006-01-02 15:04:05.000000000"), "Burning Book No ", booknum)
			b.slot[booknum].status = "Burned"

			//sleep for 1 sec to simulate transport time from incinerator
			fmt.Println("    (", wokernum, ") Going back to BookShelf <<... ")
			time.Sleep(1000000000)
		}
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("(", wokernum, ") No of books burned = ", booksburned)
}
