package core

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type _ExecutorWhereSuite struct {
	_ExecutorSuite
	e *Executor
}

func TestExecutorWhere(t *testing.T) {
	suite.Run(t, new(_ExecutorWhereSuite))
}

func (my *_ExecutorWhereSuite) SetupSuite() {
	var err error
	my._ExecutorSuite.SetupSuite()

	my.e, err = NewExecutor(my.c, my.s)
	my.Require().NoError(err)
}

func (my *_ExecutorWhereSuite) TestExecutorWhereBase() {
	input := "query{areaList(where:{id:{eq:1}}){id name}}"
	expect := ""
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}

func (my *_ExecutorWhereSuite) TestExecutorWhereNot() {
	input := "query{areaList(where:{not:{id:{le:1}}}){id name}}"
	expect := ""
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}

func (my *_ExecutorWhereSuite) TestExecutorWhereAnd() {
	input := "query{areaList(where:{and:[{id:{ge:1}},{id:{le:10}}]}){id name}}"
	expect := ""
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}

func (my *_ExecutorWhereSuite) TestExecutorWhereOr() {
	input := "query{areaList(where:{or:[{id:{ge:10}},{id:{le:1}}]}){id name}}"
	expect := ""
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}
