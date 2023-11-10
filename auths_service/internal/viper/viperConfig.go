package viperConfig

import (
	"fmt"
	"github.com/spf13/viper"
	"golang-boilerplate/config"
	"os"
	"strings"
)

// InitConfig initializes the config.
func InitConfig() *config.Configurations {
	viper.SetConfigType("yaml")

	// Expand environment variables inside the config file
	const configFile = "config.yaml"
	b, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	expand := os.ExpandEnv(string(b))
	configReader := strings.NewReader(expand)

	viper.AutomaticEnv()

	if err := viper.ReadConfig(configReader); err != nil {
		fmt.Printf("read config has failed with error: %v\n", err)
		os.Exit(1)
	}

	configs := config.Configurations{}
	if err := viper.Unmarshal(&configs); err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	return &configs
}
