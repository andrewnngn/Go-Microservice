package authJWT

import viperConfig "golang-boilerplate/internal/viper"

func IsAdmin(username string, password string) bool {
	configs := viperConfig.InitConfig()
	if username == configs.Secret.AdminAccountUsername && password == configs.Secret.AdminAccountPassword {
		return true
	}
	return false
}
