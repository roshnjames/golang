package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type rows struct {
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
		log.Fatal("ping:::", err.Error())
	} else {
		fmt.Println("databse connected successfuly👍")
		defer db.Close()

		res, err := db.Query("SELECT * FROM student")
		sl := make([]rows, 0)
		if err != nil {
			log.Fatal("query failure::::", err.Error())
		}
		for res.Next() {
			var r rows
			err = res.Scan(&r.Roll, &r.Fname, &r.Lname, &r.Gender, &r.Result, &r.Mark)
			if err != nil {
				log.Fatal("scan error:::", err.Error())
			}
			sl = append(sl, r)
		}

		for i, val := range sl {
			fmt.Println(i+1, ":::", val)
		}
	}
}
