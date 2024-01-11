package ratings

import (
	"encoding/json"
	"net/http"
	"tchipify/ratings/internal/models"
	"tchipify/ratings/internal/services/ratings"

	"github.com/sirupsen/logrus"
)

func GetRatings(w http.ResponseWriter, r *http.Request) {

	// calling service
	ratings, err := ratings.GetAllRatings()
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
	body, _ := json.Marshal(ratings)
	_, _ = w.Write(body)
	return
}
