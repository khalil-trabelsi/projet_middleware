package users

import (
	"database/sql"
	_ "database/sql"
	"errors"
	_ "errors"
	"net/http"
	_ "net/http"
	"tchipify/utilisateurs/internal/models"
	repository "tchipify/utilisateurs/internal/repositories/users"

	"github.com/gofrs/uuid"
	_ "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func CreateUser(user models.User) error {
	err := repository.CreateUser(user)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du user : %s", err.Error())
		return err
	}
	return nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving user : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}

func DeleteUser(userID uuid.UUID) error {
	err := repository.DeleteUser(userID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du commentaire : %s", err.Error())
		return err
	}

	return nil
}

func UpdateUser(userID uuid.UUID, updatedUser models.User) error {
	user, err := repository.GetUserById(userID)
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération du user : %s", err.Error())
		return err
	}

	// Mettre à jour les champs nécessaires du commentaire récupéré avec les données du commentaire mis à jour
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	err = repository.UpdateUser(user)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du user en base de données : %s", err.Error())
		return err
	}

	return nil
}
