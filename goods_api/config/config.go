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
	Mode       string       `json:"mode"`
	ConsulInfo ConsulConfig `json:"consulConfig"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}
