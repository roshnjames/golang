package main

import (
	"fmt"
	"log"
	"os"
)

func er(err error) {
	if err != nil {
		log.Default().Fatalf("::::::::::", err)
	}
}

func main() {
	//f, err := os.Create("test.txt")
	f1, err := os.OpenFile("test.txt", os.O_CREATE, 0644)
	er(err)
	f1.WriteString("I am Roshinjames ")
	defer f1.Close()
	//
	f, err := os.Open("test.txt")
	er(err)
	s := make([]byte, 16)
	f.Read(s)
	fmt.Println(string(s))
}
