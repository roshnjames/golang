package main

import (
	"fmt"
	"log"
	"os"
)

func er(err error) {
	if err != nil {
		log.Fatal("error:::::::", err)
	}
}
func main() {
	f, err := os.ReadFile("test.txt")
	er(err)
	fmt.Println(string(f))
	fmt.Println("Length of file :", len(f))
	//
	f1, err := os.Open("test.txt")
	er(err)
	f1.Close()
	fmt.Println(f1)
	s := make([]byte, 13)
	f1.Read(s)
	fmt.Println("Slice of DATA is :::", string(s))
	//s1:=make([]byte,6)
	p, err := f1.Seek(6, 1)
	er(err)
	fmt.Println(p)
}
