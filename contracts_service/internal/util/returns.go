package util

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

const (
	CONTRACT_TYPE    = "CONTRACT"
	GET_ALL_CONTRACT = "GET_ALL_CONTRACT"
)

func WrapReturn(returnType string, data any, isError bool, msgFromDev string, msgFromServer error) any {
	if returnType == CONTRACT_TYPE {
		return WrapContractReturn(data, isError, msgFromDev, msgFromServer)
	}
	if returnType == GET_ALL_CONTRACT {
		return WrapAllContractsReturn(data, isError, msgFromDev, msgFromServer)
	}
	return nil
}

func WrapContractReturn(data any, isError bool, msgFromDev string, msgFromServer error) *model.ContractReturn {
	if isError {
		return &model.ContractReturn{
			Data:          &ent.Contract{ID: -1},
			IsError:       true,
			MsgFromDev:    msgFromDev,
			MsgFromServer: msgFromServer.Error(),
		}
	}

	return &model.ContractReturn{
		Data:          data.(*ent.Contract),
		IsError:       false,
		MsgFromDev:    msgFromDev,
		MsgFromServer: "",
	}
}

func WrapAllContractsReturn(data any, isError bool, msgFromDev string, msgFromServer error) *model.ContractVendorReturn {

	if isError {
		return &model.ContractVendorReturn{
			Data:          nil,
			IsError:       true,
			MsgFromDev:    msgFromDev,
			MsgFromServer: msgFromServer.Error(),
		}
	}

	// You can add more items to the 'result' slice as needed

	return &model.ContractVendorReturn{
		Data:          data.([]*model.ContractVendor),
		IsError:       false,
		MsgFromDev:    msgFromDev,
		MsgFromServer: "",
	}
}
