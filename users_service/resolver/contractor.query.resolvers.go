package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/model"
)

// ContractorByID is the resolver for the contractorById field.
func (r *queryResolver) ContractorByID(ctx context.Context, id int) (*model.ContractorC, error) {
	return r.service.Group().ContractorByID(ctx, id)
}
