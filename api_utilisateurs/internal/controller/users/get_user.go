package users

import (
	"encoding/json"

	"net/http"
	"tchipify/utilisateurs/internal/models"
	"tchipify/utilisateurs/internal/services/users"

	"github.com/gofrs/uuid"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	userId, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "erreur lors la recuperation de user ID", http.StatusBadRequest)
		return
	}

	user, err := users.GetUserById(userId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}
