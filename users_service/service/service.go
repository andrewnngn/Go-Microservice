package service

import (
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/repository"
	"golang-boilerplate/service/group"
	"golang-boilerplate/service/user"
)

// ServiceRegistry is the interface for the service registry.
type ServiceRegistry interface {
	User() user.Service
	Group() group.Service
}

// impl is the implementation of the service registry.
type impl struct {
	user  user.Service
	group group.Service
}

// New creates a new service registry.
func New(entClient *ent.Client, logger *zap.Logger, grpc *grpcClient.GRPCClientInit) ServiceRegistry {
	repoRegistry := repository.New(entClient, grpc)
	return &impl{
		user:  user.New(repoRegistry, logger),
		group: group.New(repoRegistry, logger),
	}
}

// User returns the user service.
func (i impl) User() user.Service {
	return i.user
}

// Group returns the group service.
func (i impl) Group() group.Service {
	return i.group
}
