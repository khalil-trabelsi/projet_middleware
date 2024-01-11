package musiques

import (
	"encoding/json"
	"net/http"
	"tchipify/musiques/internal/models"
	_ "tchipify/musiques/internal/models"
	"tchipify/musiques/internal/services/musiques"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func AddSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song

	r.Header.Set("Content-type", "application/json")
	print(r.Body)
	err := json.NewDecoder(r.Body).Decode(&song)

	if err != nil {
		logrus.Errorf("Erreur de décodage JSON")
		http.Error(w, "Erreur de décodage json", http.StatusBadRequest)
		return
	}
	id, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("Erreur lors de la génération de l'identifiant UUID : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	song.Id = &id

	err = musiques.AddSong(song)
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
	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
}
