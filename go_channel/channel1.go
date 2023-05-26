package main

import (
	"fmt"
	"time"
)

func setVal(c chan string, str string) {
	c <- str //setting value to channel c
	fmt.Println(<-c)
	time.Sleep(time.Millisecond)

}

func main() {
	//creating a channel with 'chan' keyword in make function
	ch1 := make(chan string, 3)

	str1 := "Roshin James😊"
	str2 := "Hello Roshin😀"

	go setVal(ch1, str1)
	go setVal(ch1, str2)
	time.Sleep(time.Second * 2)
	//val := <-ch1 //getting value from channel ch1 and setting it to the new variable val
	//fmt.Println(val)
}
