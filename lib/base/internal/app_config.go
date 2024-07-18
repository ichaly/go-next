package internal

type AppConfig struct {
	Name  string `mapstructure:"name" jsonschema:"title=Application Name"`
	Port  string `mapstructure:"port" jsonschema:"title=Application Port"`
	Host  string `mapstructure:"host" jsonschema:"title=Application Host"`
	Root  string `mapstructure:"root" jsonschema:"title=root"`
	Debug bool   `mapstructure:"debug" jsonschema:"title=Debug"`
}
