package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Dataid    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

// WebServiceConfig
// @Description:
//
type WebServiceConfig struct {
	Name            string            `json:"name"`
	JWTInfo         JwtConfig         `json:"jwt"`
	AliSmsInfo      AliSmsConfig      `json:"aliyun_message"`
	RedisInfo       RedisConfig       `json:"redis"`
	ConsulInfo      ConsulConfig      `json:"consul"`
	JaegerInfo      JaegerConfig      `json:"jaeger_info"`
	UserServiceInfo UserServiceConfig `json:"user_service_info"`
}

// JwtConfig
// @Description: JWT
//
type JwtConfig struct {
	SigningKey string `json:"key"`
}

// AliSmsConfig
// @Description: 阿里云短信服务配置
//
type AliSmsConfig struct {
	ApiKey       string `json:"key"`
	ApiSecret    string `json:"secret"`
	SignName     string `json:"signName"`
	TemplateCode string `json:"template_code"`
}

// ConsulConfig
// @Description: Consul配置
//
type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// RedisConfig
// @Description: Redis配置
//
type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}
type WebConfig struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type JaegerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type UserServiceConfig struct {
	Name string `json:"name"`
}

// FileConfig
// @Description: 文件路劲配置
//
type FileConfig struct {
	ConfigFile string
	LogFile    string
}
