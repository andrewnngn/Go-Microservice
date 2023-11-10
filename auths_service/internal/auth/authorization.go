package authJWT

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	viperConfig "golang-boilerplate/internal/viper"
)

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}

func AuthSuccess(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}
	return nil
}

func AuthorizationRequired() fiber.Handler {
	configs := viperConfig.InitConfig()
	return jwtware.New(jwtware.Config{
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(configs.Secret.AccessTokenSecret),
		SigningMethod:  "HS256",
		TokenLookup:    "header:Authorization",
		AuthScheme:     "Bearer",
	})
}
