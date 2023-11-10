// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golang-boilerplate/ent/contract"
	"golang-boilerplate/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContractUpdate is the builder for updating Contract entities.
type ContractUpdate struct {
	config
	hooks    []Hook
	mutation *ContractMutation
}

// Where appends a list predicates to the ContractUpdate builder.
func (cu *ContractUpdate) Where(ps ...predicate.Contract) *ContractUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetVendorID sets the "vendor_id" field.
func (cu *ContractUpdate) SetVendorID(i int) *ContractUpdate {
	cu.mutation.ResetVendorID()
	cu.mutation.SetVendorID(i)
	return cu
}

// AddVendorID adds i to the "vendor_id" field.
func (cu *ContractUpdate) AddVendorID(i int) *ContractUpdate {
	cu.mutation.AddVendorID(i)
	return cu
}

// SetStatus sets the "status" field.
func (cu *ContractUpdate) SetStatus(c contract.Status) *ContractUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableStatus(c *contract.Status) *ContractUpdate {
	if c != nil {
		cu.SetStatus(*c)
	}
	return cu
}

// SetStartDate sets the "start_date" field.
func (cu *ContractUpdate) SetStartDate(t time.Time) *ContractUpdate {
	cu.mutation.SetStartDate(t)
	return cu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableStartDate(t *time.Time) *ContractUpdate {
	if t != nil {
		cu.SetStartDate(*t)
	}
	return cu
}

// ClearStartDate clears the value of the "start_date" field.
func (cu *ContractUpdate) ClearStartDate() *ContractUpdate {
	cu.mutation.ClearStartDate()
	return cu
}

// SetEndDate sets the "end_date" field.
func (cu *ContractUpdate) SetEndDate(t time.Time) *ContractUpdate {
	cu.mutation.SetEndDate(t)
	return cu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableEndDate(t *time.Time) *ContractUpdate {
	if t != nil {
		cu.SetEndDate(*t)
	}
	return cu
}

// ClearEndDate clears the value of the "end_date" field.
func (cu *ContractUpdate) ClearEndDate() *ContractUpdate {
	cu.mutation.ClearEndDate()
	return cu
}

// SetTotalAmount sets the "total_amount" field.
func (cu *ContractUpdate) SetTotalAmount(i int) *ContractUpdate {
	cu.mutation.ResetTotalAmount()
	cu.mutation.SetTotalAmount(i)
	return cu
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableTotalAmount(i *int) *ContractUpdate {
	if i != nil {
		cu.SetTotalAmount(*i)
	}
	return cu
}

// AddTotalAmount adds i to the "total_amount" field.
func (cu *ContractUpdate) AddTotalAmount(i int) *ContractUpdate {
	cu.mutation.AddTotalAmount(i)
	return cu
}

// SetRemainingAmount sets the "remaining_amount" field.
func (cu *ContractUpdate) SetRemainingAmount(i int) *ContractUpdate {
	cu.mutation.ResetRemainingAmount()
	cu.mutation.SetRemainingAmount(i)
	return cu
}

// SetNillableRemainingAmount sets the "remaining_amount" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableRemainingAmount(i *int) *ContractUpdate {
	if i != nil {
		cu.SetRemainingAmount(*i)
	}
	return cu
}

// AddRemainingAmount adds i to the "remaining_amount" field.
func (cu *ContractUpdate) AddRemainingAmount(i int) *ContractUpdate {
	cu.mutation.AddRemainingAmount(i)
	return cu
}

// Mutation returns the ContractMutation object of the builder.
func (cu *ContractUpdate) Mutation() *ContractMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContractUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContractUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContractUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContractUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ContractUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := contract.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ContractUpdate) check() error {
	if v, ok := cu.mutation.Status(); ok {
		if err := contract.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Contract.status": %w`, err)}
		}
	}
	return nil
}

