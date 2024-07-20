package internal

type DatabaseConfig struct {
	AppConfig  `mapstructure:"app"`
	DataSource `mapstructure:"database"`
}
