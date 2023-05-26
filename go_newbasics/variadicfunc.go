package main

import "fmt"

func sum(nums ...int) {
	tot := 0
	for _, val := range nums {
		tot += val
	}
	fmt.Println("\n", tot)
}
func main() {
	fmt.Println("Variadic functions can take varying num of arguments")
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4)
}

//variadic functions has the ability to take up multiple arguments
