package main

import (
	"fmt"
	"time"
)

func main() {
	tkr := time.NewTicker(time.Second * 2) //ticker creation
	fmt.Println("this Ticker works in every 2 seconds")
	for i := 0; i < 20; i++ {
		select {
		case <-tkr.C:
			fmt.Printf("\nticker is working-%d -->>", i)
		default:
			fmt.Println("\nnot ticker")
			time.Sleep(time.Second * 1 / 2)

		}
	}
	tkr.Stop()
}
