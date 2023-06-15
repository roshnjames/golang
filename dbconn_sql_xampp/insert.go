package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if db, err := sql.Open("mysql", "roshin:roshin@tcp(127.0.0.1:3306)/godb"); err != nil {
		log.Fatal("::::::", err)
	} else if pingerr := db.Ping(); pingerr != nil {
		log.Fatal("pingerr::::::", pingerr)
	} else {
		defer db.Close()
		_, err = db.Query("INSERT INTO student(rollno,fname,lname,gender,result,mark)VALUES('','mrithul','m','male','pass',39)")
		if err != nil {
			log.Fatal("insert err::::::", err)
		}
		fmt.Println("Insertion Successfull👍")
	}

}
