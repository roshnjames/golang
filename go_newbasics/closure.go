package main

import "fmt"

func main() {
	n, m := 9, 8

	//closure function
	square := func(num ...int) int {
		tot := 0
		for _, v := range num {
			tot += v
		}
		return tot
	}

	fmt.Println(square(1, 2, 3))
	fmt.Println(square(n))
	fmt.Println(square(n, m))
}

//closures in go is similar to inline functions ~ i.e functions without a name of its own
//above program implements variadic function inside a closure
