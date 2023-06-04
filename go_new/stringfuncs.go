package main

import (
	"fmt"
	s "strings"
)

func main() {
	pr := fmt.Println

	//pr(len("roshin james"))
	str := "roshin"
	var str1 string
	pr(s.ToUpper("roshin james"))
	pr(s.ToLower("ROSHIN JAMES"))
	pr(s.Index("roshin", "s"))
	pr(s.Split("roshin-j-a-mes", "-"))
	str1 = s.Clone(str)
	pr(str1)
	pr(s.Contains(str, "r"))
	pr(s.Count(str, "n"))
	pr(s.Replace(str, "r", "j", 1))
	pr(s.Repeat(str+" ", 5))

}
