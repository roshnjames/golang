package main

import (
	"fmt"
	"math"
)

const j float64 = 1 //global scope constant variable

func main() {
	var num1 int
	var ans float64
	fmt.Println("ENter the number to get square root")
	fmt.Scanln(&num1)
	num := float64(num1)
	res := root(num)

	if res > 5 {
		ans = math.Floor(res)
	} else {
		ans = math.Ceil(res)
	}
	fmt.Printf("The Result is %.2f\n", ans)
	//%.2f indicates only two values after the decimal is printed
}

func root(x float64) (y float64) {
	y = math.Sqrt(x * j)
	return
	//return variable needs only one specification and that is possible inside the return arguments field
}
