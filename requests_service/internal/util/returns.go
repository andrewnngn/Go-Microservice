package util

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

const (
	REQUEST_TYPE    = "CONTRACT"
	GET_ALL_REQUEST = "GET_ALL_REQUEST"
)

func WrapReturn(returnType string, data any, isError bool, msgFromDev string, msgFromServer error) any {
	if returnType == REQUEST_TYPE {
		return WrapRequestReturn(data, isError, msgFromDev, msgFromServer)
	}
	if returnType == GET_ALL_REQUEST {
		return WrapAllContractsReturn(data, isError, msgFromDev, msgFromServer)
	}
	return nil
}

func WrapRequestReturn(data any, isError bool, msgFromDev string, msgFromServer error) *model.RequestReturn {
	if isError {
		return &model.RequestReturn{
			Data:          &ent.Request{ID: -1},
			IsError:       true,
			MsgFromDev:    msgFromDev,
			MsgFromServer: msgFromServer.Error(),
		}
	}

	return &model.RequestReturn{
		Data:          data.(*ent.Request),
		IsError:       false,
		MsgFromDev:    msgFromDev,
		MsgFromServer: "",
	}
}

func WrapAllContractsReturn(data any, isError bool, msgFromDev string, msgFromServer error) *model.WithdrawRequestReturn {

	if isError {
		return &model.WithdrawRequestReturn{
			Data:          nil,
			IsError:       true,
			MsgFromDev:    msgFromDev,
			MsgFromServer: msgFromServer.Error(),
		}
	}

	// You can add more items to the 'result' slice as needed

	return &model.WithdrawRequestReturn{
		Data:          data.([]*model.WithdrawRequest),
		IsError:       false,
		MsgFromDev:    msgFromDev,
		MsgFromServer: "",
	}
}
