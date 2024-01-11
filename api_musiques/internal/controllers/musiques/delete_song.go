package musiques

import (
	"encoding/json"
	"net/http"
	"tchipify/musiques/internal/models"
	"tchipify/musiques/internal/services/musiques"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func DeleteSong(w http.ResponseWriter, r *http.Request) {
	songId, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "erreur lors la recuperation de user ID", http.StatusBadRequest)
		return
	}

	err = musiques.DeleteSong(songId)
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

	w.WriteHeader(http.StatusNoContent)
	return
}
