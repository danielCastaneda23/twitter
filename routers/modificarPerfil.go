package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danisper/twitter/bd"
	"github.com/danisper/twitter/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	fmt.Println(t)
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha podido modificar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
