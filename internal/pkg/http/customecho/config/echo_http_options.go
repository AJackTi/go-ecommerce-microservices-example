package config

type EchoHttpOptions struct {
	Port               int      `mapstructure:"port" validate:"required" env:"TcpPort"`
	Development        bool     `mapstructure:"development" env:"Development"`
	BasePath           string   `mapstructure:"basePath" validate:"required" env:"BasePath"`
	DebugErrorResponse bool     `mapstructure:"debugErrorResponse" env:"DebugErrorResponse"`
	IgnoreLogUrls      []string `mapstructure:"ignoreLogUrls"`
	Timeout            int      `mapstructure:"timeout" env:"Timeout"`
	Host               string   `mapstructure:"host" env:"Host"`
	Name               string   `mapstructure:"name" env:"ShortTypeName"`
}
