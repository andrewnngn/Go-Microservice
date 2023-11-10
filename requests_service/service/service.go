package service

import (
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/repository"
	"golang-boilerplate/service/request"
)

// ServiceRegistry is the interface for the service registry.
type ServiceRegistry interface {
	Request() request.Service
}

// impl is the implementation of the service registry.
type impl struct {
	request request.Service
}

// New creates a new service registry.
func New(entClient *ent.Client, logger *zap.Logger, grpc *grpcClient.GRPCClientInit) ServiceRegistry {
	repoRegistry := repository.New(entClient, grpc)
	return &impl{
		request: request.New(repoRegistry, logger),
	}
}

// Request returns the request service.
func (i impl) Request() request.Service {
	return i.request
}
