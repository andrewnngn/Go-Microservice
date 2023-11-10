package util

import (
	"golang-boilerplate/model"
)

const (
	AUTH_TYPE       = "AUTH"
	AUTH_TOKEN_TYPE = "AUTH_TOKEN"
)

func WrapReturn(returnType string, data any, isError bool, msgFromDev string, msgFromServer error) any {
	if returnType == AUTH_TOKEN_TYPE {
		return WrapAuthUserDetailsReturn(data, isError, msgFromDev, msgFromServer)
	}
	if returnType == AUTH_TYPE {
		return WrapAuthUserReturn(isError, msgFromDev, msgFromServer)
	}
	return nil
}

func WrapAuthUserDetailsReturn(data any, isError bool, msgFromDev string, msgFromServer error) *model.AuthUserDetailsReturn {
	if isError {
		return &model.AuthUserDetailsReturn{
			AccessToken:  "err",
			RefreshToken: "err",
			User: &model.UserInfo{
				FirstName: "err",
				LastName:  "err",
				Username:  "err",
				Group:     "err",
				Role:      nil,
				GroupDetails: &model.GroupOfUser{
					ID:           0,
					Name:         "err",
					Address:      "err",
					EmailAddress: "err",
				},
			},
			IsError:       true,
			MsgFromDev:    msgFromDev,
			MsgFromServer: msgFromServer.Error(),
		}
	}

	return data.(*model.AuthUserDetailsReturn)
}

func WrapAuthUserReturn(isError bool, msgFromDev string, msgFromServer error) *model.AuthReturn {
	if isError {
		return &model.AuthReturn{
			IsError:       true,
			MsgFromDev:    msgFromDev,
			MsgFromServer: msgFromServer.Error(),
		}
	}

	return &model.AuthReturn{
		IsError:       false,
		MsgFromDev:    msgFromDev,
		MsgFromServer: "",
	}
}
