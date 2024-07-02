package latihan_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// ExecContext, dapat membatalkan query dan tidak direkomendasikan untuk query select
	script := "INSERT INTO customer(id, name) VALUES ('elaina', 'Elaina')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert data")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name from customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id, ", name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at from customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("==============")
		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		if birthDate.Valid {
			fmt.Println("birthDate:", birthDate.Time)
		}
		fmt.Println("married:", married)
		fmt.Println("createdAt:", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// input yang benar
	//username := "admin"
	//password := "admin"

	// sql injection
	username := "admin' ; --"
	password := "salah"

	script := "SELECT username FROM \"user\" usr WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("success login", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// input yang benar
	username := "admin"
	password := "admin"

	// sql injection
	//username := "admin' ; --"
	//password := "salah"

	// postgres pakai $1,
	// mysql pakai ?
	script := "SELECT username FROM \"user\" WHERE username = $1 AND password = $2 LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("success login", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "kiana"
	password := "password"

	// ExecContext, dapat membatalkan query dan tidak direkomendasikan untuk query select
	script := "INSERT INTO \"user\"(username, password) VALUES ($1, $2)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert data")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "kiana@mail.com"
	comments := "halo captain"

	// ExecContext, dapat membatalkan query dan tidak direkomendasikan untuk query select
	script := "INSERT INTO comments(email, comments) VALUES ($1, $2)"
	result, err := db.ExecContext(ctx, script, email, comments)
	if err != nil {
		panic(err)
	}

	resultId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("success insert data, id :", resultId) // tidak support untuk postgres
}
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments (email, comments) VALUES ($1, $2)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "kiana" + strconv.Itoa(i) + "@mail.com"
		comments := "komentar ke : " + strconv.Itoa(i)

		_, err := statement.ExecContext(ctx, email, comments)
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment email : ", email)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	script := "INSERT INTO comments(email, comments) VALUES ($1, $2)"

	for i := 0; i < 10; i++ {
		email := "kiana" + strconv.Itoa(i) + "@mail.com"
		comments := "komentar ke : " + strconv.Itoa(i)

		_, err := tx.ExecContext(ctx, script, email, comments)
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment email : ", email)
	}

	//err = tx.Commit()
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
