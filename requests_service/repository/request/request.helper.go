package request

import (
	pb "github.com/techvify-oliver/protoserver/generated"
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

type EmailRequest struct {
	from    string
	to      string
	subject string
	content string
}

func MapDataRequest(req *ent.Request, vendor *pb.GetGroupDetailsResponse, contract *pb.GetContractResponse, contractor *pb.GetGroupDetailsResponse) *model.WithdrawRequest {
	return &model.WithdrawRequest{
		ID: req.ID,
		Contract: &model.Contract{
			ID: req.ContractID,
			Vendor: &model.Group{
				ID:      req.VendorID,
				Name:    &vendor.Name,
				Address: &vendor.Address,
			},
			ContractStatus:  &contract.Status,
			StartDate:       contract.StartDate.AsTime(),
			EndDate:         contract.EndDate.AsTime(),
			RemainingAmount: int(contract.RemainingAmount),
			TotalAmount:     int(contract.TotalAmount),
		},
		Vendor: &model.Group{
			ID:      req.VendorID,
			Name:    &vendor.Name,
			Address: &vendor.Address,
		},
		ProjectContractor: &model.Group{
			ID:      req.ContractorID,
			Name:    &contractor.Name,
			Address: &contractor.Address,
		},
		WithdrawRequestStatus: string(req.Status),
		Amount:                &req.Amount,
		CollectionDate:        req.CollectionDate,
	}
}
