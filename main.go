package main

import (
	"log"

	"github.com/danisper/twitter/bd"
	"github.com/danisper/twitter/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base BD")
		return
	}
	handlers.Handlers()
}
