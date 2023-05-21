package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter the frist number")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number")
	fmt.Scanln(&num2)
	res := sum(num1, num2)
	fmt.Printf("The Sum is %d\n", res)
}

func sum(x, y int) (ans int) {
	ans = x + y
	return
}
