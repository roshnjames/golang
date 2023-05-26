package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	log.Println("hello world")
	//log.Panic("fininsh")
	log.Fatal(errors.New("Roshin's error🈳🈹"))
	fmt.Println("😁")
}

//log package has several functions used to print out
//every function of log package includes outputting the timestamp with the given string
//log.Fatal()-> exits the program after printing out the given string
