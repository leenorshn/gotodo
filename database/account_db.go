package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/leenorshn/gotodo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) GetAccounts() []*models.Account {
	accountCollection := colHelper(db, "accounts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	qry := []bson.M{

		{
			"$lookup": bson.M{
				"from":         "users", // Child collection to join
				"localField":   "user",  // Parent collection reference holding field
				"foreignField": "_id",   // Child collection reference field
				"as":           "user",  // Arbitrary field name to store result set
			},
		},
		{
			"$unwind": "$user",
		},
	}

	var accounts []*models.Account

	cur, err := accountCollection.Aggregate(ctx, qry)

	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	for cur.Next(ctx) {

		var account *models.Account
		if err != nil {
			log.Fatal(err)
		}

		if err = cur.Decode(&account); err != nil {
			log.Fatal(err)
		}
		fmt.Println(account)

		accounts = append(accounts, account)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return accounts
}

func (db *DB) FindAccount(id string) *models.Account {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	accountCollection := colHelper(db, "accounts")
	fmt.Println(ObjectID)
	qry := []bson.M{
		{"$match": bson.M{"_id": ObjectID}},

		{
			"$lookup": bson.M{
				"from":         "users", // Child collection to join
				"localField":   "user",  // Parent collection reference holding field
				"foreignField": "_id",   // Child collection reference field
				"as":           "user",  // Arbitrary field name to store result set
			},
		},
		{
			"$unwind": "$user",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var account *models.Account

	cursor, err := accountCollection.Aggregate(ctx, qry)
	if err != nil {
		log.Fatal("accounts ", err)
	}

	for cursor.Next(ctx) {

		//fmt.Println(cursor.Current)

		if err != nil {
			log.Fatal(err)
		}

		if err = cursor.Decode(&account); err != nil {
			log.Fatal(err)
		}
		fmt.Println(account)

	}

	return account
}
