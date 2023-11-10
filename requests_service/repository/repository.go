package repository

import (
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/logger"
	"golang-boilerplate/repository/request"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	Request() request.Repository
}

// impl is the implementation of the repository registry.
type impl struct {
	request request.Repository
}

// New creates a new repository registry.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit) RepositoryRegistry {
	logg := logger.NewLogger()
	return &impl{
		request: request.New(client, grpc, logg),
	}
}

// Request returns the request repository.
func (i impl) Request() request.Repository {
	return i.request
}
