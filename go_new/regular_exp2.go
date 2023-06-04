package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	var check, str string
	str = "Roshin"
	fmt.Println("ENter chcking string ::: ~regexp(Roshin)")
	fmt.Scanf("%s", &check)

	rob, err := regexp.Compile(str)
	if err != nil {
		log.Fatal(":::::::::::", err)
	}

	fmt.Println(check)
	fmt.Println("first expression match::::", rob.MatchString(check))
	fmt.Println("second find string::", rob.FindString(check))
	fmt.Print("Index of s  :", rob.FindStringIndex("sh"))

	//new way to check start and end only to check match
	r, _ := regexp.Compile("p([a-z]+)ch") //only constraint is it should start with "p" and end in "ch" and anything can be in between

	fmt.Println("\n\nThird::::::::::::::::::::::::::::::::::::::::::", r.MatchString("peeeeeekmmeeeech"))

}