func (cu *ContractUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contract.Table,
			Columns: contract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contract.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.VendorID(); ok {
		_spec.SetField(contract.FieldVendorID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedVendorID(); ok {
		_spec.AddField(contract.FieldVendorID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(contract.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.StartDate(); ok {
		_spec.SetField(contract.FieldStartDate, field.TypeTime, value)
	}
	if cu.mutation.StartDateCleared() {
		_spec.ClearField(contract.FieldStartDate, field.TypeTime)
	}
	if value, ok := cu.mutation.EndDate(); ok {
		_spec.SetField(contract.FieldEndDate, field.TypeTime, value)
	}
	if cu.mutation.EndDateCleared() {
		_spec.ClearField(contract.FieldEndDate, field.TypeTime)
	}
	if value, ok := cu.mutation.TotalAmount(); ok {
		_spec.SetField(contract.FieldTotalAmount, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedTotalAmount(); ok {
		_spec.AddField(contract.FieldTotalAmount, field.TypeInt, value)
	}
	if value, ok := cu.mutation.RemainingAmount(); ok {
		_spec.SetField(contract.FieldRemainingAmount, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedRemainingAmount(); ok {
		_spec.AddField(contract.FieldRemainingAmount, field.TypeInt, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(contract.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ContractUpdateOne is the builder for updating a single Contract entity.
type ContractUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContractMutation
}

// SetVendorID sets the "vendor_id" field.
func (cuo *ContractUpdateOne) SetVendorID(i int) *ContractUpdateOne {
	cuo.mutation.ResetVendorID()
	cuo.mutation.SetVendorID(i)
	return cuo
}

// AddVendorID adds i to the "vendor_id" field.
func (cuo *ContractUpdateOne) AddVendorID(i int) *ContractUpdateOne {
	cuo.mutation.AddVendorID(i)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *ContractUpdateOne) SetStatus(c contract.Status) *ContractUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableStatus(c *contract.Status) *ContractUpdateOne {
	if c != nil {
		cuo.SetStatus(*c)
	}
	return cuo
}

// SetStartDate sets the "start_date" field.
func (cuo *ContractUpdateOne) SetStartDate(t time.Time) *ContractUpdateOne {
	cuo.mutation.SetStartDate(t)
	return cuo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableStartDate(t *time.Time) *ContractUpdateOne {
	if t != nil {
		cuo.SetStartDate(*t)
	}
	return cuo
}

// ClearStartDate clears the value of the "start_date" field.
func (cuo *ContractUpdateOne) ClearStartDate() *ContractUpdateOne {
	cuo.mutation.ClearStartDate()
	return cuo
}

// SetEndDate sets the "end_date" field.
func (cuo *ContractUpdateOne) SetEndDate(t time.Time) *ContractUpdateOne {
	cuo.mutation.SetEndDate(t)
	return cuo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableEndDate(t *time.Time) *ContractUpdateOne {
	if t != nil {
		cuo.SetEndDate(*t)
	}
	return cuo
}

// ClearEndDate clears the value of the "end_date" field.
func (cuo *ContractUpdateOne) ClearEndDate() *ContractUpdateOne {
	cuo.mutation.ClearEndDate()
	return cuo
}

// SetTotalAmount sets the "total_amount" field.
func (cuo *ContractUpdateOne) SetTotalAmount(i int) *ContractUpdateOne {
	cuo.mutation.ResetTotalAmount()
	cuo.mutation.SetTotalAmount(i)
	return cuo
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableTotalAmount(i *int) *ContractUpdateOne {
	if i != nil {
		cuo.SetTotalAmount(*i)
	}
	return cuo
}

// AddTotalAmount adds i to the "total_amount" field.
func (cuo *ContractUpdateOne) AddTotalAmount(i int) *ContractUpdateOne {
	cuo.mutation.AddTotalAmount(i)
	return cuo
}

// SetRemainingAmount sets the "remaining_amount" field.
func (cuo *ContractUpdateOne) SetRemainingAmount(i int) *ContractUpdateOne {
	cuo.mutation.ResetRemainingAmount()
	cuo.mutation.SetRemainingAmount(i)
	return cuo
}

// SetNillableRemainingAmount sets the "remaining_amount" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableRemainingAmount(i *int) *ContractUpdateOne {
	if i != nil {
		cuo.SetRemainingAmount(*i)
	}
	return cuo
}

// AddRemainingAmount adds i to the "remaining_amount" field.
func (cuo *ContractUpdateOne) AddRemainingAmount(i int) *ContractUpdateOne {
	cuo.mutation.AddRemainingAmount(i)
	return cuo
}

// Mutation returns the ContractMutation object of the builder.
func (cuo *ContractUpdateOne) Mutation() *ContractMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContractUpdateOne) Select(field string, fields ...string) *ContractUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contract entity.
func (cuo *ContractUpdateOne) Save(ctx context.Context) (*Contract, error) {
	var (
		err  error
		node *Contract
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Contract)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ContractMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContractUpdateOne) SaveX(ctx context.Context) *Contract {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContractUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContractUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ContractUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := contract.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ContractUpdateOne) check() error {
	if v, ok := cuo.mutation.Status(); ok {
		if err := contract.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Contract.status": %w`, err)}
		}
	}
	return nil
}

func (cuo *ContractUpdateOne) sqlSave(ctx context.Context) (_node *Contract, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contract.Table,
			Columns: contract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contract.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contract.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contract.FieldID)
		for _, f := range fields {
			if !contract.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.VendorID(); ok {
		_spec.SetField(contract.FieldVendorID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedVendorID(); ok {
		_spec.AddField(contract.FieldVendorID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(contract.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.StartDate(); ok {
		_spec.SetField(contract.FieldStartDate, field.TypeTime, value)
	}
	if cuo.mutation.StartDateCleared() {
		_spec.ClearField(contract.FieldStartDate, field.TypeTime)
	}
	if value, ok := cuo.mutation.EndDate(); ok {
		_spec.SetField(contract.FieldEndDate, field.TypeTime, value)
	}
	if cuo.mutation.EndDateCleared() {
		_spec.ClearField(contract.FieldEndDate, field.TypeTime)
	}
	if value, ok := cuo.mutation.TotalAmount(); ok {
		_spec.SetField(contract.FieldTotalAmount, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedTotalAmount(); ok {
		_spec.AddField(contract.FieldTotalAmount, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.RemainingAmount(); ok {
		_spec.SetField(contract.FieldRemainingAmount, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedRemainingAmount(); ok {
		_spec.AddField(contract.FieldRemainingAmount, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(contract.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Contract{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}