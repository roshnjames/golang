package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 9, 2, 8, 5, 6, 7, 3, 8, 4, 5, 0}
	arr2 := []string{"and", "mkjs", "hdu", "uhiue", "irini", "pkfpf"}
	fmt.Println("Before sorting----------------")

	fmt.Println(arr)
	fmt.Println(arr2)
	//sorting using inbuilt functions
	sort.Ints(arr)
	sort.Strings(arr2)
	fmt.Println("After Sorting------------------")
	fmt.Println(arr)
	fmt.Println(arr2)

}
