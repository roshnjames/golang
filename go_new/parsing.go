package main

import (
	"fmt"
	"strconv"
)

// string to other datatype conversion
func main() {
	fmt.Println(strconv.ParseInt("52", 32, 12))
	//
	str := "43.567"
	fl, _ := strconv.ParseFloat(str, 64)
	fmt.Println(fl)
	fmt.Printf("type of str -%T :::Type of fl -%T\n", str, fl)
	//
	fmt.Println(strconv.Atoi("521")) //base10 int conversion
	fmt.Println(strconv.ParseBool("False"))
}
