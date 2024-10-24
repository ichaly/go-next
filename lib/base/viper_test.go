package base

import (
	"github.com/ichaly/go-next/lib/util"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"reflect"
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

type InternalConfig struct {
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

type BaseConfig struct {
	InternalConfig `mapstructure:"app"`
}

func TestViper(t *testing.T) {
	var s struct {
		BaseConfig `mapstructure:",squash"`
		CoreConfig `mapstructure:"schema"`
	}
	viper.SetDefault("schema.default-limit", 10)
	viper.SetConfigFile("../../cfg/dev.yml")
	if err := viper.ReadInConfig(); err != nil {
		println(err)
	}
	if err := viper.Unmarshal(&s, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			return data, nil
		},
		func(f reflect.Kind, t reflect.Kind, data interface{}) (interface{}, error) {
			return data, nil
		},
		func(f reflect.Value, t reflect.Value) (interface{}, error) {
			return f.Interface(), nil
		},
	))); err != nil {
		println(err)
	}
	println(&s, viper.AllSettings())
}
