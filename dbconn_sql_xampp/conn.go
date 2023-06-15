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
		log.Fatal(":::::::::::", err)
	}
	defer db.Close()
	if pingerr := db.Ping(); pingerr != nil {
		log.Fatal("ping:::::::::::", pingerr)
	}
	fmt.Println("\n---------------------------------------------------\nConnected successfully to 'godb' database in xampp |\n---------------------------------------------------\n")
}
