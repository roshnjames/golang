package main

import (
	"errors"
	"fmt"
)

func check(s []int) ([]int, []string) {
	ans := make([]int, 10)
	er := make([]string, 10)
	for j, i := range s {
		if i%2 == 0 {
			ans = append(ans[:j], i)
			er = append(er, ".....no error")
		} else {
			ans = append(ans[:j], 0000)
			er = append(er, fmt.Sprint("\n", errors.New("odd number encountered")))
		}

	}
	return ans, er
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans, err := check(s)
	fmt.Print("\nans:", ans)
	fmt.Print("\nerrors :", err)
}
