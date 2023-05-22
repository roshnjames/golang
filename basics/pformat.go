package main

import "fmt"

func main() {
	var x, by int
	var str string = "ROshin"
	var b bool
	var f float64 = 2.044
	var ptr *float64 = &f
	x = 17
	var err error
	by, err = fmt.Printf("%T  is data type of x", x)
	fmt.Println(by)
	if err != nil {
		fmt.Print("Err from is ", err)
	}
	fmt.Print("Err from is ", err)

	fmt.Println("\n\nPrinting variable formats")
	fmt.Printf("\n\nString:%s \nboolean:%v \nfloat:%f, %g \nbinary of x :%b\n base 8 of x:%o\n Pointer format is v:%v", str, b, f, f, x, x, ptr)
	fmt.Printf("\nv can print anything:%v", x)
}
