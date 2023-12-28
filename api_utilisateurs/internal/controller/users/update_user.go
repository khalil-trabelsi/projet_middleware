package users

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"tchipify/utilisateurs/internal/models"
	"tchipify/utilisateurs/internal/services/users"

	"github.com/go-chi/chi/v5"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.FromString(chi.URLParam(r, "id"))

	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du user : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = users.UpdateUser(userID, updatedUser)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du user : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(updatedUser)
	_, _ = w.Write(response)
}
