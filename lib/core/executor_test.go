package core

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"testing"
)

type _ExecutorSuite struct {
	_CompilerSuite
	c *Compiler
}

func TestExecutor(t *testing.T) {
	suite.Run(t, new(_ExecutorSuite))
}

func (my *_ExecutorSuite) SetupSuite() {
	my._CompilerSuite.SetupSuite()
	var err error

	s := &ast.Source{Name: "metadata"}
	s.Input, err = my.m.MarshalSchema()
	my.Require().NoError(err)

	my.s, err = gqlparser.LoadSchema(s)
	my.Require().NoError(err)
	my.c = NewCompiler(my.m, my.s)
}

func (my *_ExecutorSuite) TestExecutor() {
	e, err := NewExecutor(my.c, nil)
	my.Require().NoError(err)
	r := e.Execute(context.Background(), `query getUserAndTeam{user{id}team{id}}`, nil)
	my.T().Log(r)
}
