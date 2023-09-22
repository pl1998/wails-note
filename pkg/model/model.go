package model

import (
	"changeme/pkg/helpers"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB sqlite db conn
var DB *gorm.DB

// err
var err error

// Conn get sqlite database conn
func Conn() *gorm.DB {
	DB, err = gorm.Open(sqlite.Open("note.db"), &gorm.Config{})
	helpers.CheckErr(err)
	return DB
}
