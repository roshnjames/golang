package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "roshin:roshin@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		log.Fatal("conn::::::::", err.Error())
	} else if pingerr := db.Ping(); pingerr != nil {
		log.Fatal("ping::::", pingerr.Error())
	} else {
		defer db.Close()
		roll := 4
		_, err = db.Query("UPDATE student SET mark=4 WHERE rollno=?", roll)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Updated successfully👍")

	}
}
