package main

import (
	"fmt"
	"html/template"
	"os"
)

type data struct {
	Name string
	Str  string
	Lang string
}

func main() {
	tmp, err := template.ParseFiles("rosh.html")
	if err != nil {
		template.Must(tmp, err)
	}
	dt := data{Name: "roshin", Str: "this is html template testing", Lang: "GOlang"}
	err = tmp.Execute(os.Stdout, dt)
	if err != nil {
		fmt.Println("\n\n:::::::::")
		panic(err)
	}

}
