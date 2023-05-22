package main

import "fmt"

func main() {
	var n, ele int
	var loc int = -1
	fmt.Println("Enter the no. of elements in the array")
	fmt.Scanln(&n)
	arr := [10]int{}
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the %d the element :", i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println("Enter the element you want to search for :")
	fmt.Scanln(&ele)

	for i := 0; i < len(arr); i++ {
		if arr[i] == ele {
			loc = i + 1
			break
		}
	}

	if loc != -1 {
		fmt.Printf("Element found at location %d", loc)
	} else {
		fmt.Println("Element not found in the array")
	}
}
