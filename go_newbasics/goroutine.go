package main

import (
	"fmt"
	"time"
)

func hello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello ", name, " 😊")
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	go hello("Roshin")
	hello("poothakkuzhyil")
	go hello("james")

	time.Sleep(time.Second)
}

//go routines are used to implement concurrent execution in go
//implemented with the help of 'go'keyword
//functions starting with go keyword is part of the concurrent execution
//time package is used to put to sleep mode
