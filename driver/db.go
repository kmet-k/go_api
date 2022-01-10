package driver

import (
	"context"
	"os"

	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	gotenv.Load()
}

func ConnectMongo(collection_name string) (*mongo.Collection, *mongo.Client, error) {

	auth := options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("DB")).SetAuth(auth)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("test").Collection(collection_name)

	return collection, client, nil
}
