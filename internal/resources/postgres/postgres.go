package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Init(url string) {
	var err error
	DB, err = sql.Open("postgres", url)
	Migrate()
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}
