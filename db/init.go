package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

//var db *sql.DB

func GetDB() *sql.DB {
    dbEntity, err := sql.Open("sqlite3", `D:\Code\Go\MyProj\JDBot\local_db\database.sqlite`)
    if err != nil {
        log.Panicln("opening database error:", err)
    }
    return dbEntity
}
