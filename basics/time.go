package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday()
	fmt.Println(day)

	nextday := time.Now().Weekday() + 1
	fmt.Println(nextday)

	fmt.Println("UTC:", time.Now().UTC())
	fmt.Println("Local hour is ", time.Now().Local().Hour())

}
