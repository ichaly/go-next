package internal

type TableConfig struct {
	Tables      []TableDefinition `mapstructure:"tables"`
	UseCamel    bool              `mapstructure:"use-camel"`
	Prefixes    []string          `mapstructure:"prefixes"`
	BlockList   []string          `mapstructure:"block-list"`
	TypeMapping map[string]string `mapstructure:"type-mapping"`
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
