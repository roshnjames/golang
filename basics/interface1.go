package main

import "fmt"

type motorv interface {
	mileage() float64
}

type car struct {
	speed int
	dist  int
}

type bike struct {
	speed int
	dist  int
}

func (c car) mileage() float64 {
	return float64(c.dist / c.speed)
}

func (b bike) mileage() float64 {
	return float64(b.dist / b.speed)
}

func main() {
	obj1 := car{20, 184}
	obj2 := bike{10, 200}
	fmt.Printf("\nCar mileage is %f  ", obj1.mileage())
	fmt.Printf("\nBike mileage is %f", obj2.mileage())

}

/*interface syntax======>eg:motorv

type interface_name interface{
	method_name() return_type
}

implicit implementation with a function======>no explicit implementation like java using implement keyword

func (variable struct_name) interface_method_name() interface_return_type{
	return //code_goes_here
}

*/
