package utils

import(
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDatabase(){
	var err error
	db, err = sql.Open("mysql", "mombewa:Mo201nrb#@tcp(127.0.0.1:3306)/aircraft")
	if err != nil {
		log.Fatal(err)
	}
}