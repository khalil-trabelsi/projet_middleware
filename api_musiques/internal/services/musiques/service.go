package musiques

import (
	"database/sql"
	"errors"
	"net/http"
	"tchipify/musiques/internal/models"
	repository "tchipify/musiques/internal/repositories/musiques"

	_ "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllSongs() ([]models.Song, error) {
	var err error

	songs, err := repository.GetAllSongs()
	if err != nil {
		logrus.Error("cannot retrieving songs : ", err.Error())
		return nil, &models.CustomError{
			Message: "something went wrong",
			Code:    500,
		}
	}
	return songs, nil
}

func GetSongById(id int) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

func AddSong(song models.Song) (int64, error) {
	id, err := repository.AddSong(song)
	if err != nil {
		logrus.Errorf("error create ressource : %s", err.Error())
		return 0, &models.CustomError{
			Message: "Something went wrong",
			Code:    422,
		}
	}

	return id, nil
}
