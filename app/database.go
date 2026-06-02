package app

import (
	"database/sql"
	"time"

	"github.com/abdul452/belajar-golang-restful-api/helper"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(192.168.56.128:3306)/belajar_golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
