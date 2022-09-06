package config

// NacosConfig
// @Description: 配置中心nacos
//
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Dataid    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}

// WebApiConfig
// @Description: nacos拉取的配置
//
type WebApiConfig struct {
	Name       string       `json:"name"`
	Mode       string       `json:"runMode"`
	ConsulInfo ConsulConfig `json:"consulConfig"`
	JWTInfo    JWTConfig    `json:"jwtConfig"`
}

type ConsulConfig struct {
	Host string   `json:"host"`
	Port int      `json:"port"`
	Tags []string `json:"tags"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
