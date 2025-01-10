package postgreSqlDb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConnect struct {
	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	DBNAME   string
}

type PostgreSqlDb struct {
	db *sql.DB
}

func New(conn *DBConnect) *PostgreSqlDb {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conn.HOST, conn.PORT, conn.USER, conn.PASSWORD, conn.DBNAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return &PostgreSqlDb{db: db}
}
