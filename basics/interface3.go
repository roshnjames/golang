package main

import "fmt"

// Creating an interface
type tank interface {

	// Methods
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

// Implementing methods of
// the tank interface
func (m myvalue) Tarea() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

// Main Method
func main() {

	// Accessing elements of
	// the tank interface

	obj1 := myvalue{10, 14}
	fmt.Println("Area of tank :", obj1.Tarea())
	fmt.Println("Volume of tank:", obj1.Volume())
	//another way to access the interface methods using the variables of interface type
	var t tank
	t = myvalue{10, 14}
	fmt.Println("\n\nUsing the variable of interface type itself")
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())

}
