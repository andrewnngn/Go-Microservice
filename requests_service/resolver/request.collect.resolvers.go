package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/model"
	"time"
)

// ReadyToCollect is the resolver for the readyToCollect field.
func (r *mutationResolver) ReadyToCollect(ctx context.Context, id int) (*model.WithdrawCollectReturn, error) {
	return r.service.Request().ReadyToCollect(ctx, id)
}

// ArrangeCollectionDate is the resolver for the arrangeCollectionDate field.
func (r *mutationResolver) ArrangeCollectionDate(ctx context.Context, id int, collectionDate time.Time) (string, error) {
	return r.service.Request().ArrangeCollectionDate(ctx, id, collectionDate), nil
}

// Collected is the resolver for the collected field.
func (r *mutationResolver) Collected(ctx context.Context, id int) (*model.WithdrawCollectReturn, error) {
	return r.service.Request().Collected(ctx, id)
}
