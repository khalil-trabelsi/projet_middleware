package ratings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tchipify/ratings/internal/models"
	"tchipify/ratings/internal/services/ratings"

	"github.com/sirupsen/logrus"
)

func InsertRating(w http.ResponseWriter, r *http.Request) {
	var newRating models.Rating
	r.Header.Set("Content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&newRating)

	if err != nil {
		logrus.Errorf("Erreur de décodage JSON")
		http.Error(w, "Erreur de décodage json", http.StatusBadRequest)
	}
	idRating, err := ratings.CreateRating(newRating)
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
	w.WriteHeader(http.StatusCreated)

	uri := fmt.Sprintf("/songs/%d", idRating)

	w.Header().Add("Location", uri)
	json.NewEncoder(w).Encode(map[string]int64{"idSong": idRating})
}
