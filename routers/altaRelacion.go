package routers

import (
	"net/http"

	"github.com/danisper/twitter/bd"
	"github.com/danisper/twitter/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro id ES OBLIGATORIO", http.StatusBadRequest)
		return
	}

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio error al insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se  ha logrado insertar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
