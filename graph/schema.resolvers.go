package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/leenorshn/gotodo/database"
	"github.com/leenorshn/gotodo/graph/generated"
	"github.com/leenorshn/gotodo/graph/model"
)

// Withdraw is the resolver for the withdraw field.
func (r *mutationResolver) Withdraw(ctx context.Context, from string, amount float64) (*model.Trans, error) {
	panic(fmt.Errorf("not implemented: Withdraw - withdraw"))
}

// PayRent is the resolver for the payRent field.
func (r *mutationResolver) PayRent(ctx context.Context, from string, amount float64) (*model.Trans, error) {
	panic(fmt.Errorf("not implemented: PayRent - payRent"))
}

// AchatCarburant is the resolver for the achatCarburant field.
func (r *mutationResolver) AchatCarburant(ctx context.Context, from string, amount float64, quantity float64) (*model.Trans, error) {
	panic(fmt.Errorf("not implemented: AchatCarburant - achatCarburant"))
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, phone string, password string) (*model.Auth, error) {
	panic(fmt.Errorf("not implemented: LoginUser - loginUser"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, data model.NewUser) (*model.Auth, error) {
	user := db.InsertUser(data)
	//generate user token for access with jwt
	var token = "token"
	return &model.Auth{
		Token: token,
		User:  user,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, data model.UpdateUser) (*model.User, error) {
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
func (r *queryResolver) Motars(ctx context.Context) ([]*model.Motar, error) {
	panic(fmt.Errorf("not implemented: Motars - motars"))
}

// Trans is the resolver for the trans field.
func (r *queryResolver) Trans(ctx context.Context) ([]*model.Trans, error) {
	panic(fmt.Errorf("not implemented: Trans - trans"))
}

// GetTrans is the resolver for the getTrans field.
func (r *queryResolver) GetTrans(ctx context.Context, id string) (*model.Trans, error) {
	panic(fmt.Errorf("not implemented: GetTrans - getTrans"))
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context) ([]*model.Account, error) {
	panic(fmt.Errorf("not implemented: Accounts - accounts"))
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*model.Account, error) {
	panic(fmt.Errorf("not implemented: Account - account"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := db.FindUsers()
	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return db.FindUser(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

var (
	db = database.ConnectDB()
)
