package main

import "fmt"

type student struct {
	roll int
	name string
	mark float32
}

func main() {
	var n int
	fmt.Println("ENter no. of students")
	fmt.Scanln(&n)
	var arr [5]student
	for i := 0; i < n; i++ {
		fmt.Println("ENter roll")
		fmt.Scanln(&arr[i].roll)
		fmt.Println("ENter name")
		fmt.Scanln(&arr[i].name)
		fmt.Println("ENter mark")
		fmt.Scanln(&arr[i].mark)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("roll:%d\tname:%s\tmark:%.0f\n", arr[i].roll, arr[i].name, arr[i].mark)
	}
}

//array of structure
