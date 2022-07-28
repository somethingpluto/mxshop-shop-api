package config

type ServerConfig struct {
	Name        string            `mapstructure:"name"`
	UserService UserServiceConfig `mapstructure:"user_service"`
	UserServer  UserServerConfig  `mapstructure:"user_server"`
	JWTInfo     JwtConfig         `mapstructure:"jwt_config"`
	AliSms      AliSmsConfig      `mapstructure:"alisms_config"`
	Redis       RedisConfig       `mapstructure:"redis_config"`
	ConsulInfo  ConsulConfig      `mapstructure:"consul_config"`
}

// UserServiceConfig
// @Description: gRPC服务 主机 端口
//
type UserServiceConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

// UserServerConfig
// @Description: gin监听端口
//
type UserServerConfig struct {
	Port int `mapstructure:"port"`
}

// JwtConfig
// @Description: JWT
//
type JwtConfig struct {
	SigningKey string `mapstructure:"key"`
}

type AliSmsConfig struct {
	ApiKey       string `mapstructure:"key"`
	ApiSecret    string `mapstructure:"secret"`
	SignName     string `mapstructure:"signName"`
	TemplateCode string `mapstructure:"templateCode"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type FileConfig struct {
	ConfigFile string
	LogFile    string
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
