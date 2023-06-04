package main

import (
	"errors"
	"fmt"
	"regexp"
)

func main() {
	str := "Roshin James"
	fmt.Println("ENter chcking string ::: ~regexp(Roshin James)")
	var check string
	fmt.Scanln(&check)

	//
	rexp, err := regexp.MatchString(check, str)
	if err != nil {
		fmt.Sprintln("error :::", errors.New("failed to match"))
	}
	fmt.Println("first regexp check : ", rexp)
}
