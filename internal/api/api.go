package api

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"price/internal/db/postgre_sql_db"
	"strconv"
)

func mustDBConnect() *postgreSqlDb.DBConnect {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Не найден файл .env")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Порт не порт")
	}
	return &postgreSqlDb.DBConnect{
		HOST:     os.Getenv("DB_HOST_OUTSIDE"),
		DBNAME:   os.Getenv("DB_NAME"),
		PORT:     port,
		USER:     os.Getenv("USER_NAME"),
		PASSWORD: os.Getenv("USER_PASSWORD"),
	}
}

var Storage = postgreSqlDb.New(mustDBConnect())
