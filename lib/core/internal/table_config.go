package internal

type TableConfig struct {
	Prefixes []string `mapstructure:"Prefixes"`
}

type TableDefinition struct {
	Prefixes int `mapstructure:"port"`
}
