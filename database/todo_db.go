// package mongo

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/leenorshn/gotodo/graph/model"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func (db *DB) InsertTodo(newTodo model.NewTodo) *model.Todo {
// 	startTime := time.Now()
// 	fmt.Println(startTime)
// 	ObjectID, err := primitive.ObjectIDFromHex(newTodo.User)
// 	if err != nil {
// 		log.Fatal("user malformated")
// 	}

// 	todoCollection := colHelper(db, "todos")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	result, err := todoCollection.InsertOne(ctx, bson.D{
// 		{Key: "text", Value: newTodo.Text},
// 		{Key: "user", Value: ObjectID},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	todo := &model.Todo{
// 		ID:   result.InsertedID.(primitive.ObjectID).Hex(),
// 		Text: newTodo.Text,
// 		Done: false,
// 		User: db.FindUser(newTodo.User),
// 	}
// 	endTime := time.Now()

// 	fmt.Println(endTime)
// 	fmt.Println(endTime.Sub(startTime))
// 	return todo

// }

// func (db *DB) FindAllTodos() []*model.Todo {
// 	startTime := time.Now()
// 	fmt.Println(startTime)

// 	userCollection := colHelper(db, "todos")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	qry := []bson.M{

// 		{
// 			"$lookup": bson.M{
// 				"from":         "users", // Child collection to join
// 				"localField":   "user",  // Parent collection reference holding field
// 				"foreignField": "_id",   // Child collection reference field
// 				"as":           "user",  // Arbitrary field name to store result set
// 			},
// 		},
// 		{
// 			"$unwind": "$user",
// 		},
// 	}

// 	var todos []*model.Todo

// 	cur, err := userCollection.Aggregate(ctx, qry)

// 	defer cur.Close(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for cur.Next(ctx) {

// 		var todo *model.Todo
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		//el, _ := cur.Current.Elements()
// 		//fmt.Println(el)
// 		if err = cur.Decode(&todo); err != nil {
// 			log.Fatal(err)
// 		}

// 		todos = append(todos, todo)
// 	}

// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	endTime := time.Now()

// 	fmt.Println(endTime)
// 	fmt.Println(endTime.Sub(startTime))

// 	return todos
// }
// func (db *DB) FindTodo(id string) *model.Todo {
// 	ObjectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	userCollection := colHelper(db, "todos")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	var todo *model.Todo

// 	err = userCollection.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&todo)
// 	if err != nil {
// 		panic("todo not found")
// 	}

// 	return todo

// }

package database
