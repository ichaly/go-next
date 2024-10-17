package core

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type _ExecutorWhereRelationSuite struct {
	_ExecutorSuite
	e *Executor
}

func TestExecutorWhereRelation(t *testing.T) {
	suite.Run(t, new(_ExecutorWhereRelationSuite))
}

func (my *_ExecutorWhereRelationSuite) SetupSuite() {
	var err error
	my._ExecutorSuite.SetupSuite()

	my.e, err = NewExecutor(my.c, my.s)
	my.Require().NoError(err)
}

func (my *_ExecutorWhereRelationSuite) TestExecutorWhereRelation() {
	input := `query{userList{key:id area(where:{and:[{name:{eq:"北京"}},{id:{eq:1}}]}){key:id}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name"FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE (("sys_area"."id" = 1)) LIMIT 20 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}
