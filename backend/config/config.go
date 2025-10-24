package config

type DbConfig struct {
}

type Config struct {
	BindAddress string   `json:"bindAddress" toml:"bindAddress" yaml:"bindAddress" kml:"bindAddress"`
	Db          DbConfig `json:"db"          toml:"db"          yaml:"db"          kml:"db"`
}

var Global Config

var DefaultConfig = Config{
	BindAddress: ":8080",
	Db:          DbConfig{},
}
