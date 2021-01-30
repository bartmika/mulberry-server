package db

import (
	"github.com/sdomino/scribble"
)

func ConnectDB() (*scribble.Driver, error) {
	// The location of our db.
	dir := "./my_database"

	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err := scribble.New(dir, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}
