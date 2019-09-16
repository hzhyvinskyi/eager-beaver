package dal

// Config ...
type Config struct {
	databaseURL string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		databaseURL: "user=root password=dell dbname=beaver_db sslmode=disable",
	}
}
