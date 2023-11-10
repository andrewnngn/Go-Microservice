package grpcGW

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
	"log"
)

func (s *Server) GetRequestDetails(ctx context.Context, in *pb.GetWithdrawsRequest) (*pb.GetWithdrawsResponse, error) {
	log.Printf("GetRequestDetails ....... -> %v\n", in)
	log.Printf("api gateway get req from client -> %v\n", in)
	data, err := s.Client.ServiceRequest.GetRequestDetails(ctx, in)

	if err != nil {
		return nil, err
	}
	log.Printf("api gateway send req to server -> %v\n", data)
	return data, nil
}

func (s *Server) GetRequestsByContractorID(ctx context.Context, in *pb.GetRequestByGroupIDRequest) (*pb.GetRequestByGroupIDResponse, error) {
	log.Printf("GetRequestsByContractorID ....... -> %v\n", in)
	log.Printf("api gateway get req from client -> %v\n", in)
	data, err := s.Client.ServiceRequest.GetRequestsByContractorID(ctx, in)

	if err != nil {
		return nil, err
	}
	log.Printf("api gateway send req to server -> %v\n", data)
	return data, nil
}

func (s *Server) GetRequestsByVendorID(ctx context.Context, in *pb.GetRequestByGroupIDRequest) (*pb.GetRequestByGroupIDResponse, error) {
	log.Printf("GetRequestsByVendorID ....... -> %v\n", in)
	log.Printf("api gateway get req from client -> %v\n", in)
	data, err := s.Client.ServiceRequest.GetRequestsByVendorID(ctx, in)

	if err != nil {
		return nil, err
	}
	log.Printf("api gateway send req to server -> %v\n", data)
	return data, nil
}
