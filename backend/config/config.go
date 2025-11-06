package config

type DbConfig struct {
	UseSqlite  bool   `json:"useSqlite"  toml:"useSqlite"  yaml:"useSqlite"  kdl:"useSqlite"`
	SqliteFile string `json:"sqliteFile" toml:"sqliteFile" yaml:"sqliteFile" kdl:"sqliteFile"`
}

type Config struct {
	BindAddress string   `json:"bindAddress" toml:"bindAddress" yaml:"bindAddress" kdl:"bindAddress"`
	Db          DbConfig `json:"db"          toml:"db"          yaml:"db"          kdl:"db"`
}

var Global Config

var DefaultConfig = Config{
	BindAddress: ":8080",
	Db: DbConfig{
		UseSqlite:  true,
		SqliteFile: "./db.sqlite",
	},
}
