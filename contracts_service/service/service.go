package service

import (
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/repository"
	"golang-boilerplate/service/contract"
)

// ServiceRegistry is the interface for the service registry.
type ServiceRegistry interface {
	Contract() contract.Service
}

// impl is the implementation of the service registry.
type impl struct {
	contract contract.Service
}

// New creates a new service registry.
func New(entClient *ent.Client, logger *zap.Logger, grpc *grpcClient.GRPCClientInit) ServiceRegistry {
	repoRegistry := repository.New(entClient, grpc)
	return &impl{
		contract: contract.New(repoRegistry, logger),
	}
}

// Contract returns the contract service.
func (i impl) Contract() contract.Service {
	return i.contract
}
