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
		log.Fatal("::::::", err.Error())
	} else if pingerr := db.Ping(); pingerr != nil {
		log.Fatal(":::ping:::", pingerr.Error())
	} else {
		defer db.Close()
		roll := 5
		//_,err=db.Query("DELETE FROM student WHERE rollno=?", roll)
		_, err = db.Exec("DELETE FROM student WHERE rollno=?", roll)
		if err != nil {
			log.Fatal(":::deletion query:::", err.Error())
		}
		fmt.Println("Deleted Successfully👍")
	}
}
