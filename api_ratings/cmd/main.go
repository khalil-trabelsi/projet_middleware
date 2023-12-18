package main

import (
	_ "net/http"
	"tchipify/ratings/internal/helpers"
	_ "tchipify/ratings/internal/models"

	_ "github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {

}

func init() {
	db, err := helpers.OpenDb()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS musiques (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			artistName VARCHAR(255) NOT NULL,
			title VARCHAR(255) NOT NULL,
			durationInMillis INT NOT NUll
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDb(db)
}
