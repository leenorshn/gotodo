package database

import (
	"context"
	"errors"
	"time"

	"github.com/leenorshn/gotodo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) Withdraw(from string, amount float32) (*models.Trans, error) {
	FromId, err := primitive.ObjectIDFromHex(from)
	if err != nil {
		panic(err)
	}
	transCollection := colHelper(db, "trans")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	account := db.FindAccount(from)
	if account.Balance < float64(amount) {
		defer cancel()

		return nil, errors.New("message Balance insuffisant")

	}
	defer cancel()
	newTrans := bson.D{
		{Key: "from", Value: FromId},
		{Key: "amount", Value: amount},
		{Key: "operation", Value: "withdraw"},
		{Key: "date", Value: time.Now()},
	}
	res, err := transCollection.InsertOne(ctx, newTrans)
	if err != nil {
		panic("Erreur d'insertion de donnee")
	}
	return &models.Trans{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		From:      db.FindAccount(from),
		Amount:    float64(amount),
		Operation: "withdraw",
		Date:      time.Now(),
	}, nil
}
