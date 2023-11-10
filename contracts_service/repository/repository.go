package repository

import (
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/logger"
	"golang-boilerplate/repository/contract"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	Contract() contract.Repository
}

// impl is the implementation of the repository registry.
type impl struct {
	contract contract.Repository
}

// New creates a new repository registry.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit) RepositoryRegistry {
	logg := logger.NewLogger()
	return &impl{
		contract: contract.New(client, grpc, logg),
	}
}

// Contract returns the contract repository.
func (i impl) Contract() contract.Repository {
	return i.contract
}
