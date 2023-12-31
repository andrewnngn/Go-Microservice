// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateAuthInput represents a mutation input for creating auths.
type CreateAuthInput struct {
	Username     string
	Password     string
	RefreshToken *string
	UserID       int
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

// Mutate applies the CreateAuthInput on the AuthMutation builder.
func (i *CreateAuthInput) Mutate(m *AuthMutation) {
	m.SetUsername(i.Username)
	m.SetPassword(i.Password)
	if v := i.RefreshToken; v != nil {
		m.SetRefreshToken(*v)
	}
	m.SetUserID(i.UserID)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateAuthInput on the AuthCreate builder.
func (c *AuthCreate) SetInput(i CreateAuthInput) *AuthCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAuthInput represents a mutation input for updating auths.
type UpdateAuthInput struct {
	Username          *string
	Password          *string
	ClearRefreshToken bool
	RefreshToken      *string
	UserID            *int
}

// Mutate applies the UpdateAuthInput on the AuthMutation builder.
func (i *UpdateAuthInput) Mutate(m *AuthMutation) {
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if i.ClearRefreshToken {
		m.ClearRefreshToken()
	}
	if v := i.RefreshToken; v != nil {
		m.SetRefreshToken(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateAuthInput on the AuthUpdate builder.
func (c *AuthUpdate) SetInput(i UpdateAuthInput) *AuthUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAuthInput on the AuthUpdateOne builder.
func (c *AuthUpdateOne) SetInput(i UpdateAuthInput) *AuthUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
