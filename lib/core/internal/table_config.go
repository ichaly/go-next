package internal

type TableConfig struct {
	Prefixes  []string `mapstructure:"prefixes"`
	BlockList []string `mapstructure:"block-list"`
	UseCamel  bool     `mapstructure:"use-camel"`
}

type TableDefinition struct {
	Prefixes int `mapstructure:"port"`
}
