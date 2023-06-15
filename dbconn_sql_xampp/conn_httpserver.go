package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handling(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "roshin:roshin@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		fmt.Fprintln(res, "::::", err)
	} else if pingerr := db.Ping(); pingerr != nil {
		fmt.Fprintln(res, "pingerror occured")
	} else {
		fmt.Fprintln(res, "connected successfully to 'GODB'")
	}
	defer db.Close()
}

func main() {
	http.HandleFunc("/conn", handling)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("listen:::::", err)
	}

}
