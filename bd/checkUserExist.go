package bd

import (
	"context"
	"time"

	"github.com/danisper/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExist(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}
	var resultados models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultados)
	ID := resultados.ID.Hex()
	if err != nil {
		return resultados, false, ID
	}
	return resultados, true, ID

}
