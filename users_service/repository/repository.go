package repository

import (
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/logger"
	"golang-boilerplate/repository/group"
	"golang-boilerplate/repository/user"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	User() user.Repository
	Group() group.Repository
}

// impl is the implementation of the repository registry.
type impl struct {
	user  user.Repository
	group group.Repository
}

// New creates a new repository registry.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit) RepositoryRegistry {
	logg := logger.NewLogger()
	return &impl{
		user:  user.New(client, grpc, logg),
		group: group.New(client, grpc, logg),
	}
}

// User returns the user repository.
func (i impl) User() user.Repository {
	return i.user
}

// User returns the user repository.
func (i impl) Group() group.Repository {
	return i.group
}
