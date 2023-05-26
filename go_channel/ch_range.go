package main

import "fmt"

func main() {
	ch := make(chan int, 7)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)

	//ranging over channel
	for v := range ch {
		fmt.Println(v)
	}
}
