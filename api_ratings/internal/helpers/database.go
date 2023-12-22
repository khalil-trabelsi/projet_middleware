package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:ratings.db")
	if err != nil {
		db.SetMaxOpenConns(1)
		logrus.Fatal(err)
	}

	logrus.Info("Connected to the Db")

	return db, err

}

func CloseDb(db *sql.DB) {

	err := db.Close()
	if err != nil {
		logrus.Error("error closing db : %s", err.Error())
	}

}
