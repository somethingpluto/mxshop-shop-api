package config

type ServerConfig struct {
	Name        string            `mapstructure:"name"`
	UserService UserServiceConfig `mapstructure:"user_service"`
	UserServer  UserServerConfig  `mapstructure:"user_server"`
}

type UserServiceConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type UserServerConfig struct {
	Port int `mapstructure:"port"`
}
