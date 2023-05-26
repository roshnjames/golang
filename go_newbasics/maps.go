package main

import "fmt"

func main() {
	fmt.Println("MAPS-same as dict in python\n-------------------------------------------------------")
	m := map[string]int{}
	m["city"] = 5
	m["real"] = 1
	fmt.Println(m)

	//another way with make()
	m1 := make(map[string]int)
	m1["united"] = 1
	m1["chelsea"] = 3
	fmt.Println("----------------------------------------------------------\n", m1)
	fmt.Println(m1["chelsea"], "---", m["real"])

	delete(m1, "chelsea")

	value, presence := m1["chelsea"]
	fmt.Println(presence, value)
}

//a map should onlycontain same dataype of keys and same datatype of values
//map[string]int
//key->string,,,,,value->int
