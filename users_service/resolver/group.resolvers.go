package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input ent.CreateGroupInput) (*model.GroupReturn, error) {
	return r.service.Group().CreateGroup(ctx, input)
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, id int, input ent.UpdateGroupInput) (*model.GroupReturn, error) {
	return r.service.Group().UpdateGroup(ctx, id, input)
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id int) (*model.GroupReturn, error) {
	return r.service.Group().DeleteGroup(ctx, id)
}
