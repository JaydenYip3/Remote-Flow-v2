package config

type Config struct {
	Address string
}

func Load() Config {
	address := ":9090"

	return Config{Address: address}
}
