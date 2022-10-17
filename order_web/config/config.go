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
type WebServiceConfig struct {
	Name             string `json:"name"`
	Host             string
	ConsulInfo       ConsulConfig           `json:"consul"`
	JWTInfo          JWTConfig              `json:"jwtConfig"`
	GoodsService     GoodsServiceConfig     `json:"goods_service"`
	InventoryService InventoryServiceConfig `json:"inventory_service"`
	AlipayInfo       AlipayInfoConfig       `json:"alipay_info"`
	JaegerInfo       JaegerConfig           `json:"jaeger"`
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

type GoodsServiceConfig struct {
	Name string `json:"name"`
}

type InventoryServiceConfig struct {
	Name string `json:"name"`
}

type AlipayInfoConfig struct {
	AppID        string `mapstructure:"app_id" json:"app_id"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnURL    string `mapstructure:"return_url" json:"return_url"`
}
