package config

type ServerConfig struct {
	Name        string            `mapstructure:"name"`
	UserService UserServiceConfig `mapstructure:"user_service"`
	UserServer  UserServerConfig  `mapstructure:"user_server"`
}

// UserServiceConfig
// @Description: gRPC服务 主机 端口
//
type UserServiceConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// UserServerConfig
// @Description: gin监听端口
//
type UserServerConfig struct {
	Port int `mapstructure:"port"`
}
