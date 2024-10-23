package base

import (
	"github.com/ichaly/go-next/lib/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestConfig struct {
	Config `mapstructure:",squash"`
	Test   string
}

func TestNewViper(t *testing.T) {
	v, err := NewViper("../../cfg/dev.yml")

	assert.NoError(t, err)

	c := &TestConfig{}
	err = v.Unmarshal(c)
	assert.NoError(t, err)

	val, err := util.MarshalJson(c)
	t.Log(c.Name, val)
}

type BaseConfig struct {
	Name  string `mapstructure:"name"`
	Port  string `mapstructure:"port"`
	Host  string `mapstructure:"host"`
	Root  string `mapstructure:"root"`
	Debug bool   `mapstructure:"debug"`
}

type CoreConfig struct {
	Mapping      map[string]string `mapstructure:"mapping"`
	UseCamel     bool              `mapstructure:"use-camel"`
	Prefixes     []string          `mapstructure:"prefixes"`
	BlockList    []string          `mapstructure:"block-list"`
	DefaultLimit int               `mapstructure:"default-limit"`
}

func TestViper(t *testing.T) {
	var b BaseConfig
	var c CoreConfig
	viper.SetDefault("schema.default-limit", 10)
	viper.SetConfigFile("../../cfg/dev.yml")
	if err := viper.ReadInConfig(); err != nil {
		println(err)
	}
	if err := viper.Unmarshal(&b); err != nil {
		return
	}
	println(&b, &c, viper.AllSettings())
}
