package app

import (
	"database/sql"
	"latihan-golang-depedency-injection/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Sabeso76@tcp(localhost:3306)/latihan_golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
