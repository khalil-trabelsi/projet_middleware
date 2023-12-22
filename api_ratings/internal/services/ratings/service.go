package ratings

import (
	"database/sql"
	"errors"
	"net/http"
	"tchipify/ratings/internal/models"
	repository "tchipify/ratings/internal/repositories/ratings"

	"github.com/sirupsen/logrus"
)

func GetAllRatings() ([]models.Rating, error) {
	var err error
	// calling repository
	ratings, err := repository.GetAllRatings()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return ratings, nil
}

func GetRatingById(id int) (*models.Rating, error) {
	rating, err := repository.GetRatingById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "rating not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving rating : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rating, err
}

func CreateRating(rating models.Rating) (int64, error) {
	id, err := repository.CreateRating(rating)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du rating : %s", err.Error())
		return 0, &models.CustomError{
			Message: "Something went wrong",
			Code:    422,
		}
	}
	return id, nil
}

func DeleteRating(ratingID int) error {
	err := repository.DeleteRating(ratingID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du commentaire : %s", err.Error())
		return err
	}

	return nil
}

func UpdateRating(ratingID int, song models.Rating) error {
	err := repository.UpdateRating(ratingID, song)
	if err != nil {
		logrus.Errorf("error updating ressource : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return nil
}
