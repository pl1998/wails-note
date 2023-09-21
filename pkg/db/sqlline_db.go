package db

import (
	"changeme/pkg/helpers"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Conn() *sql.DB {
	DB, err := sql.Open("sqlite3", "databases/note.db")
	helpers.CheckErr(err)
	return DB
}
