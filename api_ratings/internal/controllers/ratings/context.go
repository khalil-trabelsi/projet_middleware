package ratings

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tchipify/ratings/internal/models"

	"github.com/go-chi/chi/v5"
	_ "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		ratingId, err := strconv.Atoi(idParam)
		if err != nil {
			logrus.Errorf("parsing error : %s", err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse id (%s) as Int", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "ratingId", ratingId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
