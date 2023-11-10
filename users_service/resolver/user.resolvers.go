package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*model.UserReturn, error) {
	return r.service.User().CreateUser(ctx, input)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*model.UserReturn, error) {
	return r.service.User().UpdateUser(ctx, id, input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.UserReturn, error) {
	return r.service.User().DeleteUser(ctx, id)
}
