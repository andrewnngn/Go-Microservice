package auth

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	"golang-boilerplate/ent/auth"
	authJWT "golang-boilerplate/internal/auth"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/internal/util"
	"golang-boilerplate/model"
)

// Repository is the interface for the auth repository.
type Repository interface {
	GetToken(ctx context.Context, input model.LoginInput) (*model.AuthUserDetailsReturn, error)
	ExchangeToken(ctx context.Context, refreshToken string) (*model.AuthUserDetailsReturn, error)
	Logout(ctx context.Context, refreshToken string) (string, error)

	CreateAccount(ctx context.Context, input ent.CreateAuthInput) (*model.AuthReturn, error)
	UpdateAccount(ctx context.Context, id int, input ent.UpdateAuthInput) (*model.AuthReturn, error)
	DeleteAccount(ctx context.Context, accountID int) (*model.AuthReturn, error)

	// GRPC
	getUserGRPC(client pb.ServiceClient, userID int) (*pb.GetUserResponse, error)
}

// impl is the implementation of the auth repository.
type impl struct {
	client *ent.Client
	grpc   *grpcClient.GRPCClientInit
	logg   *zap.Logger
}

// New creates a new auth repository.
func New(client *ent.Client, grpc *grpcClient.GRPCClientInit, logg *zap.Logger) Repository {
	return &impl{
		client: client,
		grpc:   grpc,
		logg:   logg,
	}
}

func (i impl) CreateAccount(ctx context.Context, input ent.CreateAuthInput) (*model.AuthReturn, error) {
	passwordHash, err := util.GenerateFromPassword([]byte(input.Password), util.DefaultParams)
	if err != nil {
		return util.WrapReturn(util.AUTH_TYPE, nil, true, "Hash password failed", err).(*model.AuthReturn), nil
	}
	input.Password = string(passwordHash)

	_, err = i.client.Auth.Create().SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.AUTH_TYPE, nil, true, "Create account failed", err).(*model.AuthReturn), nil
	}

	user1, err := i.getUserGRPC(i.grpc.Service, input.UserID)
	if err != nil {
		return util.WrapReturn(util.AUTH_TYPE, nil, true, "Get user info failed", err).(*model.AuthReturn), nil
	}

	err = SendMailUserCreated(user1.User.Email)
	if err != nil {
		return nil, err
	}

	return util.WrapReturn(util.AUTH_TYPE, nil, false, "Create account successfully", err).(*model.AuthReturn), nil
}

func (i impl) UpdateAccount(ctx context.Context, accountID int, input ent.UpdateAuthInput) (*model.AuthReturn, error) {

	if input.Password != nil {
		passwordHash, err := util.GenerateFromPassword([]byte(*input.Password), util.DefaultParams)
		if err != nil {
			return util.WrapReturn(util.AUTH_TYPE, nil, true, "Hash password failed", err).(*model.AuthReturn), nil
		}
		*input.Password = string(passwordHash)
	}

	_, err := i.client.Auth.UpdateOneID(accountID).SetInput(input).Save(ctx)
	if err != nil {
		return util.WrapReturn(util.AUTH_TYPE, nil, true, "Update account failed", err).(*model.AuthReturn), nil
	}

	return util.WrapReturn(util.AUTH_TYPE, nil, false, "Update account successfully", err).(*model.AuthReturn), nil
}

func (i impl) DeleteAccount(ctx context.Context, accountID int) (*model.AuthReturn, error) {
	account1, err := i.client.Auth.Get(ctx, accountID)
	if err != nil {
		return util.WrapReturn(util.AUTH_TYPE, nil, true, "Get account failed", err).(*model.AuthReturn), nil
	}

	err = i.client.Auth.DeleteOne(account1).Exec(ctx)
	if err != nil {
		return util.WrapReturn(util.AUTH_TYPE, nil, true, "Delete account failed", err).(*model.AuthReturn), nil
	}

	return util.WrapReturn(util.AUTH_TYPE, nil, false, "Delete account successfully", err).(*model.AuthReturn), nil
}

