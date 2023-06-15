package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type row1 struct {
	Roll   int
	Fname  string
	Lname  string
	Gender string
	Result string
	Mark   int
}

func reqHandler(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "roshin:roshin@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		fmt.Fprintln(res, "Connection Error", err.Error())
	} else if pingerr := db.Ping(); pingerr != nil {
		fmt.Fprintln(res, "Ping Error", err.Error())
	} else {
		defer db.Close()
		sl := make([]row1, 0)
		result, err := db.Query("SELECT * FROM student")
		if err != nil {
			fmt.Fprintln(res, "Query error", err.Error())
		}
		for result.Next() {
			var r row1
			if err = result.Scan(&r.Roll, &r.Fname, &r.Lname, &r.Gender, &r.Result, &r.Mark); err != nil {
				fmt.Fprintln(res, "Result Scanning Failed", err.Error())
			}
			sl = append(sl, r)
		}
		for i, val := range sl {
			fmt.Fprintln(res, i+1, "->", val)
		}
	}
}

func main() {
	http.HandleFunc("/go", reqHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listener Failed:::", err.Error())
	}

}
