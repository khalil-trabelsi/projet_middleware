package musiques

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"tchipify/musiques/internal/models"
	"net/http"
	"strconv"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		songId, err := strconv.Atoi(idParam)
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

		ctx := context.WithValue(r.Context(), "songId", songId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
