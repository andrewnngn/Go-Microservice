package authJWT

import (
	"errors"
	viperConfig "golang-boilerplate/internal/viper"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

func CreateToken(payload any, role string, authID int, userID int, username string) (Token, error) {
	configs := viperConfig.InitConfig()
	var msgToken Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["payload"] = payload
	claims["username"] = username
	claims["role"] = role
	claims["auth_id"] = authID
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = time.Now().Unix()
	claims["https://hasura.io/jwt/claims"] = map[string]interface{}{
		"x-hasura-default-role":  role,
		"x-hasura-allowed-roles": []string{role},
		"x-hasura-user-id":       userID,
		"x-hasura-org-id":        configs.Trivial.AppName,
		"x-hasura-custom":        configs.Trivial.ServerHeader,
	}

	t, err := token.SignedString([]byte(configs.Secret.AccessTokenSecret))

	if err != nil {
		return msgToken, err
	}

	msgToken.AccessToken = t

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["payload"] = payload
	rtClaims["username"] = username
	rtClaims["role"] = role
	rtClaims["auth_id"] = authID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rtClaims["iat"] = time.Now().Unix()
	rtClaims["https://hasura.io/jwt/claims"] = map[string]interface{}{
		"x-hasura-default-role":  role,
		"x-hasura-allowed-roles": []string{role},
		"x-hasura-user-id":       userID,
		"x-hasura-org-id":        configs.Trivial.AppName,
		"x-hasura-custom":        configs.Trivial.ServerHeader,
	}

	rt, err := token.SignedString([]byte(configs.Secret.RefreshTokenSecret))

	if err != nil {
		return msgToken, err
	}

	msgToken.RefreshToken = rt
	return msgToken, nil
}

func VerifyToken(tokenStr string, isAccessToken bool) (jwt.MapClaims, error) {
	configs := viperConfig.InitConfig()

	var secret []byte
	if isAccessToken {
		secret = []byte(configs.Secret.AccessTokenSecret)
	} else {
		secret = []byte(configs.Secret.RefreshTokenSecret)
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the algorithm is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	// Convert token claims to jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
