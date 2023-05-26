package main

import (
	"fmt"
	"reflect"
)

func main() {
	var letter rune
	letter = 'B'
	fmt.Printf("type :%T \nvalue  :%v\n-------------------------------", letter, letter)

	//single literals are considered as rune by default in golang
	x := 'r'
	fmt.Println("\n", reflect.TypeOf(x))
	fmt.Printf("\n%T ---%v", x, x)
}

//a RUNE is a datatype to represent literals in go and holds the unicode value of the character in int32
//x:='y'===>variable x is of datatype rune and stores unicode value of character 'y' in int32 format
//'' quotes are used to specify literals for rune datatype
