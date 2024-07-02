package latihan_golang_database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
)

//func TestOpenConnection(t *testing.T) {
//	urlExample := "postgres://nikawaweb:Sabeso76@localhost:5432/latihan_golang"
//	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
//	conn, err := pgx.Connect(context.Background(), urlExample)
//	if err != nil {
//		fmt.Printf("Unable to connect to database: %v\n\n", err)
//		os.Exit(1)
//	}
//	defer func(conn *pgx.Conn, ctx context.Context) {
//		err := conn.Close(ctx)
//		if err != nil {
//			fmt.Println("Error", err.Error())
//		}
//	}(conn, context.Background())
//}

func TestOpenConnectionPzn(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://nikawaweb:Sabeso76@localhost:5432/latihan_golang?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)
}
