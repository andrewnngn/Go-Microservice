package request

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
)

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

func (i impl) removeAmountCollectedGRPCSuccess(c pb.ServiceClient, contractID int, amount int) bool {
	//logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("contractID", contractID))
	input := &pb.RemoveAmountAfterCollectedRequest{
		ContractId: int64(contractID),
		Amount:     int32(amount),
	}

	res, err := c.RemoveAmountFromContractAfterCollected(context.Background(), input)
	if err != nil || !res.IsSuccess {
		return false
	}
	//logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return true
}

func (i impl) sendMailToMailService(c pb.ServiceClient, email EmailRequest) bool {
	//logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("contractID", contractID))
	input := &pb.EmailRequest{
		From:    email.from,
		To:      email.to,
		Subject: email.subject,
		Content: email.content,
	}

	res, err := c.SendMailToMailService(context.Background(), input)
	if err != nil || !res.IsSent {
		return false
	}
	//logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return true
}
