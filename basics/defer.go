package main

import "fmt"

func main() {
	var n int
	fmt.Println("ENter the no of iterations")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		defer fmt.Println(i)
	}
	fmt.Print("after loop\nDEfER by default follows LIFO order")
}

//defer keyword is used to make a statement get executed only after the whole function is executed.
//In case of multiple defers it follows LIFO order
