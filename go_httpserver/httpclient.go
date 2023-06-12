package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://gobyexample.com/http-client")
	if err != nil {
		log.Fatal(":::::::::", err)
	}
	fmt.Println("status::", res.Status)
	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 11; i++ {
		fmt.Println(scanner.Text())
		//fmt.Println("-------------------", string(scanner.Bytes()))
	}

}
