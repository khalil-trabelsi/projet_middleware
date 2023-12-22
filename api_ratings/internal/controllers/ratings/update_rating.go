package ratings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tchipify/ratings/internal/models"
	"tchipify/ratings/internal/services/ratings"

	"github.com/sirupsen/logrus"
)

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating

	ctx := r.Context()
	ratingId, _ := ctx.Value("ratingId").(int)

	r.Header.Set("Content-type", "application/json")
	error2 := json.NewDecoder(r.Body).Decode(&rating)

	if error2 != nil {
		logrus.Errorf("Erreur de décodage JSON")
		http.Error(w, "Erreur de décodage json", http.StatusBadRequest)
	}
	erro := ratings.UpdateRating(ratingId, rating)
	if erro != nil {
		// logging error
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
