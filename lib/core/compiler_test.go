package core

import (
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
}
