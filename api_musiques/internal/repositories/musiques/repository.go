package musiques

import (
	"tchipify/musiques/internal/helpers"
	"tchipify/musiques/internal/models"

	"github.com/gofrs/uuid"
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
		err = rows.Scan(&data.Id, &data.Title, &data.Artist, &data.Filename, &data.Published)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	_ = rows.Close()

	return songs, err

}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDb()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM musiques WHERE id=?", id)
	helpers.CloseDb(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Title, &song.Artist, &song.Filename, &song.Published)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func AddSong(song models.Song) error {
	db, err := helpers.OpenDb()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO musiques (id, title, artist, filename, published) VALUES  (?,?,?,?,?)", song.Id, song.Title, song.Artist, song.Filename, song.Published)

	helpers.CloseDb(db)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSong(songId uuid.UUID) error {
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

func Update(idSong uuid.UUID, song models.Song) error {
	db, err := helpers.OpenDb()
	if err != nil {
		logrus.Errorf("Error when opening DB: %s", err.Error())
		return err
	}
	res, err := db.Exec("UPDATE musiques SET artist = ?, title= ?, filename = ?,  published = ? WHERE id = ?", song.Artist, song.Title, song.Filename, song.Published, idSong)

	rows, err := res.RowsAffected()
	logrus.Printf("%d", rows)
	helpers.CloseDb(db)
	if err != nil {
		logrus.Errorf("Repository : Error in updating song %s", err.Error())
		return err
	}
	return nil
}
