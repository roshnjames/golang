// pointer in go
package main

import "fmt"

func main() {
	var ptr *int //pointer declaration
	var x, y int = 3, 2
	fmt.Printf("initial value is %d\n", y)
	pt := &x
	ptr = &y
	*ptr = 4
	pro := x * *ptr
	fmt.Println(pt)
	fmt.Printf("%v", ptr)
	fmt.Printf("\n%d is the new value\n", y)
	fmt.Println(pro)
}
