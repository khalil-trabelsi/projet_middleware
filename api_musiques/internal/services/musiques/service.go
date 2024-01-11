package musiques

import (
	"database/sql"
	"errors"
	"net/http"
	"tchipify/musiques/internal/models"
	repository "tchipify/musiques/internal/repositories/musiques"

	"github.com/gofrs/uuid"
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

func GetSongById(id uuid.UUID) (*models.Song, error) {
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

func AddSong(song models.Song) error {
	err := repository.AddSong(song)
	if err != nil {
		logrus.Errorf("error creating ressource : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    422,
		}
	}

	return nil
}

func DeleteSong(idSong uuid.UUID) error {
	err := repository.DeleteSong(idSong)
	if err != nil {
		logrus.Errorf("error deleting ressource : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return nil
}

func UpdateSong(idSong uuid.UUID, song models.Song) error {
	err := repository.Update(idSong, song)
	if err != nil {
		logrus.Errorf("error updating ressource : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return nil
}
