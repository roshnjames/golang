package main

import "fmt"

func alias() {
	panic("mone error⚠️⚠️⚠️")
}

func recovery() {
	r := recover()
	if r != nil {
		fmt.Println("error catched by recover from panic:::::", r)
	}
}

func main() {

	//fmt.Println(recover())

	defer recovery()
	alias()

}
