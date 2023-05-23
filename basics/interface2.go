package main

import "fmt"

const pi float64 = 3.14

//interface
type shape interface {
	area() float64
	perimeter() float64
	sum() string
}

//structures
type circle struct {
	radius float32
}

type triangle struct {
	base   float32
	height float32
}

type rectangle struct {
	length  float32
	breadth float32
}

//interface implicit implementation

//circle------------------------------------------
func (c circle) area() float64 {
	return pi * float64(c.radius*c.radius)
}

func (c circle) perimeter() float64 {
	return pi * 2 * float64(c.radius)
}

func (c circle) sum() string {
	str := fmt.Sprint(c.area() + c.perimeter())
	return str
}

//triangle

func (t triangle) area() float64 {
	return 0.5 * float64(t.base*t.height)
}

func (t triangle) perimeter() float64 {
	return float64(t.height * t.height * t.height)
}

func (t triangle) sum() string {
	str := fmt.Sprint(t.area() + t.perimeter())
	return str
}

//rectangle

func (r rectangle) area() float64 {
	return float64(r.length * r.breadth)
}

func (r rectangle) perimeter() float64 {
	return 2 * float64(r.length*r.breadth)
}

func (r rectangle) sum() string {
	str := fmt.Sprint(r.area() + r.perimeter())
	return str
}

//main function
func main() {
	obj1 := circle{3.00}
	obj2 := triangle{3.00, 4.00}
	obj3 := rectangle{3.00, 4.00}

	fmt.Printf("\n\nThe area of circle is %f", obj1.area())
	fmt.Printf("\nThe perimeter of circle is %f", obj1.perimeter())
	fmt.Printf("\n---The sum in type(%T) for circle is %v ", obj1.sum(), obj1.sum())
	fmt.Printf("\n\n\nThe area of triangle is %f", obj2.area())
	fmt.Printf("\nThe perimeter of triangle is %f", obj2.perimeter())
	fmt.Printf("\n---The sum in type:(%T) for triangle is %s", obj2.sum(), obj2.sum())
	fmt.Printf("\n\n\nThe area of rectangle is %f", obj3.area())
	fmt.Printf("\nThe perimeter of rectangle is %f", obj3.perimeter())
	fmt.Printf("\n---The sum in type:(%T) for rectangle is %s", obj3.sum(), obj3.sum())
}
