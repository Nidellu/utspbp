package controller

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func connect() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	fmt.Println(dbHost)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_uts_pbp")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
