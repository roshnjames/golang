package main

import (
	"fmt"
)

// generic function
func sum[input int | int64 | float64 | float32](arr []input) input {
	var net input
	for _, i := range arr {
		net += i
	}
	return net
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr1 := []int64{1, 2, 3, 4, 5}
	arr2 := []float64{1.02, 2.98, 3.43, 4.33, 5.90}
	ans := sum(arr)
	ans1 := sum(arr1)
	ans2 := sum(arr2)

	fmt.Println("sum is :", ans)
	fmt.Println("sum1 is :", ans1)
	fmt.Println("sum2 is :", ans2)
}

//generics means a generalised common function to deal with same functionality for different input datatypes
//sum is the common general function and it can deal with 'int | int64 | float64 | float32' datatypes as per required
