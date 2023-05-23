package main

import "fmt"

func main() {

	sl := make([]int, 3, 10)
	fmt.Println(sl)
	sl1 := append(sl, 1)
	fmt.Println(sl1)
	sl1[3] = 0
	fmt.Println(sl1)

	fmt.Printf("\nthe maximum capacity is :%d", cap(sl1))
	fmt.Printf("\nThe length of the created slice after change is %d \nand before change is %d", len(sl1), len(sl))

	var by int
	var err error
	by, err = fmt.Println("\n", sl1[1:3])

	fmt.Println("bytes printed is:", by, err)

}

//format==var_name:=make([]data_type,length(optional),capacity)
//cap()-finds max capacity
//len()
//new_var_name=append(var_name,item)~returns a new slice
