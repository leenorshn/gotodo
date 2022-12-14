package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/leenorshn/gotodo/graph/model"
	"github.com/leenorshn/gotodo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) InsertUser(newUser model.NewUser) *models.User {
	userCollection := colHelper(db, "users")
	accountCollection := colHelper(db, "accounts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		panic("Erreur d'insertion de donnee")
	}
	resAccount, errAc := accountCollection.InsertOne(ctx, bson.D{
		{Key: "user", Value: res.InsertedID},
	})
	if errAc != nil {
		log.Fatalln("Erreur de creation de compte", errAc)
	}

	fmt.Println(resAccount.InsertedID)

	return &models.User{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Name:  newUser.Name,
		Phone: newUser.Phone,
	}
}

func (db *DB) DeleteUser(id string) bool {
	userCollection := colHelper(db, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := userCollection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return res.DeletedCount > 0

}

func (db *DB) FindUsers() []*models.User {

	userCollection := colHelper(db, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{})
	cur, err := userCollection.Find(ctx, bson.D{}, opts)
	defer cur.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var users []*models.User

	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	for cur.Next(ctx) {

		var user *models.User
		if err != nil {
			log.Fatal(err)
		}

		if err = cur.Decode(&user); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func (db *DB) FindUser(id string) *models.User {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	userCollection := colHelper(db, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user *models.User

	err = userCollection.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&user)
	if err != nil {
		log.Fatal("user not found ", err)
	}

	return user

}
