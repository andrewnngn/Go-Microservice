// Code generated by ent, DO NOT EDIT.

package ent

import (
	"golang-boilerplate/ent/contract"
	"time"
)

// CreateContractInput represents a mutation input for creating contracts.
type CreateContractInput struct {
	VendorID        int
	Status          *contract.Status
	StartDate       *time.Time
	EndDate         *time.Time
	TotalAmount     *int
	RemainingAmount *int
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

// Mutate applies the CreateContractInput on the ContractMutation builder.
func (i *CreateContractInput) Mutate(m *ContractMutation) {
	m.SetVendorID(i.VendorID)
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.StartDate; v != nil {
		m.SetStartDate(*v)
	}
	if v := i.EndDate; v != nil {
		m.SetEndDate(*v)
	}
	if v := i.TotalAmount; v != nil {
		m.SetTotalAmount(*v)
	}
	if v := i.RemainingAmount; v != nil {
		m.SetRemainingAmount(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateContractInput on the ContractCreate builder.
func (c *ContractCreate) SetInput(i CreateContractInput) *ContractCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateContractInput represents a mutation input for updating contracts.
type UpdateContractInput struct {
	VendorID        *int
	Status          *contract.Status
	ClearStartDate  bool
	StartDate       *time.Time
	ClearEndDate    bool
	EndDate         *time.Time
	TotalAmount     *int
	RemainingAmount *int
}

// Mutate applies the UpdateContractInput on the ContractMutation builder.
func (i *UpdateContractInput) Mutate(m *ContractMutation) {
	if v := i.VendorID; v != nil {
		m.SetVendorID(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearStartDate {
		m.ClearStartDate()
	}
	if v := i.StartDate; v != nil {
		m.SetStartDate(*v)
	}
	if i.ClearEndDate {
		m.ClearEndDate()
	}
	if v := i.EndDate; v != nil {
		m.SetEndDate(*v)
	}
	if v := i.TotalAmount; v != nil {
		m.SetTotalAmount(*v)
	}
	if v := i.RemainingAmount; v != nil {
		m.SetRemainingAmount(*v)
	}
}

// SetInput applies the change-set in the UpdateContractInput on the ContractUpdate builder.
func (c *ContractUpdate) SetInput(i UpdateContractInput) *ContractUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateContractInput on the ContractUpdateOne builder.
func (c *ContractUpdateOne) SetInput(i UpdateContractInput) *ContractUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
