package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	generated "golang-boilerplate/graphql"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client  *ent.Client
	service service.ServiceRegistry
	logger  *zap.Logger
	grpc    *grpcClient.GRPCClientInit
}

// NewExecutableSchema creates an ExecutableSchema instance.
func NewExecutableSchema(client *ent.Client, logger *zap.Logger, grpcConn *grpcClient.GRPCClientInit) graphql.ExecutableSchema {
	service1 := service.New(client, logger, grpcConn)

	config := generated.Config{
		Resolvers: &Resolver{client: client, service: service1, logger: logger, grpc: grpcConn},
	}

	return generated.NewExecutableSchema(config)
}
