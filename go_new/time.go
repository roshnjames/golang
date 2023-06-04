package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now().Weekday())
	t1 := time.Now()
	fmt.Println(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), t1.Location())
	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), t1.Location()) //creating a custom datetime
	fmt.Print("\n-------------------------------\n")
	fmt.Println("today is ", t2.Day(), ",", t2.Weekday(), " of ", t2.Month(), ", ", t2.Year())

}
