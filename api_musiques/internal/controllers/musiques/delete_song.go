package musiques

import (
	"encoding/json"
	"net/http"
	"tchipify/musiques/internal/models"
	"tchipify/musiques/internal/services/musiques"

	_ "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func DeleteSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(int)

	err := musiques.DeleteSong(songId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
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

	w.WriteHeader(http.StatusNoContent)
	return
}
