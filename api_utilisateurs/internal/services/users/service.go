package users

import (
	_ "database/sql"
	_ "errors"
	_ "net/http"
	"tchipify/utilisateurs/internal/models"
	repository "tchipify/utilisateurs/internal/repositories/users"

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
