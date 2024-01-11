package musiques

import (
	"fmt"
	"tchipify/ratings/internal/helpers"
	"tchipify/ratings/internal/models"

	"github.com/sirupsen/logrus"
)

func GetAllRatings() ([]models.Rating, error) {
	db, err := helpers.OpenDb()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings ")
	helpers.CloseDb(db)
	if err != nil {
		return nil, err
	}

	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id, &data.UserId, &data.SongId, &data.Content, &data.Date, &data.Rating)
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
	fmt.Println(id)
	row := db.QueryRow("SELECT * FROM ratings WHERE id=?  ", id)
	helpers.CloseDb(db)

	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.UserId, &rating.SongId, &rating.Content, &rating.Date, &rating.Rating)
	logrus.Println(err)

	if err != nil {

		return nil, err // Autres erreurs lors du scan
	}
	return &rating, err
}

func CreateRating(rating models.Rating) (int64, error) {
	fmt.Println(rating)
	db, err := helpers.OpenDb()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return 0, err
	}
	defer helpers.CloseDb(db)

	result, err := db.Exec("INSERT INTO ratings (user_id, song_id, content, rating, date ) VALUES (?, ?, ?, ?,?) ",
		rating.UserId, rating.SongId, rating.Content, rating.Rating, rating.Date)
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
	res, err := db.Exec("UPDATE ratings SET content = ?, rating = ? WHERE   id = ?", rating.Content, rating.Rating, ratingID)

	logrus.Printf("%d", res)
	helpers.CloseDb(db)
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

	_, err = db.Exec("DELETE FROM ratings WHERE  id=?", ratingID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du rating dans la base de données : %s", err.Error())
		return err
	}

	return nil
}
