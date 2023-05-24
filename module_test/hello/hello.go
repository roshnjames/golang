package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {

	message := greetings.Hello("Roshin James")
	fmt.Println(message)
}
