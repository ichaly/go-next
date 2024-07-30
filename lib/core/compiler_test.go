package core

import (
	"github.com/ichaly/go-next/lib/core/internal/intro"
	"github.com/ichaly/go-next/lib/util"
	"github.com/stretchr/testify/suite"
	"testing"
)

type _CompilerSuite struct {
	_MetadataSuite
	m *Metadata
}

func TestCompiler(t *testing.T) {
	suite.Run(t, new(_CompilerSuite))
}

func (my *_CompilerSuite) SetupSuite() {
	my._MetadataSuite.SetupSuite()

	var err error
	my.m, err = NewMetadata(my.v, my.d)
	my.Require().NoError(err)
}

func (my *_CompilerSuite) TestCompiler() {
	c, err := NewCompiler(my.m, my.d)
	my.Require().NoError(err)
	query := "{user{id}team{id}}"
	_, _ = c.Compile(query)
}

func (my *_CompilerSuite) TestIntrospection() {
	c, err := NewCompiler(my.m, my.d)
	my.Require().NoError(err)

	s := intro.New(c.schema)
	my.T().Log(util.MustMarshalJson(s))
}
