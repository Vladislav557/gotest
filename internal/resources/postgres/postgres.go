package postgres

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

type Database struct {
    DB *sql.DB
}

func (db *Database) Init(url string) {
    var err error
    db.DB, err = sql.Open("postgres", url)
    if err != nil {
        log.Fatal(err)
    }
    if err = db.DB.Ping(); err != nil {
        log.Fatal(err)
    }
}

func (db *Database) Close() {
    if err := db.DB.Close(); err != nil {
        log.Fatal(err)
    }
}