package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-boilerplate/ent"
	graphql1 "golang-boilerplate/graphql"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Contracts is the resolver for the contracts field.
func (r *queryResolver) Contracts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ContractOrder, where *ent.ContractWhereInput) (*ent.ContractConnection, error) {
	return r.client.Contract.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithContractOrder(orderBy),
			ent.WithContractFilter(where.Filter),
		)
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
