package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	//no recieve variables for channel ch
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*By default channels are unbuffered, meaning that they will only
accept sends (chan <-) if there is a corresponding receive (<- chan)
ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.*/
