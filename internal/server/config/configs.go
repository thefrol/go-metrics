package config

type Config struct {
	//address string
	Address string `env:"ADDRESS"`
}

func LoadConfig() *Config {
	var cfg Config
	return &cfg

}
