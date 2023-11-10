package auth

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
	"go.uber.org/zap"
	"golang-boilerplate/internal/logger"
)

func (i impl) getUserGRPC(c pb.ServiceClient, userID int) (*pb.GetUserResponse, error) {
	logger.NewLogger().Info("sending to grpc api gateway ->", zap.Any("userID", userID))
	input := &pb.GetUserRequest{
		UserId: int64(userID),
	}

	res, err := c.GetUser(context.Background(), input)

	if err != nil || res.IsError {
		return res, err
	}
	logger.NewLogger().Info("response from grpc server ->", zap.Any("res", res))
	return res, nil
}
