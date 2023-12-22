package musiques

import (
	"encoding/json"
	"net/http"
	"tchipify/musiques/internal/models"
	_ "tchipify/musiques/internal/models"
	"tchipify/musiques/internal/services/musiques"

	"fmt"

	"github.com/sirupsen/logrus"
)

func UpdateSong(w http.ResponseWriter, r *http.Request) {

	var song models.Song

	ctx := r.Context()
	songId, _ := ctx.Value("songId").(int)

	// Test if the song exist in the Db else Throw an error to the client
	_, err := musiques.GetSongById(songId)
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

	r.Header.Set("Content-type", "application/json")
	error2 := json.NewDecoder(r.Body).Decode(&song)

	if error2 != nil {
		logrus.Errorf("Erreur de décodage JSON")
		http.Error(w, "Erreur de décodage json", http.StatusBadRequest)
	}
	erro := musiques.UpdateSong(songId, song)
	if erro != nil {
		// logging error
		logrus.Errorf("error : %s", erro.Error())
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

	uri := fmt.Sprintf("/songs/%d", song.Id)

	w.Header().Add("Location", uri)
}
