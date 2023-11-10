package grpcGW

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
	"log"
)

func (s *Server) GetContractDetails(ctx context.Context, in *pb.GetContractRequest) (*pb.GetContractResponse, error) {
	log.Printf("GetContractDetails ....... -> %v\n", in)
	log.Printf("api gateway get req from client -> %v\n", in)
	data, err := s.Client.ServiceContract.GetContractDetails(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("api gateway send req to server -> %v\n", data)
	return data, nil
}

func (s *Server) GetContractsByVendorID(ctx context.Context, in *pb.GetContractByGroupIDRequest) (*pb.GetContractByGroupIDResponse, error) {
	log.Printf("GetContractsByVendorID ....... -> %v\n", in)
	log.Printf("api gateway get req from client -> %v\n", in)
	data, err := s.Client.ServiceContract.GetContractsByVendorID(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("api gateway send req to server -> %v\n", data)
	return data, nil
}

func (s *Server) RemoveAmountFromContractAfterCollected(ctx context.Context, in *pb.RemoveAmountAfterCollectedRequest) (*pb.RemoveAmountAfterCollectedResponse, error) {
	log.Printf("RemoveAmountFromContractAfterCollected ....... -> %v\n", in)
	log.Printf("api gateway get req from client -> %v\n", in)
	data, err := s.Client.ServiceContract.RemoveAmountFromContractAfterCollected(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("api gateway send req to server -> %v\n", data)
	return data, nil
}
