package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

// CreateContract is the resolver for the createContract field.
func (r *mutationResolver) CreateContract(ctx context.Context, input ent.CreateContractInput) (*model.ContractReturn, error) {
	return r.service.Contract().CreateContract(ctx, input)
}

// UpdateContract is the resolver for the updateContract field.
func (r *mutationResolver) UpdateContract(ctx context.Context, id int, input ent.UpdateContractInput) (*model.ContractReturn, error) {
	return r.service.Contract().UpdateContract(ctx, id, input)
}

// DeleteContract is the resolver for the deleteContract field.
func (r *mutationResolver) DeleteContract(ctx context.Context, id int) (*model.ContractReturn, error) {
	return r.service.Contract().DeleteContract(ctx, id)
}

// QueryAllContracts is the resolver for the QueryAllContracts field.
func (r *queryResolver) QueryAllContracts(ctx context.Context) (*model.ContractVendorReturn, error) {
	return r.service.Contract().QueryAllContracts(ctx)
}
