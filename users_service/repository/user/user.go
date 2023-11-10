package user

import (
	"context"
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/util"
	"golang-boilerplate/model"
)

// Repository is the interface for the user repository.
type Repository interface {
	CreateUser(ctx context.Context, input ent.CreateUserInput) (*model.UserReturn, error)
	UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*model.UserReturn, error)
	DeleteUser(ctx context.Context, id int) (*model.UserReturn, error)
}

// impl is the implementation of the user repository.
type impl struct {
	client *ent.Client
	grpc   *grpcClient.GRPCClientInit
	logg   *zap.Logger
}

// New creates a new user repository.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit, logg *zap.Logger) Repository {
	return &impl{
		client: client,
		grpc:   grpc,
		logg:   logg,
	}
}

func (i impl) CreateUser(ctx context.Context, input ent.CreateUserInput) (*model.UserReturn, error) {

	//if input.GroupsID != nil {
	//	userGroup, err := i.client.Group.Get(ctx, *input.GroupsID)
	//	if err != nil {
	//		return util.WrapReturn(util.USER_TYPE, nil, true, "Group not found", err).(*model.UserReturn), nil
	//	}
	//	*input.GroupsID = userGroup.ID
	//}
	//
	//user1, err := i.client.User.Create().SetInput(input).Save(ctx)
	//if err != nil {
	//	return util.WrapReturn(util.USER_TYPE, nil, true, "Create user failed", err).(*model.UserReturn), nil
	//}
	//return util.WrapReturn(util.USER_TYPE, user1, false, "", err).(*model.UserReturn), nil
	return &model.UserReturn{
		Data:          nil,
		IsError:       false,
		MsgFromDev:    "h",
		MsgFromServer: "h",
	}, nil
}

func (i impl) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*model.UserReturn, error) {
	user1, err := i.client.User.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.USER_TYPE, nil, true, "Update user failed", err).(*model.UserReturn), nil
	}

	return util.WrapReturn(util.USER_TYPE, user1, false, "", err).(*model.UserReturn), nil
}

func (i impl) DeleteUser(ctx context.Context, id int) (*model.UserReturn, error) {
	user1, err := i.client.User.Get(ctx, id)
	if err != nil {
		return util.WrapReturn(util.USER_TYPE, nil, true, "Get user failed", err).(*model.UserReturn), nil
	}

	err = i.client.User.DeleteOne(user1).Exec(ctx)
	if err != nil {
		return util.WrapReturn(util.USER_TYPE, nil, true, "Delete user failed", err).(*model.UserReturn), nil
	}

	return util.WrapReturn(util.USER_TYPE, user1, false, "", err).(*model.UserReturn), nil
}
