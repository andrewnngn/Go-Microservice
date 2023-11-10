package contract

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
	"go.uber.org/zap"
	"golang-boilerplate/internal/logger"
)

func (i impl) getGroupGRPC(c pb.ServiceClient, groupID int) (*pb.GetGroupDetailsResponse, error) {
	logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("groupID", groupID))
	input := &pb.GetGroupDetailsRequest{
		GroupId: int64(groupID),
	}

	res, err := c.GetGroupDetails(context.Background(), input)

	if err != nil {
		return nil, err
	}
	logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return res, nil
}
