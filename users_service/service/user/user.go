package user

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
	"golang-boilerplate/repository"

	"go.uber.org/zap"
)

// Service is the interface for the user service.
type Service interface {
	CreateUser(ctx context.Context, input ent.CreateUserInput) (*model.UserReturn, error)
	UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*model.UserReturn, error)
	DeleteUser(ctx context.Context, id int) (*model.UserReturn, error)
}

// impl is the implementation of the user service.
type impl struct {
	repoRegistry repository.RepositoryRegistry
	logger       *zap.Logger
}

// New creates a new user service.
func New(repoRegistry repository.RepositoryRegistry, logger *zap.Logger) Service {
	return &impl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (i impl) CreateUser(ctx context.Context, input ent.CreateUserInput) (*model.UserReturn, error) {
	return i.repoRegistry.User().CreateUser(ctx, input)
}

func (i impl) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*model.UserReturn, error) {
	return i.repoRegistry.User().UpdateUser(ctx, id, input)
}

func (i impl) DeleteUser(ctx context.Context, id int) (*model.UserReturn, error) {
	return i.repoRegistry.User().DeleteUser(ctx, id)
}
