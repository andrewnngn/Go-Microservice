package grpcGW

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
	"log"
)

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("GetUser ....... -> %v\n", in)
	log.Printf("api gateway get req from client -> %v\n", in)
	data, err := s.Client.ServiceUser.GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("api gateway send req to server -> %v\n", data)
	return data, nil
}
