package internal

type TableConfig struct {
	Tables    []TableDefinition `mapstructure:"tables"`
	Mapping   map[string]string `mapstructure:"mapping"`
	UseCamel  bool              `mapstructure:"use-camel"`
	Prefixes  []string          `mapstructure:"prefixes"`
	BlockList []string          `mapstructure:"block-list"`
}

type TableDefinition struct {
	Name    string             `mapstructure:"name"`
	Type    string             `mapstructure:"type"`
	Table   string             `mapstructure:"table"`
	Columns []ColumnDefinition `mapstructure:"columns"`
}

type ColumnDefinition struct {
	Name      string `mapstructure:"name"`
	Type      string `mapstructure:"type"`
	RelatedTo string `mapstructure:"related-to"`
}
