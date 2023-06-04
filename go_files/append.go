package main

import (
	"os"
)

func er(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	er(err)
	_, err = f.WriteString("\nroshin james")
	er(err)
	f.Truncate(34)
	f.Close()

}
