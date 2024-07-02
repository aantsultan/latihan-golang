package app

import (
	"database/sql"
	"latihan-golang-restful-api-nkw/helper"
	"time"
)

func NewDB() *sql.DB {
	username := helper.GetProperties("datasource.username")
	password := helper.GetProperties("datasource.password")
	url := "postgres://" + username + ":" + password + "@localhost:5432/nikawamobile?sslmode=disable"
	db, err := sql.Open("postgres", url)
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
