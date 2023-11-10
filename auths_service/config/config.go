package config

type Configurations struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Port     PortConfig     `mapstructure:"port"`
	Trivial  TrivialConfig  `mapstructure:"trivial"`
	Secret   SecretConfig   `mapstructure:"secret"`
}

type SecretConfig struct {
	AccessTokenSecret    string `mapstructure:"accessToken"`
	RefreshTokenSecret   string `mapstructure:"refreshToken"`
	AdminAccountUsername string `mapstructure:"adminAccountUsername"`
	AdminAccountPassword string `mapstructure:"adminAccountPassword"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}

type PortConfig struct {
	GraphqlPort    string `mapstructure:"graphqlPort"`
	GrpcClientPort string `mapstructure:"grpcClientPort"`
}

type TrivialConfig struct {
	AppName        string `mapstructure:"appName"`
	ServerHeader   string `mapstructure:"serverHeader"`
	PlaygroundName string `mapstructure:"playgroundName"`
}
