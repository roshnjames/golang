package main

import (
	"fmt"
	"net/url"
)

func main() {
	link := "https://github.com/roshnjames/golang"
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	fmt.Println("user :", u.User.Username())
	fmt.Println("Host :", u.Host)
	fmt.Println("Port :", u.Port())
	fmt.Println("Path :", u.Path)
	fmt.Println("Fragment :", u.Fragment)
	fmt.Println("scheme :", u.Scheme)
}
