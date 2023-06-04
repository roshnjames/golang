package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Float64())
	s2 := rand.NewSource(6)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",", r2.Intn(100))
}
