package contract

import (
	"context"
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/util"
	"golang-boilerplate/model"
)

// Repository is the interface for the contract repository.
type Repository interface {
	CreateContract(ctx context.Context, input ent.CreateContractInput) (*model.ContractReturn, error)
	UpdateContract(ctx context.Context, id int, input ent.UpdateContractInput) (*model.ContractReturn, error)
	DeleteContract(ctx context.Context, id int) (*model.ContractReturn, error)

	QueryAllContracts(ctx context.Context) (*model.ContractVendorReturn, error)
}

// impl is the implementation of the contract repository.
type impl struct {
	client *ent.Client
	grpc   *grpcClient.GRPCClientInit
	logg   *zap.Logger
}

// New creates a new contract repository.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit, logg *zap.Logger) Repository {
	return &impl{
		client: client,
		grpc:   grpc,
		logg:   logg,
	}
}

func (i impl) CreateContract(ctx context.Context, input ent.CreateContractInput) (*model.ContractReturn, error) {
	contract1, err := i.client.Contract.Create().SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.CONTRACT_TYPE, nil, true, "Create contract failed", err).(*model.ContractReturn), nil
	}

	return util.WrapReturn(util.CONTRACT_TYPE, contract1, false, "", err).(*model.ContractReturn), nil
}

func (i impl) UpdateContract(ctx context.Context, id int, input ent.UpdateContractInput) (*model.ContractReturn, error) {
	contract1, err := i.client.Contract.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.CONTRACT_TYPE, nil, true, "Update contract failed", err).(*model.ContractReturn), nil
	}

	return util.WrapReturn(util.CONTRACT_TYPE, contract1, false, "", err).(*model.ContractReturn), nil
}

func (i impl) DeleteContract(ctx context.Context, id int) (*model.ContractReturn, error) {
	contract1, err := i.client.Contract.Get(ctx, id)
	if err != nil {
		return util.WrapReturn(util.CONTRACT_TYPE, nil, true, "Get contract failed", err).(*model.ContractReturn), nil
	}

	err = i.client.Contract.DeleteOne(contract1).Exec(ctx)
	if err != nil {
		return util.WrapReturn(util.CONTRACT_TYPE, nil, true, "Delete contract failed", err).(*model.ContractReturn), nil
	}

	return util.WrapReturn(util.CONTRACT_TYPE, contract1, false, "", err).(*model.ContractReturn), nil
}

func (i impl) QueryAllContracts(ctx context.Context) (*model.ContractVendorReturn, error) {
	allResults, err := i.client.Contract.Query().All(ctx)
	if err != nil {
		return util.WrapReturn(util.GET_ALL_CONTRACT, nil, true, "Get all contracts failed", err).(*model.ContractVendorReturn), nil
	}
	var allContracts []*model.ContractVendor
	for _, contract := range allResults {

		dataGrpc, err := i.getGroupGRPC(i.grpc.Service, contract.VendorID)
		if err != nil {
			return util.WrapReturn(util.GET_ALL_CONTRACT, nil, true, "Get vendor of contract failed", err).(*model.ContractVendorReturn), nil
		}

		allContracts = append(allContracts, &model.ContractVendor{
			ID: contract.ID,
			Vendor: &model.Vendor{
				ID:   contract.VendorID,
				Name: dataGrpc.Name,
			},
			ContractStatus:  string(contract.Status),
			StartDate:       *contract.StartDate,
			EndDate:         *contract.EndDate,
			RemainingAmount: contract.RemainingAmount,
			TotalAmount:     contract.TotalAmount,
		})
	}

	return util.WrapReturn(util.GET_ALL_CONTRACT, allContracts, false, "", nil).(*model.ContractVendorReturn), nil
}