// GetToken DONE
func (i impl) GetToken(ctx context.Context, input model.LoginInput) (*model.AuthUserDetailsReturn, error) {
	if authJWT.IsAdmin(input.Username, input.Password) {
		token, err := authJWT.CreateToken(nil, "admin", -1, -1, "admin")
		if err != nil {
			return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Create token failed", err).(*model.AuthUserDetailsReturn), nil
		}
		dataSend := &model.AuthUserDetailsReturn{
			AccessToken:   token.AccessToken,
			RefreshToken:  token.RefreshToken,
			User:          nil,
			IsError:       false,
			MsgFromDev:    "Admin Login Successfully",
			MsgFromServer: "",
		}
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, dataSend, false, "", nil).(*model.AuthUserDetailsReturn), nil
	}

	auth1, err := i.client.Auth.Query().Where(auth.Username(input.Username)).Select().First(ctx)
	if err != nil {
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Wrong Username", err).(*model.AuthUserDetailsReturn), nil
	}

	err = util.CompareHashAndPassword([]byte(auth1.Password), []byte(input.Password))
	if err != nil {
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Wrong Password", err).(*model.AuthUserDetailsReturn), nil
	}

	dataGrpc, err := i.getUserGRPC(i.grpc.Service, auth1.UserID)
	if err != nil {
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Get user info failed", err).(*model.AuthUserDetailsReturn), nil
	}

	token, err := authJWT.CreateToken(nil, dataGrpc.User.Role, auth1.ID, auth1.UserID, auth1.Username)
	if err != nil {
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Create token failed", err).(*model.AuthUserDetailsReturn), nil
	}

	// Map dataGrpc
	dataSend := MapDataAuthTokenSend(dataGrpc, &token, auth1.Username)
	return util.WrapReturn(util.AUTH_TOKEN_TYPE, dataSend, false, "", nil).(*model.AuthUserDetailsReturn), nil
}

func (i impl) ExchangeToken(_ context.Context, refreshToken string) (*model.AuthUserDetailsReturn, error) {
	payload, err := authJWT.VerifyToken(refreshToken, false)
	if err != nil {
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Refresh token not exist", err).(*model.AuthUserDetailsReturn), nil
	}

	dataGrpc, err := i.getUserGRPC(i.grpc.Service, int(payload["user_id"].(float64)))
	if err != nil {
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Get user info failed", err).(*model.AuthUserDetailsReturn), nil
	}

	token, err := authJWT.CreateToken(nil, dataGrpc.User.Role, int(payload["auth_id"].(float64)), int(payload["user_id"].(float64)), payload["username"].(string))
	if err != nil {
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Create token failed", err).(*model.AuthUserDetailsReturn), nil
	}

	_, err = i.client.Auth.UpdateOneID(int(payload["auth_id"].(float64))).SetRefreshToken(refreshToken).Save(context.Background())
	if err != nil {
		return util.WrapReturn(util.AUTH_TOKEN_TYPE, nil, true, "Save refresh token failed", err).(*model.AuthUserDetailsReturn), nil
	}

	// Map dataGrpc
	dataSend := MapDataAuthTokenSend(dataGrpc, &token, payload["username"].(string))
	return util.WrapReturn(util.AUTH_TOKEN_TYPE, dataSend, false, "", nil).(*model.AuthUserDetailsReturn), nil
}

func (i impl) Logout(ctx context.Context, refreshToken string) (string, error) {
	payload, err := authJWT.VerifyToken(refreshToken, false)
	if err != nil {
		return "Wrong refresh token", err
	}

	_, err = i.client.Auth.UpdateOneID(int(payload["auth_id"].(float64))).ClearRefreshToken().Save(ctx)
	if err != nil {
		return "Logout failed", err
	}

	return "Logout successfully", nil
}
