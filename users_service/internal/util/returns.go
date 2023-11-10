package util

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/model"
)

const (
	USER_TYPE  = "USER"
	GROUP_TYPE = "GROUP"
)

func WrapReturn(returnType string, data any, isError bool, msgFromDev string, msgFromServer error) any {
	if returnType == USER_TYPE {
		return WrapUserReturn(data, isError, msgFromDev, msgFromServer)
	}
	if returnType == GROUP_TYPE {
		return WrapGroupReturn(data, isError, msgFromDev, msgFromServer)
	}
	return nil
}

func WrapUserReturn(data any, isError bool, msgFromDev string, msgFromServer error) *model.UserReturn {
	if isError {
		return &model.UserReturn{
			Data:          &ent.User{ID: -1},
			IsError:       true,
			MsgFromDev:    msgFromDev,
			MsgFromServer: msgFromServer.Error(),
		}
	}

	return &model.UserReturn{
		Data:          data.(*ent.User),
		IsError:       false,
		MsgFromDev:    msgFromDev,
		MsgFromServer: "",
	}
}

func WrapGroupReturn(data any, isError bool, msgFromDev string, msgFromServer error) *model.GroupReturn {
	if isError {
		return &model.GroupReturn{
			Data:          &ent.Group{ID: -1},
			IsError:       true,
			MsgFromDev:    msgFromDev,
			MsgFromServer: msgFromServer.Error(),
		}
	}

	return &model.GroupReturn{
		Data:          data.(*ent.Group),
		IsError:       false,
		MsgFromDev:    msgFromDev,
		MsgFromServer: "",
	}
}
