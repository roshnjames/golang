package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	tmr := time.NewTimer(time.Second * 7)      //timer
	tkr := time.NewTicker(time.Second * 1 / 8) //ticker

	for i := 0; i < 30; i++ {
		select {
		case <-tkr.C:
			fmt.Printf("\nTicker Working-->%d", i)

		case <-tmr.C:
			fmt.Printf("\n\nTimer working==>>%d", i)
			tkr.Stop()
			tmr.Stop()
			os.Exit(0)
		default:
			fmt.Printf("\ninvalid=>>%d", i)
			time.Sleep(time.Second * 3)
		}
	}
}
