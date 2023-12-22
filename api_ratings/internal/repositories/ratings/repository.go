package musiques

import (
	"tchipify/ratings/internal/helpers"
	"tchipify/ratings/internal/models"

	_ "github.com/sirupsen/logrus"
)

func GetAllRatings() ([]models.Rating, error) {
	db, err := helpers.OpenDb()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings")
	helpers.CloseDb(db)
	if err != nil {
		return nil, err
	}

	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id, &data.UserId, &data.SongId, &data.Comment, &data.Note)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, data)
	}

	_ = rows.Close()

	return ratings, err
}
