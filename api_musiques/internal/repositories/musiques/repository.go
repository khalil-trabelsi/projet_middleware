package musiques

import (
	"tchipify/musiques/internal/helpers"
	"tchipify/musiques/internal/models"

	_ "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
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

	result, err := db.Exec("INSERT INTO musiques (artistName, title, durationInMillis) VALUES  ( ?, ?, ?)", song.ArtistName, song.Title, song.DurationInMillis)

	helpers.CloseDb(db)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteSong(songId int) error {
	db, err := helpers.OpenDb()
	if err != nil {
		return err
	}
	logrus.Warn("", songId)
	res, err := db.Exec("DELETE from musiques WHERE id = ?", songId)

	rows, err := res.RowsAffected()
	logrus.Printf("%d", rows)
	helpers.CloseDb(db)
	if err != nil {
		return err
	}
	return nil
}

func Update(idSong int, song models.Song) error {
	db, err := helpers.OpenDb()
	if err != nil {
		logrus.Errorf("Error when opening DB: %s", err.Error())
		return err
	}
	res, err := db.Exec("UPDATE musiques SET artistName = ?, title= ?, durationInMillis = ? WHERE id = ?", song.ArtistName, song.Title, song.DurationInMillis, idSong)

	rows, err := res.RowsAffected()
	logrus.Printf("%d", rows)
	helpers.CloseDb(db)
	if err != nil {
		logrus.Errorf("Repository : Error in updating song %s", err.Error())
		return err
	}
	return nil
}
