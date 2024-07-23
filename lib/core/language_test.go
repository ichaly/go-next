package core

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type _LanguageSuite struct {
	_MetadataSuite
	m *Metadata
}

func TestLanguage(t *testing.T) {
	suite.Run(t, new(_LanguageSuite))
}

func (my *_LanguageSuite) SetupSuite() {
	my._MetadataSuite.SetupSuite()

	var err error
	my.m, err = NewMetadata(my.v, my.d)
	my.Require().NoError(err)
}

func (my *_LanguageSuite) TestLanguage() {
	language, err := NewLanguage(my.m)
	my.Require().NoError(err)
	my.T().Log(language.build())
}
