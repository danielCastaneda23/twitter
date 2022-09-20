package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danisper/twitter/bd"
	"github.com/danisper/twitter/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweets(registro)
	if err != nil {
		http.Error(w, "Ocurrio un Error al intentar insertar el registro"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
