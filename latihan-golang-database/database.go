package latihan_golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	url := "postgres://nikawaweb:Sabeso76@localhost:5432/latihan_golang?sslmode=disable"
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err.Error())
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
