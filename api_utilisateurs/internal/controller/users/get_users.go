package users

import (
	"encoding/json"
	"net/http"
	"tchipify/utilisateurs/internal/models"
	"tchipify/utilisateurs/internal/services/users"

	"github.com/sirupsen/logrus"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	// calling service
	users, err := users.GetAllUsers()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(users)
	_, _ = w.Write(body)
	return
}
