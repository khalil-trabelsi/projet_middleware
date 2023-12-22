package ratings

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"tchipify/ratings/internal/services/ratings"
)

func DeleteRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ratingId, _ := ctx.Value("ratingId").(int)
	logrus.Printf("%d", ratingId)
	err := ratings.DeleteRating(ratingId)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du rating : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
