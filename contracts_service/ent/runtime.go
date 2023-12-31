// Code generated by ent, DO NOT EDIT.

package ent

import (
	"golang-boilerplate/ent/contract"
	"golang-boilerplate/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contractFields := schema.Contract{}.Fields()
	_ = contractFields
	// contractDescTotalAmount is the schema descriptor for total_amount field.
	contractDescTotalAmount := contractFields[4].Descriptor()
	// contract.DefaultTotalAmount holds the default value on creation for the total_amount field.
	contract.DefaultTotalAmount = contractDescTotalAmount.Default.(int)
	// contractDescRemainingAmount is the schema descriptor for remaining_amount field.
	contractDescRemainingAmount := contractFields[5].Descriptor()
	// contract.DefaultRemainingAmount holds the default value on creation for the remaining_amount field.
	contract.DefaultRemainingAmount = contractDescRemainingAmount.Default.(int)
	// contractDescCreatedAt is the schema descriptor for created_at field.
	contractDescCreatedAt := contractFields[6].Descriptor()
	// contract.DefaultCreatedAt holds the default value on creation for the created_at field.
	contract.DefaultCreatedAt = contractDescCreatedAt.Default.(func() time.Time)
	// contractDescUpdatedAt is the schema descriptor for updated_at field.
	contractDescUpdatedAt := contractFields[7].Descriptor()
	// contract.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	contract.DefaultUpdatedAt = contractDescUpdatedAt.Default.(func() time.Time)
	// contract.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	contract.UpdateDefaultUpdatedAt = contractDescUpdatedAt.UpdateDefault.(func() time.Time)
}
