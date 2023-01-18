package config

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
	PostgresMaxConn  int32
}

func Load() Config {

	cfg := Config{}

	cfg.HTTPPort = ":9090"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "postgres"
	cfg.PostgresDatabase = "day9"
	cfg.PostgresPassword = "0003"
	cfg.PostgresPort = "3003"
	cfg.PostgresMaxConn = 30

	return cfg
}