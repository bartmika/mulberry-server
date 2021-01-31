package db

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		databaseHost,
		databasePort,
		databaseUser,
		databasePassword,
		databaseName,
	)

	dbInstance, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = dbInstance.Ping()
	if err != nil {
		return nil, err
	}
	return dbInstance, nil
}
