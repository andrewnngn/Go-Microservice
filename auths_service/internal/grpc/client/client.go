package grpcClient

import pb "github.com/techvify-oliver/protoserver/generated"

type GRPCClientInit struct {
	Service pb.ServiceClient
}
