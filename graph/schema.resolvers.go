package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/leenorshn/gotodo/database"
	"github.com/leenorshn/gotodo/graph/generated"
	"github.com/leenorshn/gotodo/graph/model"
	"github.com/leenorshn/gotodo/models"
)

// Withdraw is the resolver for the withdraw field.
func (r *mutationResolver) Withdraw(ctx context.Context, from string, amount float64) (*models.Trans, error) {
	trans, err := db.Withdraw(from, float32(amount))
	if err != nil {
		log.Fatalln("Erreur survenu", err)
	}

	return trans, nil
}

// PayRent is the resolver for the payRent field.
func (r *mutationResolver) PayRent(ctx context.Context, from string, amount float64) (*models.Trans, error) {
	panic(fmt.Errorf("not implemented: PayRent - payRent"))
}

// AchatCarburant is the resolver for the achatCarburant field.
func (r *mutationResolver) AchatCarburant(ctx context.Context, from string, amount float64, quantity float64) (*models.Trans, error) {
	panic(fmt.Errorf("not implemented: AchatCarburant - achatCarburant"))
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, phone string, password string) (*models.Auth, error) {
	panic(fmt.Errorf("not implemented: LoginUser - loginUser"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, data model.NewUser) (*models.Auth, error) {
	user := db.InsertUser(data)
	//generate user token for access with jwt
	var token = "token"
	return &models.Auth{
		Token: token,
		User:  user,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, data model.UpdateUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	isDeleted := db.DeleteUser(id)

	return isDeleted, nil
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, phone *string, newPassword string) (bool, error) {
	panic(fmt.Errorf("not implemented: ChangePassword - changePassword"))
}

// Motars is the resolver for the motars field.
func (r *queryResolver) Motars(ctx context.Context) ([]*models.Motar, error) {
	panic(fmt.Errorf("not implemented: Motars - motars"))
}

// Trans is the resolver for the trans field.
func (r *queryResolver) Trans(ctx context.Context) ([]*models.Trans, error) {
	panic(fmt.Errorf("not implemented: Trans - trans"))
}

// GetTrans is the resolver for the getTrans field.
func (r *queryResolver) GetTrans(ctx context.Context, id string) (*models.Trans, error) {
	panic(fmt.Errorf("not implemented: GetTrans - getTrans"))
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context) ([]*models.Account, error) {
	accounts := db.GetAccounts()

	return accounts, nil
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*models.Account, error) {
	account := db.FindAccount(id)

	return account, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	users := db.FindUsers()
	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return db.FindUser(id), nil
}

// Date is the resolver for the date field.
func (r *transResolver) Date(ctx context.Context, obj *models.Trans) (string, error) {
	panic(fmt.Errorf("not implemented: Date - date"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Trans returns generated.TransResolver implementation.
func (r *Resolver) Trans() generated.TransResolver { return &transResolver{r} }

//func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type transResolver struct{ *Resolver }

//type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	db = database.ConnectDB()
)
