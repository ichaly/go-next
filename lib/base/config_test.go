package base

import (
	"github.com/ichaly/go-next/lib/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

var v *viper.Viper

func init() {
	val, err := NewViper("../../cfg/dev.yml")
	if err != nil {
		panic(err)
	}
	v = val
}

func TestNewConfig(t *testing.T) {
	config, err := NewConfig(v)
	assert.NoError(t, err)
	assert.NotNil(t, config)

	str, err := util.MarshalJson(config)
	t.Log(str)
}
