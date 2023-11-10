package repository

import (
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/logger"
	"golang-boilerplate/repository/auth"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	Auth() auth.Repository
}

// impl is the implementation of the repository registry.
type impl struct {
	auth auth.Repository
}

// New creates a new repository registry.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit) RepositoryRegistry {
	logg := logger.NewLogger()
	return &impl{
		auth: auth.New(client, grpc, logg),
	}
}

// Auth returns the auth repository.
func (i impl) Auth() auth.Repository {
	return i.auth
}
