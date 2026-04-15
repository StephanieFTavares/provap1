package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb+srv://stephaniefortava98_db_user:73wmoAWXeS3DTcis@provap1.50jg4vu.mongodb.net/condlink?retryWrites=true&w=majority"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Erro ao conectar no MongoDB")
	}

	DB = client.Database("condlink")

	log.Println("Mongo conectado 🚀")
}
