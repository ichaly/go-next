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
	s.Input, err = my.m.Marshal()
	my.Require().NoError(err)

	my.s, err = gqlparser.LoadSchema(s)
	my.Require().NoError(err)
	my.c = NewCompiler(my.m, my.s)
}

func (my *_ExecutorSuite) TestExecutorBase() {
	e, err := NewExecutor(my.c, my.s)
	my.Require().NoError(err)
	r := e.Execute(context.Background(), `query{areaList{id name}}`, nil)
	my.T().Log(r)
}

func (my *_ExecutorSuite) TestExecutorRecursive() {
	e, err := NewExecutor(my.c, my.s)
	my.Require().NoError(err)
	r := e.Execute(context.Background(), `query{areaList{id name children{id name}}}`, nil)
	my.T().Log(r)
}

func (my *_ExecutorSuite) TestExecutorOne2Many() {
	e, err := NewExecutor(my.c, my.s)
	my.Require().NoError(err)
	r := e.Execute(context.Background(), `query{areaList{key:id userList{key:id}}}`, nil)
	my.T().Log(r)
}

func (my *_ExecutorSuite) TestExecutorMany2One() {
	e, err := NewExecutor(my.c, my.s)
	my.Require().NoError(err)
	r := e.Execute(context.Background(), `query{userList{key:id area{key:id}}}`, nil)
	my.T().Log(r)
}

func (my *_ExecutorSuite) TestExecutorMany2Many() {
	e, err := NewExecutor(my.c, my.s)
	my.Require().NoError(err)
	r := e.Execute(context.Background(), `query{userList{key:id teamList{key:id}}}`, nil)
	my.T().Log(r)
}
