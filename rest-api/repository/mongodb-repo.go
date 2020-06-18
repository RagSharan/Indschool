/*  install mongodb driver
go get github.com/mongodb/mongo-go-driver
*/

package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongorepo struct{}

type MongoDbRepo interface {
	ConnectMongoDatabase() *mongo.Collection
}

func NewMongoRepo() MongoDbRepo {
	return &mongorepo{}
}

var (
	client       *mongo.Client
	url          string = "mongodb://localhost:27017"
	databaseName string = "UserDatabase"
	//collection   *mongo.Collection
)

func (*mongorepo) ConnectMongoDatabase() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	UserCollection := client.Database("mydb").Collection("users")
	return UserCollection

}
