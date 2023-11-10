package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

// CreateRequest is the resolver for the createRequest field.
func (r *mutationResolver) CreateRequest(ctx context.Context, input ent.CreateRequestInput) (*model.RequestReturn, error) {
	return r.service.Request().CreateRequest(ctx, input)
}

// UpdateRequest is the resolver for the updateRequest field.
func (r *mutationResolver) UpdateRequest(ctx context.Context, id int, input ent.UpdateRequestInput) (*model.RequestReturn, error) {
	return r.service.Request().UpdateRequest(ctx, id, input)
}

// DeleteRequest is the resolver for the deleteRequest field.
func (r *mutationResolver) DeleteRequest(ctx context.Context, id int) (*model.RequestReturn, error) {
	return r.service.Request().DeleteRequest(ctx, id)
}

// QueryAllRequests is the resolver for the QueryAllRequests field.
func (r *queryResolver) QueryAllRequests(ctx context.Context) (*model.WithdrawRequestReturn, error) {
	return r.service.Request().QueryAllRequests(ctx)
}
