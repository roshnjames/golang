package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server creation at port 8080")

	http.HandleFunc("/hello", func(res http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(res, "hello Roshin James😊")
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(":::::::::", err)

	}
	fmt.Println("server created successfully at port 8080😊♨️")
}
