package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

// GetToken is the resolver for the getToken field.
func (r *mutationResolver) GetToken(ctx context.Context, input model.LoginInput) (*model.AuthUserDetailsReturn, error) {
	return r.service.Auth().GetToken(ctx, input)
}

// ExchangeToken is the resolver for the exchangeToken field.
func (r *mutationResolver) ExchangeToken(ctx context.Context, refreshToken string) (*model.AuthUserDetailsReturn, error) {
	return r.service.Auth().ExchangeToken(ctx, refreshToken)
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context, refreshToken string) (string, error) {
	return r.service.Auth().Logout(ctx, refreshToken)
}

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input ent.CreateAuthInput) (*model.AuthReturn, error) {
	return r.service.Auth().CreateAccount(ctx, input)
}

// UpdateAccount is the resolver for the updateAccount field.
func (r *mutationResolver) UpdateAccount(ctx context.Context, id int, input ent.UpdateAuthInput) (*model.AuthReturn, error) {
	return r.service.Auth().UpdateAccount(ctx, id, input)
}

// DeleteAccount is the resolver for the deleteAccount field.
func (r *mutationResolver) DeleteAccount(ctx context.Context, accountID int) (*model.AuthReturn, error) {
	return r.service.Auth().DeleteAccount(ctx, accountID)
}
