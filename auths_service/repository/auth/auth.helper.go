package auth

import (
	pb "github.com/techvify-oliver/protoserver/generated"
	authJWT "golang-boilerplate/internal/auth"
	"golang-boilerplate/model"
)

func MapDataAuthTokenSend(data *pb.GetUserResponse, token *authJWT.Token, username string) *model.AuthUserDetailsReturn {
	userGRPC := data.User
	groupGRPC := userGRPC.Group
	dataSend := &model.AuthUserDetailsReturn{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		User: &model.UserInfo{
			FirstName: userGRPC.FirstName,
			LastName:  userGRPC.LastName,
			Username:  username,
			Group:     groupGRPC.GroupName,
			Role:      &userGRPC.Role,
			GroupDetails: &model.GroupOfUser{
				ID:           int(groupGRPC.GroupId),
				Name:         groupGRPC.GroupName,
				Address:      groupGRPC.GroupAddress,
				EmailAddress: groupGRPC.GroupEmail,
			},
		},
		IsError:       false,
		MsgFromDev:    "",
		MsgFromServer: "",
	}
	return dataSend
}
