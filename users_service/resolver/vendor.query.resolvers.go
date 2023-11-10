package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/model"
)

// VendorByID is the resolver for the vendorById field.
func (r *queryResolver) VendorByID(ctx context.Context, id int) (*model.VendorReturn, error) {
	return r.service.Group().VendorByID(ctx, id)
}
