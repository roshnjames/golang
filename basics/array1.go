package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 3 {
			defer fmt.Println(arr[i]) //printing 3 at end using defer
		} else {
			fmt.Println(arr[i])
		}
	}
}
