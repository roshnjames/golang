package main

import "fmt"

type i interface {
	method() float64
}

type first struct {
	num int
}

type second struct {
	first
	val int
}

/*func (f first) method() float64 {
	return float64(f.num)
}*/

func (s second) method() float64 {
	return float64(s.first.num * s.val)
}

func main() {
	x := second{first: first{num: 2}, val: 3}
	fmt.Println(x.method())
}

//embedding a structure inside another structure ...similar to inheritance in other programming languages
//first->struct 1
//second->struct 2
