package main

import (
	"net/http"
	"tchipify/ratings/internal/controllers/ratings"
	"tchipify/ratings/internal/helpers"
	_ "tchipify/ratings/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	router := chi.NewRouter()
	// Our routes
	router.Get("/ratings", ratings.GetRatings)
	router.Get("/ratings/{id}", ratings.GetRating)
	// router.Put("/ratings/{id}", ratings.UpdateRating)
	// router.Delete("/ratings/{id}", ratings.DeleteRating)
	logrus.Info("[INFO] Web server started. Now listening on *:8084")
	logrus.Fatalln(http.ListenAndServe(":8084", router))
}

func init() {
	db, err := helpers.OpenDb()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS ratings (
			id INTEGER  PRIMARY KEY AUTOINCREMENT,
			song_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            comment VARCHAR(300) NOT NULL,
			rating REAL CHECK (rating BETWEEN 0 AND 5)
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDb(db)
}
