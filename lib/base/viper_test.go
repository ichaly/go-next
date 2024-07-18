package base

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewViper(t *testing.T) {

	// 调用 NewViper 函数
	v, err := NewViper("../../cfg/application.yml")

	// 断言错误为空
	assert.NoError(t, err)

	// 断言返回的 Viper 实例不为空
	assert.NotNil(t, v)
}
