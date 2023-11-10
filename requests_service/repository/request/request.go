package request

import (
	"context"
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	"golang-boilerplate/ent/request"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/util"
	"golang-boilerplate/model"
	"strconv"
	"time"
)

// Repository is the interface for the request repository.
type Repository interface {
	CreateRequest(ctx context.Context, input ent.CreateRequestInput) (*model.RequestReturn, error)
	UpdateRequest(ctx context.Context, id int, input ent.UpdateRequestInput) (*model.RequestReturn, error)
	DeleteRequest(ctx context.Context, id int) (*model.RequestReturn, error)

	QueryAllRequests(ctx context.Context) (*model.WithdrawRequestReturn, error)

	ReadyToCollect(ctx context.Context, id int) (*model.WithdrawCollectReturn, error)
	Collected(ctx context.Context, id int) (*model.WithdrawCollectReturn, error)
	ArrangeCollectionDate(ctx context.Context, id int, collectionDate time.Time) string
}

// impl is the implementation of the request repository.
type impl struct {
	client *ent.Client
	grpc   *grpcClient.GRPCClientInit
	logg   *zap.Logger
}

func (i impl) ReadyToCollect(ctx context.Context, id int) (*model.WithdrawCollectReturn, error) {
	req, err := i.client.Request.UpdateOneID(id).SetStatus(request.StatusReadyToCollect).Save(ctx)
	if err != nil {
		return nil, err
	}

	vendor, err := i.getGroupGRPC(i.grpc.Service, req.VendorID)
	if err != nil {
		return nil, err
	}

	contractor, err := i.getGroupGRPC(i.grpc.Service, req.ContractorID)
	if err != nil {
		return nil, err
	}

	contract, err := i.getContractGRPC(i.grpc.Service, req.ContractID)
	if err != nil {
		return nil, err
	}

	return &model.WithdrawCollectReturn{
		Data:          MapDataRequest(req, vendor, contract, contractor),
		IsError:       false,
		MsgFromDev:    "Updated to ready to collect",
		MsgFromServer: "",
	}, nil
}

func (i impl) Collected(ctx context.Context, id int) (*model.WithdrawCollectReturn, error) {
	req, err := i.client.Request.UpdateOneID(id).SetStatus(request.StatusCollected).Save(ctx)
	if err != nil {
		return nil, err
	}

	vendor, err := i.getGroupGRPC(i.grpc.Service, req.VendorID)
	if err != nil {
		return nil, err
	}

	contractor, err := i.getGroupGRPC(i.grpc.Service, req.ContractorID)
	if err != nil {
		return nil, err
	}

	contract, err := i.getContractGRPC(i.grpc.Service, req.ContractID)
	if err != nil {
		return nil, err
	}
	if i.removeAmountCollectedGRPCSuccess(i.grpc.Service, req.ContractID, req.Amount) {
		contract.RemainingAmount -= int32(req.Amount)
		return &model.WithdrawCollectReturn{
			Data:          MapDataRequest(req, vendor, contract, contractor),
			IsError:       false,
			MsgFromDev:    "Collected to cable drums successfully",
			MsgFromServer: "",
		}, nil
	}

	return &model.WithdrawCollectReturn{
		Data:          nil,
		IsError:       true,
		MsgFromDev:    "Collect failed for some reasons",
		MsgFromServer: err.Error(),
	}, nil
}

func (i impl) ArrangeCollectionDate(ctx context.Context, id int, collectionDate time.Time) string {
	_, err := i.client.Request.UpdateOneID(id).SetCollectionDate(collectionDate).Save(ctx)
	if err != nil {
		return err.Error()
	}
	return "Arrange collection date successfully"
}

// New creates a new request repository.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit, logg *zap.Logger) Repository {
	return &impl{
		client: client,
		grpc:   grpc,
		logg:   logg,
	}
}

