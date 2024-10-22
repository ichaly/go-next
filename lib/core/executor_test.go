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
	e *Executor
}

func TestExecutor(t *testing.T) {
	suite.Run(t, new(_ExecutorSuite))
}

func (my *_ExecutorSuite) doCase(input, expect string) {
	r := my.e.Execute(context.Background(), input, nil)
	my.T().Log(r.Sql)
	my.Require().Equal(expect, r.Sql)
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

	my.e, err = NewExecutor(my.c, my.s)
	my.Require().NoError(err)
}

func (my *_ExecutorSuite) TestExecutorBase() {
	input := `query{areaList{id name}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorOne2Many() {
	input := `query{areaList{key:id userList{key:id areaId}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "key","__sj_1"."json" AS "userList" FROM ( SELECT "sys_area"."id" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_user_1"."id" AS "key","sys_user_1"."area_id" AS "areaId" FROM ( SELECT "sys_user"."id","sys_user"."area_id" FROM "sys_user" WHERE ((("sys_user"."area_id") = "sys_area_0"."id")) LIMIT 10 ) AS"sys_user_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorMany2One() {
	input := `query{userList{key:id area{key:id}}}`
	expect := `SELECT jsonb_build_object('userList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_user_0"."id" AS "key","__sj_1"."json" AS "area" FROM ( SELECT "sys_user"."id","sys_user"."area_id" FROM "sys_user" LIMIT 10 ) AS"sys_user_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_area_1"."id" AS "key" FROM ( SELECT "sys_area"."id" FROM "sys_area" WHERE ((("sys_area"."id") = "sys_user_0"."area_id")) LIMIT 10 ) AS"sys_area_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorMany2Many() {
	input := `query{userList{key:id teamList{key:id}}}`
	expect := `SELECT jsonb_build_object('userList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_user_0"."id" AS "key","__sj_1"."json" AS "teamList" FROM ( SELECT "sys_user"."id" FROM "sys_user" LIMIT 10 ) AS"sys_user_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_team_1"."id" AS "key" FROM ( SELECT "sys_team"."id" FROM "sys_team" INNER JOIN sys_edge ON (("sys_edge" . "user_id" = "sys_user_0" . "id")) WHERE ((("sys_team"."id") = "sys_edge"."team_id")) LIMIT 10 ) AS"sys_team_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorRecursiveParents() {
	input := `query{areaList{id name parents{id name}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name","__sj_1"."json" AS "parents" FROM ( SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_area_1"."id" AS "id","sys_area_1"."name" AS "name" FROM ( WITH RECURSIVE "__rcte_sys_area" AS ((SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" WHERE "sys_area".id = "sys_area_0".id LIMIT 1 ) UNION ALL  SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" , "__rcte_sys_area" WHERE (((("__rcte_sys_area"."pid") IS NOT NULL) AND (("__rcte_sys_area"."pid") != "__rcte_sys_area"."id") AND (("__rcte_sys_area"."pid") = "sys_area"."id")))) SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name" FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS  "sys_area" LIMIT 10 ) AS"sys_area_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorRecursiveChildren() {
	input := `query{areaList{id name children{id name}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name","__sj_1"."json" AS "children" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_area_1"."id" AS "id","sys_area_1"."name" AS "name" FROM ( WITH RECURSIVE "__rcte_sys_area" AS ((SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" WHERE "sys_area".id = "sys_area_0".id LIMIT 1 ) UNION ALL  SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" , "__rcte_sys_area" WHERE (((("sys_area"."pid") IS NOT NULL) AND (("sys_area"."pid") != "sys_area"."id") AND (("sys_area"."pid") = "__rcte_sys_area"."id")))) SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name" FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS  "sys_area" LIMIT 10 ) AS"sys_area_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorRecursiveParentsAndChildren() {
	input := `query{areaList{id name parents{id name}children{id name}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name","__sj_1"."json" AS "parents","__sj_2"."json" AS "children" FROM ( SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_area_1"."id" AS "id","sys_area_1"."name" AS "name" FROM ( WITH RECURSIVE "__rcte_sys_area" AS ((SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" WHERE "sys_area".id = "sys_area_0".id LIMIT 1 ) UNION ALL  SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" , "__rcte_sys_area" WHERE (((("__rcte_sys_area"."pid") IS NOT NULL) AND (("__rcte_sys_area"."pid") != "__rcte_sys_area"."id") AND (("__rcte_sys_area"."pid") = "sys_area"."id")))) SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name" FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS  "sys_area" LIMIT 10 ) AS"sys_area_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_2.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_2.*) AS json FROM (  SELECT "sys_area_2"."id" AS "id","sys_area_2"."name" AS "name" FROM ( WITH RECURSIVE "__rcte_sys_area" AS ((SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" WHERE "sys_area".id = "sys_area_0".id LIMIT 1 ) UNION ALL  SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" , "__rcte_sys_area" WHERE (((("sys_area"."pid") IS NOT NULL) AND (("sys_area"."pid") != "sys_area"."id") AND (("sys_area"."pid") = "__rcte_sys_area"."id")))) SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name" FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS  "sys_area" LIMIT 10 ) AS"sys_area_2" ) AS "__sr_2" ) AS "__sj_2" ) AS "__sj_2" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereBase() {
	input := "query{areaList(where:{id:{eq:1}}){id name}}"
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE (("sys_area"."id" = 1)) LIMIT 10 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereNot() {
	input := "query{areaList(where:{not:{id:{le:1}}}){id name}}"
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE (NOT (("sys_area"."id" <= 1))) LIMIT 10 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereAnd() {
	input := "query{areaList(where:{and:[{id:{ge:1}},{id:{le:10}}]}){id name}}"
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE ((("sys_area"."id" >= 1) AND ("sys_area"."id" <= 10))) LIMIT 10 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereOr() {
	input := "query{areaList(where:{or:[{id:{ge:10}},{id:{le:1}}]}){id name}}"
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" WHERE ((("sys_area"."id" >= 10) OR ("sys_area"."id" <= 1))) LIMIT 10 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereOne2Many() {
	input := `query{areaList{key:id userList(where:{and:[{id:{ge:1}},{id:{le:3}}]}){key:id}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "key","__sj_1"."json" AS "userList" FROM ( SELECT "sys_area"."id" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_user_1"."id" AS "key" FROM ( SELECT "sys_user"."id" FROM "sys_user" WHERE (((("sys_user"."id" >= 1) AND ("sys_user"."id" <= 3)) AND (("sys_user"."area_id") = "sys_area_0"."id"))) LIMIT 10 ) AS"sys_user_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereMany2One() {
	input := `query{userList{key:id area(where:{and:[{name:{eq:"北京"}},{id:{eq:1}}]}){key:id}}}`
	expect := `SELECT jsonb_build_object('userList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_user_0"."id" AS "key","__sj_1"."json" AS "area" FROM ( SELECT "sys_user"."id","sys_user"."area_id" FROM "sys_user" LIMIT 10 ) AS"sys_user_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_area_1"."id" AS "key" FROM ( SELECT "sys_area"."id" FROM "sys_area" WHERE (((("sys_area"."name" = '北京') AND ("sys_area"."id" = 1)) AND (("sys_area"."id") = "sys_user_0"."area_id"))) LIMIT 10 ) AS"sys_area_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereMany2Many() {
	input := `query{userList{key:id teamList(where:{and:[{id:{ge:1}},{id:{le:3}}]}){key:id}}}`
	expect := `SELECT jsonb_build_object('userList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_user_0"."id" AS "key","__sj_1"."json" AS "teamList" FROM ( SELECT "sys_user"."id" FROM "sys_user" LIMIT 10 ) AS"sys_user_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_team_1"."id" AS "key" FROM ( SELECT "sys_team"."id" FROM "sys_team" INNER JOIN sys_edge ON (("sys_edge" . "user_id" = "sys_user_0" . "id")) WHERE (((("sys_team"."id" >= 1) AND ("sys_team"."id" <= 3)) AND (("sys_team"."id") = "sys_edge"."team_id"))) LIMIT 10 ) AS"sys_team_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereRecursiveParents() {
	input := `query{areaList{id name parents(where:{and:[{id:{ge:1}},{id:{le:3}}]}){id name}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name","__sj_1"."json" AS "parents" FROM ( SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_area_1"."id" AS "id","sys_area_1"."name" AS "name" FROM ( WITH RECURSIVE "__rcte_sys_area" AS ((SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" WHERE "sys_area".id = "sys_area_0".id LIMIT 1 ) UNION ALL  SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" , "__rcte_sys_area" WHERE (((("sys_area"."id" >= 1) AND ("sys_area"."id" <= 3)) AND ((("__rcte_sys_area"."pid") IS NOT NULL) AND (("__rcte_sys_area"."pid") != "__rcte_sys_area"."id") AND (("__rcte_sys_area"."pid") = "sys_area"."id"))))) SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name" FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS  "sys_area" LIMIT 10 ) AS"sys_area_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorWhereRecursiveParentsAndChildren() {
	input := `query{areaList{id name parents(where:{and:[{id:{ge:1}},{id:{le:3}}]}){id name}children(where:{and:[{id:{ge:1}},{id:{le:3}}]}){id name}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name","__sj_1"."json" AS "parents","__sj_2"."json" AS "children" FROM ( SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_area_1"."id" AS "id","sys_area_1"."name" AS "name" FROM ( WITH RECURSIVE "__rcte_sys_area" AS ((SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" WHERE "sys_area".id = "sys_area_0".id LIMIT 1 ) UNION ALL  SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" , "__rcte_sys_area" WHERE (((("sys_area"."id" >= 1) AND ("sys_area"."id" <= 3)) AND ((("__rcte_sys_area"."pid") IS NOT NULL) AND (("__rcte_sys_area"."pid") != "__rcte_sys_area"."id") AND (("__rcte_sys_area"."pid") = "sys_area"."id"))))) SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name" FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS  "sys_area" LIMIT 10 ) AS"sys_area_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_2.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_2.*) AS json FROM (  SELECT "sys_area_2"."id" AS "id","sys_area_2"."name" AS "name" FROM ( WITH RECURSIVE "__rcte_sys_area" AS ((SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" WHERE "sys_area".id = "sys_area_0".id LIMIT 1 ) UNION ALL  SELECT "sys_area"."id","sys_area"."name","sys_area"."pid" FROM "sys_area" , "__rcte_sys_area" WHERE (((("sys_area"."id" >= 1) AND ("sys_area"."id" <= 3)) AND ((("sys_area"."pid") IS NOT NULL) AND (("sys_area"."pid") != "sys_area"."id") AND (("sys_area"."pid") = "__rcte_sys_area"."id"))))) SELECT "sys_area"."id" AS "id","sys_area"."name" AS "name" FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS  "sys_area" LIMIT 10 ) AS"sys_area_2" ) AS "__sr_2" ) AS "__sj_2" ) AS "__sj_2" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorBaseLimit() {
	input := `query{userList(limit:2){id}}`
	expect := `SELECT jsonb_build_object('userList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_user_0"."id" AS "id" FROM ( SELECT "sys_user"."id" FROM "sys_user" LIMIT 2 ) AS"sys_user_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorDistinct() {
	input := `query{areaList(distinct:["id","name"]){id name}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT DISTINCT ON ("sys_area_0"."id", "sys_area_0"."name") "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" LIMIT 10 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorDistinctMany2Many() {
	input := `query{userList{key:id teamList(distinct:["id"]){key:id}}}`
	expect := `SELECT jsonb_build_object('userList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_user_0"."id" AS "key","__sj_1"."json" AS "teamList" FROM ( SELECT "sys_user"."id" FROM "sys_user" LIMIT 10 ) AS"sys_user_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT DISTINCT ON ("sys_team_1"."id") "sys_team_1"."id" AS "key" FROM ( SELECT "sys_team"."id" FROM "sys_team" INNER JOIN sys_edge ON (("sys_edge" . "user_id" = "sys_user_0" . "id")) WHERE ((("sys_team"."id") = "sys_edge"."team_id")) LIMIT 10 ) AS"sys_team_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorOffset() {
	input := `query{areaList(offset:50){id name}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area" LIMIT 10 OFFSET 50 ) AS"sys_area_0" ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}

func (my *_ExecutorSuite) TestExecutorSortOneToMany() {
	input := `query{areaList(sort:{name:DESC_NULLS_LAST,id:ASC}){id name userList(sort:{id:ASC}){id areaId}}}`
	expect := `SELECT jsonb_build_object('areaList', __sj_0.json) AS __root FROM (SELECT true) AS __root_x LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_0.*) AS json FROM (  SELECT "sys_area_0"."id" AS "id","sys_area_0"."name" AS "name","__sj_1"."json" AS "userList" FROM ( SELECT "sys_area"."id","sys_area"."name" FROM "sys_area"  ORDER BY  "sys_area"."name" DESC NULLS LAST, "sys_area"."id" ASC LIMIT 10 ) AS"sys_area_0" LEFT OUTER JOIN LATERAL ( SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json FROM (  SELECT to_jsonb(__sr_1.*) AS json FROM (  SELECT "sys_user_1"."id" AS "id","sys_user_1"."area_id" AS "areaId" FROM ( SELECT "sys_user"."id","sys_user"."area_id" FROM "sys_user" WHERE ((("sys_user"."area_id") = "sys_area_0"."id"))  ORDER BY  "sys_user"."id" ASC LIMIT 10 ) AS"sys_user_1" ) AS "__sr_1" ) AS "__sj_1" ) AS "__sj_1" ON true  ) AS "__sr_0" ) AS "__sj_0" ) AS "__sj_0" ON true`
	my.doCase(input, expect)
}
