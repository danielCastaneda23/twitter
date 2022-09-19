package bd

import (
	"github.com/danisper/twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := CheckUserExist(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true

}
