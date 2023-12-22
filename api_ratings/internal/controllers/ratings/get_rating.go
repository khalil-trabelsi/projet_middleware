package ratings

import (
	"encoding/json"

	"net/http"
	"tchipify/ratings/internal/models"
	"tchipify/ratings/internal/services/ratings"

	_ "github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func GetRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ratingId, _ := ctx.Value("songId").(int)
	rating, err := ratings.GetRatingById(ratingId)
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
