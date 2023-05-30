package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	//t := template.New("t")
	tmp, err := template.ParseFiles("text_template.go")
	if err != nil {
		template.Must(tmp, err)
	} else {
		fmt.Println("\ninside----------\n")
		err := tmp.Execute(os.Stdout, nil)
		if err != nil {
			log.Fatal("logged:::::", err)
		}
		tmp, err = template.ParseFiles("rosh.txt")
		if err == nil {
			err = tmp.Execute(os.Stdout, "roshin")

		}
	}

}
