package base

import (
	"github.com/spf13/viper"
)

type Config struct {
	App   *App   `mapstructure:"app" jsonschema:"title=App"`
	Oss   *Oss   `mapstructure:"oss" jsonschema:"title=oss"`
	Oauth *Oauth `mapstructure:"oauth" jsonschema:"title=oauth"`
}

type App struct {
	Name  string `mapstructure:"name" jsonschema:"title=Application Name"`
	Port  string `mapstructure:"port" jsonschema:"title=Application Port"`
	Host  string `mapstructure:"host" jsonschema:"title=Application Host"`
	Root  string `mapstructure:"root" jsonschema:"title=root"`
	Debug bool   `mapstructure:"debug" jsonschema:"title=Debug"`
}

type Oss struct {
	Vendor    string `mapstructure:"vendor"`
	Domain    string `mapstructure:"domain"`
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

func NewConfig(v *viper.Viper) (*Config, error) {
	c := &Config{}
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

type Jwt struct {
	Secret string `mapstructure:"secret"`
}

type Oauth struct {
	Jwt      Jwt    `mapstructure:"url"`
	Passkey  string `mapstructure:"passkey"`
	LoginUri string `mapstructure:"login_uri"`
}
