package musiques

import (
	"tchipify/ratings/internal/helpers"
	"tchipify/ratings/internal/models"

	"github.com/sirupsen/logrus"
)

func GetAllRatings() ([]models.Rating, error) {
	db, err := helpers.OpenDb()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM rating")
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
	row := db.QueryRow("SELECT * FROM rating WHERE id=?", id)
	helpers.CloseDb(db)

	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.SongId, &rating.UserId, &rating.Comment, &rating.Note)

	if err != nil {

		return nil, err // Autres erreurs lors du scan
	}
	return &rating, err
}

func CreateRating(rating models.Rating) (int64, error) {
	db, err := helpers.OpenDb()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return 0, err
	}
	defer helpers.CloseDb(db)

	result, err := db.Exec("INSERT INTO rating (song_id, user_id, comment, note ) VALUES (?, ?, ?, ?)",
		rating.SongId, rating.UserId, rating.Comment, rating.Note)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion du rating dans la base de données : %s", err.Error())
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func UpdateRating(ratingID int, rating models.Rating) error {
	db, err := helpers.OpenDb()
	if err != nil {
		logrus.Errorf("Error when opening DB: %s", err.Error())
		return err
	}
	res, err := db.Exec("UPDATE rating SET user_id = ?, song_id= ?, comment = ?, note = ? WHERE id = ?", rating.UserId, rating.SongId, rating.Comment, rating.Note, ratingID)

	rows, err := res.RowsAffected()
	logrus.Printf("%d", rows)
	defer helpers.CloseDb(db)
	if err != nil {
		logrus.Errorf("Repository : Error in updating rating %s", err.Error())
		return err
	}
	return nil
}

func DeleteRating(ratingID int) error {
	db, err := helpers.OpenDb()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDb(db)

	_, err = db.Exec("DELETE FROM rating WHERE id=?", ratingID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du rating dans la base de données : %s", err.Error())
		return err
	}

	return nil
}
