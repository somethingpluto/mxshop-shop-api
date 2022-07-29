package config

// ServerConfig
// @Description: 服务配置总
//
type ServerConfig struct {
	Name        string            `mapstructure:"name"`
	UserService UserServiceConfig `mapstructure:"user_service"`
	UserServer  UserServerConfig  `mapstructure:"user_server"`
	JWTInfo     JwtConfig         `mapstructure:"jwt_config"`
	AliSms      AliSmsConfig      `mapstructure:"alisms_config"`
	Redis       RedisConfig       `mapstructure:"redis_config"`
	ConsulInfo  ConsulConfig      `mapstructure:"consul_config"`
	RuntimeInfo RuntimeConfig     `mapstructure:"runtime_config"`
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
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// JwtConfig
// @Description: JWT
//
type JwtConfig struct {
	SigningKey string `mapstructure:"key"`
}

// AliSmsConfig
// @Description: 阿里云短信服务配置
//
type AliSmsConfig struct {
	ApiKey       string `mapstructure:"key"`
	ApiSecret    string `mapstructure:"secret"`
	SignName     string `mapstructure:"signName"`
	TemplateCode string `mapstructure:"templateCode"`
}

// RedisConfig
// @Description: Redis配置
//
type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// FileConfig
// @Description: 文件路劲配置
//
type FileConfig struct {
	ConfigFile string
	LogFile    string
}

// ConsulConfig
// @Description: Consul配置
//
type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type RuntimeConfig struct {
	Mode string `mapstructure:"mode"`
}
