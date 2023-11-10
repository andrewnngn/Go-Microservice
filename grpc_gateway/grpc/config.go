package grpcGW

import (
	pb "github.com/techvify-oliver/protoserver/generated"
	"grpc_server_testing/rabbit"
)

type GRPCClientInit struct {
	ServiceUser     pb.ServiceClient
	ServiceContract pb.ServiceClient
	ServiceRequest  pb.ServiceClient
}
type Server struct {
	pb.ServiceServer
	Client   *GRPCClientInit
	Rabbitmq *rabbit.RabbitMQClient
}
