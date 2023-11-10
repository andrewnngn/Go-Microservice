// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golang-boilerplate/ent/request"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RequestCreate is the builder for creating a Request entity.
type RequestCreate struct {
	config
	mutation *RequestMutation
	hooks    []Hook
}

// SetContractID sets the "contract_id" field.
func (rc *RequestCreate) SetContractID(i int) *RequestCreate {
	rc.mutation.SetContractID(i)
	return rc
}

// SetVendorID sets the "vendor_id" field.
func (rc *RequestCreate) SetVendorID(i int) *RequestCreate {
	rc.mutation.SetVendorID(i)
	return rc
}

// SetContractorID sets the "contractor_id" field.
func (rc *RequestCreate) SetContractorID(i int) *RequestCreate {
	rc.mutation.SetContractorID(i)
	return rc
}

// SetAmount sets the "amount" field.
func (rc *RequestCreate) SetAmount(i int) *RequestCreate {
	rc.mutation.SetAmount(i)
	return rc
}

// SetStatus sets the "status" field.
func (rc *RequestCreate) SetStatus(r request.Status) *RequestCreate {
	rc.mutation.SetStatus(r)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RequestCreate) SetNillableStatus(r *request.Status) *RequestCreate {
	if r != nil {
		rc.SetStatus(*r)
	}
	return rc
}

// SetCollectionDate sets the "collection_date" field.
func (rc *RequestCreate) SetCollectionDate(t time.Time) *RequestCreate {
	rc.mutation.SetCollectionDate(t)
	return rc
}

// SetNillableCollectionDate sets the "collection_date" field if the given value is not nil.
func (rc *RequestCreate) SetNillableCollectionDate(t *time.Time) *RequestCreate {
	if t != nil {
		rc.SetCollectionDate(*t)
	}
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RequestCreate) SetCreatedAt(t time.Time) *RequestCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RequestCreate) SetNillableCreatedAt(t *time.Time) *RequestCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RequestCreate) SetUpdatedAt(t time.Time) *RequestCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RequestCreate) SetNillableUpdatedAt(t *time.Time) *RequestCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// Mutation returns the RequestMutation object of the builder.
func (rc *RequestCreate) Mutation() *RequestMutation {
	return rc.mutation
}

// Save creates the Request in the database.
func (rc *RequestCreate) Save(ctx context.Context) (*Request, error) {
	var (
		err  error
		node *Request
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Request)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RequestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RequestCreate) SaveX(ctx context.Context) *Request {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RequestCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RequestCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RequestCreate) defaults() {
	if _, ok := rc.mutation.Status(); !ok {
		v := request.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := request.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := request.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RequestCreate) check() error {
	if _, ok := rc.mutation.ContractID(); !ok {
		return &ValidationError{Name: "contract_id", err: errors.New(`ent: missing required field "Request.contract_id"`)}
	}
	if _, ok := rc.mutation.VendorID(); !ok {
		return &ValidationError{Name: "vendor_id", err: errors.New(`ent: missing required field "Request.vendor_id"`)}
	}
	if _, ok := rc.mutation.ContractorID(); !ok {
		return &ValidationError{Name: "contractor_id", err: errors.New(`ent: missing required field "Request.contractor_id"`)}
	}
	if _, ok := rc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Request.amount"`)}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Request.status"`)}
	}
	if v, ok := rc.mutation.Status(); ok {
		if err := request.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Request.status": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Request.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Request.updated_at"`)}
	}
	return nil
}

func (rc *RequestCreate) sqlSave(ctx context.Context) (*Request, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RequestCreate) createSpec() (*Request, *sqlgraph.CreateSpec) {
	var (
		_node = &Request{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: request.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: request.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.ContractID(); ok {
		_spec.SetField(request.FieldContractID, field.TypeInt, value)
		_node.ContractID = value
	}
	if value, ok := rc.mutation.VendorID(); ok {
		_spec.SetField(request.FieldVendorID, field.TypeInt, value)
		_node.VendorID = value
	}
	if value, ok := rc.mutation.ContractorID(); ok {
		_spec.SetField(request.FieldContractorID, field.TypeInt, value)
		_node.ContractorID = value
	}
	if value, ok := rc.mutation.Amount(); ok {
		_spec.SetField(request.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(request.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.CollectionDate(); ok {
		_spec.SetField(request.FieldCollectionDate, field.TypeTime, value)
		_node.CollectionDate = &value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(request.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(request.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// RequestCreateBulk is the builder for creating many Request entities in bulk.
type RequestCreateBulk struct {
	config
	builders []*RequestCreate
}

// Save creates the Request entities in the database.
func (rcb *RequestCreateBulk) Save(ctx context.Context) ([]*Request, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Request, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RequestCreateBulk) SaveX(ctx context.Context) []*Request {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RequestCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RequestCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}