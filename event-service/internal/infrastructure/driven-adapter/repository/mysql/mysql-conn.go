package mysql

import (
	"database/sql"
	"log"
)

func RepositoryConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:secret@tcp(event-msql-repo:3306)/events_service")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
