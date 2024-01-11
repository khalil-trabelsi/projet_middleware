package ratings

import (
	"net/http"

	"strconv"
	"tchipify/ratings/internal/services/ratings"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func DeleteRating(w http.ResponseWriter, r *http.Request) {
	rating_id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "erreur lors la recuperation de user ID", http.StatusBadRequest)
		return
	}
	err = ratings.DeleteRating(rating_id)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du rating : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
