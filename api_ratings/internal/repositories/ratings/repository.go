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

func GetRatingById(id int) (*models.Rating, error) {
	db, err := helpers.OpenDb()
	if err != nil {

		return nil, err
	}
	row := db.QueryRow("SELECT * FROM ratings WHERE id=?", id)
	helpers.CloseDb(db)

	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.SongId, &rating.UserId, &rating.Comment, &rating.Note)

	if err != nil {

		return nil, err // Autres erreurs lors du scan
	}
	return &rating, err
}
