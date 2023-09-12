package envs

type Config struct {
	//address string
	Address        string `env:"ADDRESS"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	PollInterval   int    `env:"POLL_INTERVAL"`
}

func LoadConfig() *Config {
	var cfg Config
	return &cfg

}
