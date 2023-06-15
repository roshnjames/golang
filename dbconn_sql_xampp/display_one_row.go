package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type row struct {
	Roll   int
	Fname  string
	Lname  string
	Gender string
	Result string
	Mark   int
}

func main() {
	db, err := sql.Open("mysql", "roshin:roshin@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		log.Fatal(err.Error())
	} else if pingerr := db.Ping(); pingerr != nil {
		log.Fatal(pingerr.Error())
	} else {
		defer db.Close()
		roll := 1
		var r row
		res := db.QueryRow("SELECT * FROM student WHERE rollno=?", roll)
		if res == nil {
			fmt.Println("empty result")
		}
		err = res.Scan(&r.Roll, &r.Fname, &r.Lname, &r.Gender, &r.Result, &r.Mark)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("OUTPUT ::\n", r)
	}
}
