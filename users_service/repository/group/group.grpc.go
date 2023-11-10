package group

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
)

func (i impl) getContractsGRPCByVendor(c pb.ServiceClient, vendorID int) (*pb.GetContractByGroupIDResponse, error) {
	//logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("contractID", contractID))
	input := &pb.GetContractByGroupIDRequest{
		GroupId: int64(vendorID),
	}

	res, err := c.GetContractsByVendorID(context.Background(), input)

	if err != nil {
		return nil, err
	}
	//logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return res, nil
}

func (i impl) getRequestsGRPCByContractor(c pb.ServiceClient, groupID int) (*pb.GetRequestByGroupIDResponse, error) {
	//logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("contractID", contractID))
	input := &pb.GetRequestByGroupIDRequest{
		GroupId: int64(groupID),
	}
	res, err := c.GetRequestsByContractorID(context.Background(), input)
	if err != nil {
		return nil, err
	}
	//logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return res, nil
}

func (i impl) getRequestsGRPCByVendor(c pb.ServiceClient, groupID int) (*pb.GetRequestByGroupIDResponse, error) {
	//logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("contractID", contractID))
	input := &pb.GetRequestByGroupIDRequest{
		GroupId: int64(groupID),
	}

	res, err := c.GetRequestsByVendorID(context.Background(), input)

	if err != nil {
		return nil, err
	}
	//logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return res, nil
}

func (i impl) getGroupGRPC(c pb.ServiceClient, groupID int) (*pb.GetGroupDetailsResponse, error) {
	//logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("groupID", groupID))
	input := &pb.GetGroupDetailsRequest{
		GroupId: int64(groupID),
	}

	res, err := c.GetGroupDetails(context.Background(), input)

	if err != nil {
		return nil, err
	}
	//logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return res, nil
}

func (i impl) getContractGRPC(c pb.ServiceClient, contractID int) (*pb.GetContractResponse, error) {
	//logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("contractID", contractID))
	input := &pb.GetContractRequest{
		ContractId: int64(contractID),
	}
	res, err := c.GetContractDetails(context.Background(), input)
	if err != nil {
		return nil, err
	}
	//logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return res, nil
}
