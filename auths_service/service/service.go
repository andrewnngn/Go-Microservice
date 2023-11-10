package service

import (
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/repository"
	"golang-boilerplate/service/auth"
)

// ServiceRegistry is the interface for the service registry.
type ServiceRegistry interface {
	Auth() auth.Service
}

// impl is the implementation of the service registry.
type impl struct {
	auth auth.Service
}

// New creates a new service registry.
func New(entClient *ent.Client, logger *zap.Logger, grpc *grpcClient.GRPCClientInit) ServiceRegistry {
	repoRegistry := repository.New(entClient, grpc)
	return &impl{
		auth: auth.New(repoRegistry, logger),
	}
}

// Auth returns the auth service.
func (i impl) Auth() auth.Service {
	return i.auth
}
