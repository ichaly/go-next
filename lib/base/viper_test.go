package base

import (
	"github.com/ichaly/go-next/lib/util"
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
