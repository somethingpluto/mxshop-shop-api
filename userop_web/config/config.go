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
type WebServiceConfig struct {
	Name          string `json:"name"`
	Host          string
	ConsulInfo    ConsulConfig        `json:"consul"`
	JWTInfo       JWTConfig           `json:"jwtConfig"`
	UseropService UserOPServiceConfig `json:"userop_service"`
	GoodsService  GoodsServiceConfig  `json:"goods_service"`
	JaegerInfo    JaegerConfig        `json:"jaeger"`
}

type ConsulConfig struct {
	Host string   `json:"host"`
	Port int      `json:"port"`
	Tags []string `json:"tags"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type JaegerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type UserOPServiceConfig struct {
	Name string `json:"name"`
}

type GoodsServiceConfig struct {
	Name string `json:"name"`
}

type FilePathConfig struct {
	ConfigFile string
	LogFile    string
}
