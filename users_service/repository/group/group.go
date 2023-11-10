package group

import (
	"context"
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/util"
	"golang-boilerplate/model"
	"log"
	"time"
)

// Repository is the interface for the group repository.
type Repository interface {
	CreateGroup(ctx context.Context, input ent.CreateGroupInput) (*model.GroupReturn, error)
	UpdateGroup(ctx context.Context, id int, input ent.UpdateGroupInput) (*model.GroupReturn, error)
	DeleteGroup(ctx context.Context, id int) (*model.GroupReturn, error)

	ContractorByID(ctx context.Context, id int) (*model.ContractorC, error)
	VendorByID(ctx context.Context, id int) (*model.VendorReturn, error)
}

// impl is the implementation of the group repository.
type impl struct {
	client *ent.Client
	grpc   *grpcClient.GRPCClientInit
	logg   *zap.Logger
}

func (i impl) ContractorByID(ctx context.Context, id int) (*model.ContractorC, error) {
	ctr, err := i.client.Group.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	requests1, err := i.getRequestsGRPCByContractor(i.grpc.Service, ctr.ID)
	var res []*model.WithdrawRequestC
	log.Println(requests1)
	for _, v := range requests1.WithdrawRequests {
		contract1, err := i.getContractGRPC(i.grpc.Service, int(v.ContractId))
		if err != nil {
			return nil, err
		}

		vendor1, err := i.getGroupGRPC(i.grpc.Service, int(v.VendorId))
		if err != nil {
			return nil, err
		}
		res = append(res, &model.WithdrawRequestC{
			ID: int(v.Id),
			ContractC: &model.ContractC{
				ID:              int(v.ContractId),
				StartDate:       contract1.StartDate.AsTime(),
				EndDate:         contract1.EndDate.AsTime(),
				RemainingAmount: int(contract1.RemainingAmount),
				TotalAmount:     int(contract1.TotalAmount),
			},
			VendorC: &model.VendorC{
				ID:      int(v.VendorId),
				Name:    vendor1.Name,
				Address: vendor1.Address,
			},
		})
	}

	return &model.ContractorC{
		ID:               ctr.ID,
		Name:             ctr.Name,
		Address:          ctr.Address,
		WithdrawRequests: res,
		IsError:          false,
		MsgFromDev:       "",
		MsgFromServer:    "",
	}, nil
}

func (i impl) VendorByID(ctx context.Context, id int) (*model.VendorReturn, error) {
	vdr, err := i.client.Group.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	var contracts []*model.ContractV
	var requests []*model.WithdrawRequestV
	contracts1, err := i.getContractsGRPCByVendor(i.grpc.Service, vdr.ID)
	request1, err := i.getRequestsGRPCByVendor(i.grpc.Service, vdr.ID)

	for _, v := range contracts1.Contracts {
		contracts = append(contracts, &model.ContractV{
			ID:        int(v.Id),
			StartDate: v.StartDate.AsTime(),
			EndDate:   v.EndDate.AsTime(),
		})
	}

	for _, v := range request1.WithdrawRequests {
		if v.CollectionDate != nil {
			requests = append(requests, &model.WithdrawRequestV{
				ID:             int(v.Id),
				CollectionDate: v.CollectionDate.AsTime(),
			})
		} else {
			requests = append(requests, &model.WithdrawRequestV{
				ID:             int(v.Id),
				CollectionDate: time.Date(1700, 12, 12, 12, 12, 12, 12, time.UTC),
			})
		}

	}
	return &model.VendorReturn{
		ID:                vdr.ID,
		Name:              vdr.Name,
		Address:           vdr.Address,
		ContractsV:        contracts,
		WithdrawRequestsV: requests,
		IsError:           false,
		MsgFromDev:        "",
		MsgFromServer:     "",
	}, nil
}

// New creates a new group repository.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit, logg *zap.Logger) Repository {
	return &impl{
		client: client,
		grpc:   grpc,
		logg:   logg,
	}
}

func (i impl) CreateGroup(ctx context.Context, input ent.CreateGroupInput) (*model.GroupReturn, error) {
	group1, err := i.client.Group.Create().SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.GROUP_TYPE, nil, true, "Create group failed", err).(*model.GroupReturn), nil
	}

	return util.WrapReturn(util.GROUP_TYPE, group1, false, "", err).(*model.GroupReturn), nil
}

func (i impl) UpdateGroup(ctx context.Context, id int, input ent.UpdateGroupInput) (*model.GroupReturn, error) {
	group1, err := i.client.Group.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.GROUP_TYPE, nil, true, "Update group failed", err).(*model.GroupReturn), nil
	}

	return util.WrapReturn(util.GROUP_TYPE, group1, false, "", err).(*model.GroupReturn), nil
}

func (i impl) DeleteGroup(ctx context.Context, id int) (*model.GroupReturn, error) {
	group1, err := i.client.Group.Get(ctx, id)
	if err != nil {
		return util.WrapReturn(util.GROUP_TYPE, nil, true, "Get group failed", err).(*model.GroupReturn), nil
	}

	err = i.client.Group.DeleteOne(group1).Exec(ctx)
	if err != nil {
		return util.WrapReturn(util.GROUP_TYPE, nil, true, "Delete group failed", err).(*model.GroupReturn), nil
	}

	return util.WrapReturn(util.GROUP_TYPE, group1, false, "", err).(*model.GroupReturn), nil
}
