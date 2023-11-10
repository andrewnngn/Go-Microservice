package config

type Configurations struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Port     PortConfig     `mapstructure:"port"`
	Trivial  TrivialConfig  `mapstructure:"trivial"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}

type PortConfig struct {
	GraphqlPort    string `mapstructure:"graphqlPort"`
	GrpcClientPort string `mapstructure:"grpcClientPort"`
	GrpcServerPort string `mapstructure:"grpcServerPort"`
}

type TrivialConfig struct {
	AppName        string `mapstructure:"appName"`
	ServerHeader   string `mapstructure:"serverHeader"`
	PlaygroundName string `mapstructure:"playgroundName"`
}
