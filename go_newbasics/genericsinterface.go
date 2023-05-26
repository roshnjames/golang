package main

import (
	"fmt"
	"log"
)

// creating an interface to include all possible datatypes
type generics interface {
	int | int8 | int32 | int64 | float32 | float64
}

// generic function with the interface
func sum[input generics](arr []input) (ans input) {
	ans = 0
	for _, val := range arr {
		ans += val
	}
	return
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	ans1 := sum(slice)
	slice1 := []int8{0, 1, 2, 3, 4, 5}
	ans2 := sum(slice1)
	slice2 := []int64{0, 1, 2, 3, 4, 5}
	ans3 := sum(slice2)
	slice4 := []float64{0.00, 1.987, 2.2341, 3.2, 4.01, 5.9991}
	ans4 := sum(slice4)

	//printing
	fmt.Println("int :", ans1)
	fmt.Println("int8 :", ans2)
	fmt.Println("int64 :", ans3)
	_, err := fmt.Println("float64 :", ans4)
	if err != nil {
		log.Fatal("ERROR OCCURED")
	}

}
