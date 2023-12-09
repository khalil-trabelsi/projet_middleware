package musiques

import (
	"github.com/gofrs/uuid"
	"tchipify/musiques/helpers"
	"tchipify/musiques/models"
)

func getAllSongs() {
	db, err := helpers.OpenDb()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select * from musiques")
	helpers.CloseDb()

	songs := []models.Song{}

	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.ArtistName, &data.Title, &data.DurationInMillis)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	_ = rows.Close()

	return songs, err


}