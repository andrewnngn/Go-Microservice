package grpcGW

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
	"log"
)

func (s *Server) GetGroupDetails(ctx context.Context, in *pb.GetGroupDetailsRequest) (*pb.GetGroupDetailsResponse, error) {
	log.Printf("GetGroupDetails ....... -> %v\n", in)
	log.Printf("api gateway get req from client -> %v\n", in)
	data, err := s.Client.ServiceUser.GetGroupDetails(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("api gateway send req to server -> %v\n", data)
	return data, nil
}
