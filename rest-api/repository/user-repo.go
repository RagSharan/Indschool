package repository

import (
	"context"
	"fmt"
	"log"

	"../entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct{}

func NewUserRepository() UserRepo {
	return &repo{}
}

var (
	mongoref       = NewMongoRepo()
	UserCollection = mongoref.ConnectMongoDatabase()
)

type UserRepo interface {
	CreateUser(user entity.User) (entity.User, error)
	FindUser(userId string) (entity.User, error)
	FindUserList() ([]*entity.User, error)
}

func (*repo) CreateUser(user entity.User) (entity.User, error) {
	//obj_id := bson.NewObjectId()
	//user.Id = obj_id
	result, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Inserted a Single Document: ", result.InsertedID)
	//user.ID = json.NewEncoderinsertResult
	//json.NewEncoder().Encode(result)
	return user, err
}

func (*repo) FindUser(userId string) (entity.User, error) {
	filter := bson.D{}
	var user entity.User

	err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func (*repo) FindUserList() ([]*entity.User, error) {
	filter := bson.D{}

	var results []*entity.User
	findOptions := options.Find()
	findOptions.SetLimit(2)
	cur, err := UserCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem entity.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results, err
}

func deleteUser() {
	filter := bson.D{}
	deleteResult, err := UserCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
