package main

import (
	"fmt"
	"time"
)

func diviByThree(ch chan int, n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			ch <- i
		}
		time.Sleep(100 * time.Millisecond)
	}
	defer close(ch)
}

func main() {
	ch := make(chan int)

	go diviByThree(ch, 100)

	for v := range ch {
		fmt.Println(v)
	}
}
