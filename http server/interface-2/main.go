package main

import "fmt"

func main() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		go func(n int) {
			ch <- fmt.Sprintf("Goroutine %d done", n)
		}(i)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
