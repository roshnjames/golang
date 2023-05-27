package main

import (
	"fmt"
	"sync"
	"time"
)

func display(c chan int, wg *sync.WaitGroup) {
	for v := range c {
		fmt.Println("\n", <-v, "\ti=", v)
		wg.Done()
	}
}

func main() {
	ch := make(chan int, 10)
	wg := new(sync.WaitGroup)
	tkr := time.NewTicker(time.Second * 1)

	for i := 0; i < 10; {
		select {
		case <-tkr.C:
			ch <- i
			wg.Add(1)

		default:
			fmt.Print("\ntiker running...")
			time.Sleep(time.Second * 2)

		}
	}
	tkr.Stop()
	//close(ch)
	display(ch, wg)
	wg.Wait()
}

/*Rate limiting is an important mechanism for controlling resource utilization and
maintaining quality of service. Go elegantly supports rate limiting with goroutines,
 channels, and tickers.*/
