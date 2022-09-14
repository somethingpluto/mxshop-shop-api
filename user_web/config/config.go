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
	WebInfo     WebConfig     `json:"web"`
	JWTInfo     JwtConfig     `json:"jwt"`
	AliSmsInfo  AliSmsConfig  `json:"aliyun_message"`
	RedisInfo   RedisConfig   `json:"redis"`
	ConsulInfo  ConsulConfig  `json:"consul"`
	ServiceInfo ServiceConfig `json:"service"`
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
type ServiceConfig struct {
	Name string `json:"name"`
}

type WebConfig struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// FileConfig
// @Description: 文件路劲配置
//
type FileConfig struct {
	ConfigFile string
	LogFile    string
}
