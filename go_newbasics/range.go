package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5

	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	fmt.Println("ARRAY----------------------------------------------")

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for val := range arr {
		fmt.Print(arr[val], "\t\a")
	}

	by, _ := fmt.Println("\nSLICE__________________________________________")
	fmt.Println(by, " bytes printed")
	s := make([]int, 11)
	for i := 0; i < 10; i++ {
		s = append(s[:i], i)
	}
	fmt.Println()
	for v := range s {
		fmt.Print(s[v], "\t")
	}

}

//range is similar to range in python
//m=map[]/[]slice/[20]array
//for _,i:=range m{//}

//  _,err=fmt.Println() ==>'_' character discards the first value returned
//range generally returns two values depdending on the data structure on which we are applying it
//in case of arrays and slices range returns index as the first return value and the value as the second return value.
//we can choose to recieve our required value by discarding the unnecessary value using the '_' operator
