package simpleprogs

import (
	"fmt"
	"time"
)

func Spinner() {
	for {
		for _, r := range `aAbBcC` {
			fmt.Printf("\r%c", r)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
