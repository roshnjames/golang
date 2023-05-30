package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	t := template.New("t1")
	t1 := template.Must(t.Parse("Value: {{.}}\n")) //whatever data is passed from execute() will get displayed at {{.}}
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
	fmt.Println("\n-++++\n\n")

	tel, err := template.ParseFiles("rosh.txt", "rosh1.txt")
	if err != nil {
		template.Must(tel, err)
	} else {
		fmt.Println("ENetered")
		name := "roshin"
		err := tel.ExecuteTemplate(os.Stdout, "rosh.txt", name)
		err1 := tel.ExecuteTemplate(os.Stdout, "rosh1.txt", nil)
		if err != nil || err1 != nil {
			log.Fatal("error from log ", err)
		}
	}
}
