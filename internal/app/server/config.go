package server

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		//BindAddr: ":" + os.Getenv("PORT"),
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
