package musiques
import (
	_"github.com/gofrs/uuid"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"tchipify/musiques/internal/models"
	"tchipify/musiques/internal/services/musiques"
	"net/http"
)
func GetSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songId, _ := ctx.Value("songId").(int)

	song, err := musiques.GetSongById(songId)
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

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}