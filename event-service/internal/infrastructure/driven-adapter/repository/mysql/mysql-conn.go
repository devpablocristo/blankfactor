package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func RepositoryConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", openConn())
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return db, nil
}

func openConn() string {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", USER, PASS, HOST, PORT, DBNAME)
}
