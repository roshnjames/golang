package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "input to the channel is done like this"

	out := <-ch
	fmt.Println(out)

}
