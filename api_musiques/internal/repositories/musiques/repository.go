package musiques

import (
	"tchipify/musiques/internal/helpers"
	"tchipify/musiques/internal/models"
	_ "github.com/gofrs/uuid"
	
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDb()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select * from musiques")
	helpers.CloseDb(db)
	if err != nil {
		return nil, err
	}

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


func GetSongById(id int) (*models.Song, error) {
	db, err := helpers.OpenDb()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM musiques WHERE id=?", id)
	helpers.CloseDb(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.ArtistName, &song.Title, &song.DurationInMillis)
	if err != nil {
		return nil, err
	}
	return &song, err
}


func AddSong(song models.Song) (int64, error) {
	db, err := helpers.OpenDb()
	if err != nil {
		return 0, err
	}
	result, err := db.Exec("INSERT INTO musiques (id, artistName, title, durationInMillis) VALUES  (?, ?, ?, ?)", song.Id, song.ArtistName,song.Title, song.DurationInMillis)
	helpers.CloseDb(db)
	if (err != nil) {
		return 0, err
	}
	id,  err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}