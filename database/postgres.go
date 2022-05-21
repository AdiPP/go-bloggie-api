package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Open() *sqlx.DB {
	//Init DB
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SCHEMA"),
	)
	dbConn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return dbConn
}

func CreateSchema() (sql.Result, error) {
	//Init DB
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	dbConn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return dbConn.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", os.Getenv("DB_SCHEMA")))
}
