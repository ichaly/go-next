package core

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"testing"
)

func (my *_ExecutorSuite) test(key string) {
	data := map[string]struct {
		Input  string
		Expect string
	}{
		"Base": {
			Input:  `query{areaList{id name}}`,
			Expect: `SELECT jsonb_build_object('areaList', __sj_4294967303.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967303.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967303.*) AS json FROM (  SELECT "sys_area_4294967303"."id" AS "id","sys_area_4294967303"."name" AS "name"FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" LIMIT 20 ) AS"sys_area_4294967303") AS __sr_4294967303) AS __sj_4294967303) AS __sj_4294967303 ON true`,
		},
		"One2Many": {
			Input:  `query{areaList{key:id userList{key:id}}}`,
			Expect: `SELECT jsonb_build_object('areaList', __sj_4294967303.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967303.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967303.*) AS json FROM (  SELECT "sys_area_4294967303"."id" AS "key","__sj_4294967319"."json" AS "userList"FROM ( SELECT "sys_area"."id" FROM "sys_area" LIMIT 20 ) AS"sys_area_4294967303" LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967319.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967319.*) AS json FROM (  SELECT "sys_user_4294967319"."id" AS "key"FROM ( SELECT "sys_user"."id" FROM "sys_user" WHERE ("sys_user"."area_id" = "sys_area_4294967303"."id") LIMIT 20 ) AS"sys_user_4294967319") AS __sr_4294967319) AS __sj_4294967319) AS __sj_4294967319 ON true) AS __sr_4294967303) AS __sj_4294967303) AS __sj_4294967303 ON true`,
		},
		"Many2One": {
			Input:  `query{userList{key:id area{key:id}}}`,
			Expect: `SELECT jsonb_build_object('userList', __sj_4294967303.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967303.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967303.*) AS json FROM (  SELECT "sys_user_4294967303"."id" AS "key","__sj_4294967319"."json" AS "area"FROM ( SELECT "sys_user"."id","sys_user"."area_id" FROM "sys_user" LIMIT 20 ) AS"sys_user_4294967303" LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967319.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967319.*) AS json FROM (  SELECT "sys_area_4294967319"."id" AS "key"FROM ( SELECT "sys_area"."id" FROM "sys_area" WHERE ("sys_area"."id" = "sys_user_4294967303"."area_id") LIMIT 20 ) AS"sys_area_4294967319") AS __sr_4294967319) AS __sj_4294967319) AS __sj_4294967319 ON true) AS __sr_4294967303) AS __sj_4294967303) AS __sj_4294967303 ON true`,
		},
		"Many2Many": {
			Input:  `query{userList{key:id teamList{key:id}}}`,
			Expect: `SELECT jsonb_build_object('userList', __sj_4294967303.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967303.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967303.*) AS json FROM (  SELECT "sys_user_4294967303"."id" AS "key","__sj_4294967319"."json" AS "teamList"FROM ( SELECT "sys_user"."id" FROM "sys_user" LIMIT 20 ) AS"sys_user_4294967303" LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967319.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967319.*) AS json FROM (  SELECT "sys_team_4294967319"."id" AS "key"FROM ( SELECT "sys_team"."id" FROM "sys_team" INNER JOIN sys_edge ON (("sys_edge" . "user_id" = "sys_user_4294967303" . "id")) WHERE ("sys_team"."id" = "sys_edge"."team_id") LIMIT 20 ) AS"sys_team_4294967319") AS __sr_4294967319) AS __sj_4294967319) AS __sj_4294967319 ON true) AS __sr_4294967303) AS __sj_4294967303) AS __sj_4294967303 ON true`,
		},
		"RecursiveChildren": {
			Input:  `query{areaList{id name children{id name}}}`,
			Expect: `SELECT jsonb_build_object('areaList', __sj_4294967303.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967303.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967303.*) AS json FROM (  SELECT "sys_area_4294967303"."id" AS "id","sys_area_4294967303"."name" AS "name","__sj_4294967320"."json" AS "children"FROM ( SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" LIMIT 20 ) AS"sys_area_4294967303" LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4294967320.json), '[]') AS json FROM (SELECT to_jsonb(__sr_4294967320.*) AS json FROM (  SELECT "sys_area_4294967320"."id" AS "id","sys_area_4294967320"."name" AS "name"FROM ( WITH RECURSIVE "__rcte_sys_area" AS ((SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" LIMIT 1 ) UNION ALL  SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name","sys_area"."pid" FROM "sys_area" , "__rcte_sys_area") SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name" FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS  "sys_area" LIMIT 20 ) AS"sys_area_4294967320") AS __sr_4294967320) AS __sj_4294967320) AS __sj_4294967320 ON true) AS __sr_4294967303) AS __sj_4294967303) AS __sj_4294967303 ON true`,
		},
	}
	item := data[key]
	input, expect := item.Input, item.Expect
	e, err := NewExecutor(my.c, my.s)
	my.Require().NoError(err)
	r := e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
}

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
	my.test("Base")
}

func (my *_ExecutorSuite) TestExecutorOne2Many() {
	my.test("One2Many")
}

func (my *_ExecutorSuite) TestExecutorMany2One() {
	my.test("Many2One")
}

func (my *_ExecutorSuite) TestExecutorMany2Many() {
	my.test("Many2Many")
}

func (my *_ExecutorSuite) TestExecutorRecursiveChildren() {
	my.test("RecursiveChildren")
}
