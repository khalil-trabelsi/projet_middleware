package ratings

import (
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
