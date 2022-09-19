package routers

import (
	"encoding/json"
	"net/http"

	"github.com/danisper/twitter/bd"
	"github.com/danisper/twitter/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los  datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.CheckUserExist(t.Email)
	if encontrado {
		http.Error(w, "El usuario ya se encuentra registrado", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
