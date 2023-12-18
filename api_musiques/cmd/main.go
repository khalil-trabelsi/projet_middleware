package main

import (
	"net/http"
	"tchipify/musiques/internal/controllers/musiques"
	"tchipify/musiques/internal/helpers"
	_ "tchipify/musiques/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api/v1/songs", func(r chi.Router) {
		r.Get("/", musiques.GetSongs)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(musiques.Ctx)
			r.Get("/", musiques.GetSong)
		})
		r.Post("/", musiques.AddSong)
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
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
