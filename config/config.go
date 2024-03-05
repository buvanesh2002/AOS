package config

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConfig() *mongo.Client {

	mongourl := viper.GetString("mongourl")
	log.Println("================", mongourl)
	Ctx := context.Background()
	defer Ctx.Done()
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(mongourl))
	if err != nil {
		log.Println("Error connecting to", err)
	} else {
		log.Println("mongo connected successfully")
	}
	return client
}