func (i impl) CreateRequest(ctx context.Context, input ent.CreateRequestInput) (*model.RequestReturn, error) {

	contractor, err := i.getGroupGRPC(i.grpc.Service, input.ContractorID)
	if err != nil {
		return util.WrapReturn(util.REQUEST_TYPE, nil, true, "Contractor not exist", err).(*model.RequestReturn), nil
	}

	_, err = i.getGroupGRPC(i.grpc.Service, input.VendorID)
	if err != nil {
		return util.WrapReturn(util.REQUEST_TYPE, nil, true, "Vendor not exist", err).(*model.RequestReturn), nil
	}

	_, err = i.getContractGRPC(i.grpc.Service, input.ContractID)
	if err != nil {
		return util.WrapReturn(util.REQUEST_TYPE, nil, true, "Contract not exist", err).(*model.RequestReturn), nil
	}

	request2, err := i.client.Request.Create().SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.REQUEST_TYPE, nil, true, "Create request failed", err).(*model.RequestReturn), nil
	}

	emailSend := EmailRequest{
		from:    "system@microservice.com",
		to:      contractor.Email,
		subject: "A new withdrawal request has been created with id " + strconv.Itoa(request2.ID) + ".",
		content: "A withdrawal request has been created for " + strconv.Itoa(request2.Amount) + " cable drums. Please check the status of the request.",
	}

	sent := i.sendMailToMailService(i.grpc.Service, emailSend)
	if !sent {
		return util.WrapReturn(util.REQUEST_TYPE, nil, true, "Send email failed", err).(*model.RequestReturn), nil
	}

	return util.WrapReturn(util.REQUEST_TYPE, request2, false, "", err).(*model.RequestReturn), nil
}

func (i impl) UpdateRequest(ctx context.Context, id int, input ent.UpdateRequestInput) (*model.RequestReturn, error) {
	request1, err := i.client.Request.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.REQUEST_TYPE, nil, true, "Update request failed", err).(*model.RequestReturn), nil
	}

	return util.WrapReturn(util.REQUEST_TYPE, request1, false, "", err).(*model.RequestReturn), nil
}

func (i impl) DeleteRequest(ctx context.Context, id int) (*model.RequestReturn, error) {
	request1, err := i.client.Request.Get(ctx, id)
	if err != nil {
		return util.WrapReturn(util.REQUEST_TYPE, nil, true, "Get request failed", err).(*model.RequestReturn), nil
	}

	err = i.client.Request.DeleteOne(request1).Exec(ctx)
	if err != nil {
		return util.WrapReturn(util.REQUEST_TYPE, nil, true, "Delete request failed", err).(*model.RequestReturn), nil
	}

	return util.WrapReturn(util.REQUEST_TYPE, request1, false, "", err).(*model.RequestReturn), nil
}

func (i impl) QueryAllRequests(ctx context.Context) (*model.WithdrawRequestReturn, error) {
	allResults, err := i.client.Request.Query().All(ctx)
	if err != nil {
		return util.WrapReturn(util.GET_ALL_REQUEST, nil, true, "Get all requests failed", err).(*model.WithdrawRequestReturn), nil
	}
	var allRequest []*model.WithdrawRequest
	for _, req := range allResults {

		vendor, err := i.getGroupGRPC(i.grpc.Service, req.VendorID)
		if err != nil {
			return util.WrapReturn(util.GET_ALL_REQUEST, nil, true, "Get vendor failed", err).(*model.WithdrawRequestReturn), nil
		}
		contractor, err := i.getGroupGRPC(i.grpc.Service, req.ContractorID)
		if err != nil {
			return util.WrapReturn(util.GET_ALL_REQUEST, nil, true, "Get contractor failed", err).(*model.WithdrawRequestReturn), nil
		}
		contract, err := i.getContractGRPC(i.grpc.Service, req.ContractID)

		if err != nil {
			return util.WrapReturn(util.GET_ALL_REQUEST, nil, true, "Get contract failed", err).(*model.WithdrawRequestReturn), nil
		}

		allRequest = append(allRequest, MapDataRequest(req, vendor, contract, contractor))
	}

	return util.WrapReturn(util.GET_ALL_REQUEST, allRequest, false, "", nil).(*model.WithdrawRequestReturn), nil

}
