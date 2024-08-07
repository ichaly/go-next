package core

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type _ExecutorSuite struct {
	_MetadataSuite
	m *Metadata
}

func TestExecutor(t *testing.T) {
	suite.Run(t, new(_ExecutorSuite))
}

func (my *_ExecutorSuite) SetupSuite() {
	my._MetadataSuite.SetupSuite()

	var err error
	my.m, err = NewMetadata(my.v, my.d)
	my.Require().NoError(err)
}

func (my *_ExecutorSuite) TestExecutor() {
	c, err := NewExecutor(my.m, my.d)
	my.Require().NoError(err)
	query := "query getUserAndTeam{user{id}team{id}}"
	_ = c.Execute(context.Background(), query, nil)
}
