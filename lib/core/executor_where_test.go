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
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name"FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE (("sys_area"."id" = 1)) LIMIT 20 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}

func (my *_ExecutorWhereSuite) TestExecutorWhereNot() {
	input := "query{areaList(where:{not:{id:{le:1}}}){id name}}"
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name"FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE (NOT (("sys_area"."id" <= 1))) LIMIT 20 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}

func (my *_ExecutorWhereSuite) TestExecutorWhereAnd() {
	input := "query{areaList(where:{and:[{id:{ge:1}},{id:{le:10}}]}){id name}}"
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name"FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE ((("sys_area"."id" >= 1) AND ("sys_area"."id" <= 10))) LIMIT 20 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}

func (my *_ExecutorWhereSuite) TestExecutorWhereOr() {
	input := "query{areaList(where:{or:[{id:{ge:10}},{id:{le:1}}]}){id name}}"
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name"FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE ((("sys_area"."id" >= 10) OR ("sys_area"."id" <= 1))) LIMIT 20 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}
