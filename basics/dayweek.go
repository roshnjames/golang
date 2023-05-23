package main

import "fmt"

func main() {
	var num int
	var day string
	fmt.Println("Enter the no. from 1-7")
	fmt.Scanf("%d\n", &num)
	if num < 1 {
		fmt.Println("Value should be greater than 1")
	} else if num > 7 {
		fmt.Println("Value should be below 7")
	} else {
		switch num {
		case 1:
			day = "Sunday"
		case 2:
			day = "Monday"
		case 3:
			day = "Tuesday"
		case 4:
			day = "Wednesday"
		case 5:
			day = "Thursday"
		case 6:
			day = "Friday"
		case 7:
			day = "Saturday"
		default:
			day = "Invalid"
		}
		fmt.Printf("%v", day)
	}

}
