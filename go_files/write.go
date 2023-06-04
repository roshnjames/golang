package main

import (
	"fmt"
	"log"
	"os"
)

func er(err error) {
	if err != nil {
		log.Fatal("::::::::", err)
	}
}
func main() {

	f, err := os.Create("test.txt")
	er(err)
	//1-------------
	b := []byte("This is Roshin james")
	_, err = f.Write(b)
	er(err)
	//2-------------------
	_, err = f.WriteString("\ngolang file management")
	er(err)
	//3----------------------
	_, err = f.WriteAt(b, 11)
	er(err)
	f.Close()

	//
	//f.Open("test.txt")
	//
	//
	//
	s := make([]byte, 50)
	f1, err := os.OpenFile("test.txt", os.O_RDWR, 0644)
	f1.Read(s)
	fmt.Println("file content is :::", string(s))

}
