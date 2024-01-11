package ratings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tchipify/ratings/internal/models"
	"tchipify/ratings/internal/services/ratings"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating

	rating_id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "erreur lors la recuperation de user ID", http.StatusBadRequest)
		return
	}
	r.Header.Set("Content-type", "application/json")
	error2 := json.NewDecoder(r.Body).Decode(&rating)

	if error2 != nil {
		logrus.Errorf("Erreur de décodage JSON")
		http.Error(w, "Erreur de décodage json", http.StatusBadRequest)
	}
	erro := ratings.UpdateRating(rating_id, rating)
	if erro != nil {
		logrus.Errorf("error in controller update rating : %s", erro.Error())
		customError, isCustom := erro.(*models.CustomError)
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

	uri := fmt.Sprintf("/songs/%d", rating.Id)

	w.Header().Add("Location", uri)
}
