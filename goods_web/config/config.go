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

// WebApiConfig
// @Description: nacos拉取的配置
//
type WebApiConfig struct {
	Name       string       `json:"name"`
	Host       string       `json:"host"`
	ConsulInfo ConsulConfig `json:"consul"`
	JWTInfo    JWTConfig    `json:"jwt"`
	JaegerInfo JaegerConfig `json:"jaeger"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type JaegerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}
