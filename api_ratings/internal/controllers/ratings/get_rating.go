package ratings

import (
	"encoding/json"
	"strconv"

	"net/http"
	"tchipify/ratings/internal/models"
	"tchipify/ratings/internal/services/ratings"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func GetRating(w http.ResponseWriter, r *http.Request) {
	rating_id, err := strconv.Atoi(chi.URLParam(r, "id"))
	rating, err := ratings.GetRatingById(rating_id)
	if err != nil {
		logrus.Errorf("error in controller : %s", err.Error())
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
	body, _ := json.Marshal(rating)
	_, _ = w.Write(body)
	return
}
