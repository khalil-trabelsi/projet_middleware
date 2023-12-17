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

func AddSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song

	r.Header.Set("Content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&song)

	if err != nil {
		logrus.Errorf("Erreur de décodage JSON")
		http.Error(w, "Erreur de décodage json", http.StatusBadRequest)
	}

	idSong, err := musiques.AddSong(song)
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

	uri := fmt.Sprintf("/songs/%d", idSong)

	w.Header().Add("Location", uri)
	json.NewEncoder(w).Encode(map[string]int64{"idSong": idSong})
}
