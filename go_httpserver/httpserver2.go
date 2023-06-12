package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func server(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/server" {
		er := errors.New("Wrong Path")
		fmt.Fprintln(res, er, "⚠️")
	} else if req.Method != "GET" {
		er := errors.New("Wrong request Method")
		fmt.Fprintln(res, er, "⚠️")
	} else {
		fmt.Fprintln(res, "Hello Roshin James😀... Server Worked")
	}
}
func main() {

	http.HandleFunc("/server", server)
	//------
	fmt.Println("Server creation and error handling at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Listening Failed:::::::", err)
	}
}
