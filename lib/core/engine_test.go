package core

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type _engineSuite struct {
	suite.Suite
	d *gorm.DB
	v *viper.Viper
	m *Metadata
	c *Compiler
}

func TestEngine(t *testing.T) {
	suite.Run(t, new(_engineSuite))
}

func (my *_engineSuite) SetupSuite() {
	v, err := base.NewViper("../../cfg/dev.yml")
	my.Require().NoError(err)
	d, err := base.NewConnect(v, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{})
	my.Require().NoError(err)
	m, err := NewMetadata(d, v)
	my.Require().NoError(err)
	c, err := NewCompiler(m, d)
	my.Require().NoError(err)

	my.v = v
	my.d = d
	my.m = m
	my.c = c
}

func (my *_engineSuite) TestEngine() {
	e, err := NewEngine(my.m, my.c)
	my.Require().NoError(err)

	s, err := e.Schema()
	my.Require().NoError(err)

	params := graphql.Params{Schema: s, RequestString: `{ user { id } }`}
	r := graphql.Do(params)

	str, err := util.MarshalJson(r)
	my.Require().NoError(err)

	my.T().Log(str)
}
