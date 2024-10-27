-- mutation{sysArea(insert:{name:"黎巴嫩" sysUser:{name:"管理员",connect:{id:{equals:1}}}}){id}}
WITH "sys_area" AS (INSERT INTO "public"."sys_area" ("name") SELECT '黎巴嫩' :: text RETURNING "public"."sys_area".*),
     "sys_user"
         AS ( UPDATE "public"."sys_user" SET "area_id" = "_x_sys_area"."id" FROM "sys_area" _x_sys_area WHERE (("sys_user"."id") = '1') RETURNING "public"."sys_user".*)
SELECT jsonb_build_object('sysArea', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_area_0"."id" AS "id"
                                              FROM (SELECT "sys_area"."id" FROM "sys_area" LIMIT 20) AS "sys_area_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;

-- mutation{sysArea(insert:{name:"黎巴嫩",sysUser:{name:"管理员"}}){id}}
WITH "sys_area" AS (INSERT INTO "public"."sys_area" ("name") SELECT '黎巴嫩' :: text RETURNING "public"."sys_area".*),
     "sys_user"
         AS (INSERT INTO "public"."sys_user" ("name", "area_id") SELECT '管理员' :: text, "sys_area"."id" FROM "sys_area" RETURNING "public"."sys_user".*)
SELECT jsonb_build_object('sysArea', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_area_0"."id" AS "id"
                                              FROM (SELECT "sys_area"."id" FROM "sys_area" LIMIT 20) AS "sys_area_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;

-- mutation{sysUser(insert:{name:"操作员",sysEdge:{userId:11,teamId:1}}){id}}
WITH "sys_user" AS (INSERT INTO "public"."sys_user" ("name") SELECT '操作员' :: text RETURNING "public"."sys_user".*),
     "sys_edge"
         AS (INSERT INTO "public"."sys_edge" ("team_id", "user_id") SELECT '1' :: bigint, "sys_user"."id" FROM "sys_user" RETURNING "public"."sys_edge".*)
SELECT jsonb_build_object('sysUser', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_user_0"."id" AS "id"
                                              FROM (SELECT "sys_user"."id" FROM "sys_user" LIMIT 20) AS "sys_user_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;

-- mutation{sysUser(insert:{name:"操作员" sysTeam:{name:"北京总部"}sysEdge:{userId:11,teamId:1}}){id}}
WITH "sys_team"
         AS (INSERT INTO "public"."sys_team" ("name") SELECT '北京总部' :: character varying(255) RETURNING "public"."sys_team".*),
     "sys_user"
         AS (INSERT INTO "public"."sys_user" ("name") SELECT '操作员' :: text, "sys_team"."id" FROM "sys_team" RETURNING "public"."sys_user".*),
     "sys_edge"
         AS (INSERT INTO "public"."sys_edge" ("user_id", "team_id") SELECT "sys_user"."id", "sys_team"."id"
                                                                    FROM "sys_user",
                                                                         "sys_team" RETURNING "public"."sys_edge".*)
SELECT jsonb_build_object('sysUser', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_user_0"."id" AS "id"
                                              FROM (SELECT "sys_user"."id" FROM "sys_user" LIMIT 20) AS "sys_user_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;

-- mutation{sysUser(insert:{name:"管理员1",sysArea:{name:"伊拉克",weight:1}}where:{id:{equals:1}}){id}}
WITH "sys_area"
         AS (INSERT INTO "public"."sys_area" ("name", "weight") SELECT '伊拉克' :: text, '1' :: integer RETURNING "public"."sys_area".*),
     "sys_user"
         AS (INSERT INTO "public"."sys_user" ("name", "area_id") SELECT '管理员1' :: text, "sys_area"."id" FROM "sys_area" RETURNING "public"."sys_user".*)
SELECT jsonb_build_object('sysUser', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_user_0"."id" AS "id"
                                              FROM (SELECT "sys_user"."id"
                                                    FROM "sys_user"
                                                    WHERE (("sys_user"."id") = '1')
                                                    LIMIT 20) AS "sys_user_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;

-- mutation{sysUser(insert:{name:"管理员1" sysArea:{name:"伊拉克",weight:1,sysUser:{name:"管理员2"}}}){id}}
WITH "sys_area"
         AS (INSERT INTO "public"."sys_area" ("name", "weight") SELECT '伊拉克' :: text, '1' :: integer RETURNING "public"."sys_area".*),
     sys_user_1
         AS (INSERT INTO "public"."sys_user" ("name", "area_id") SELECT '管理员2' :: text, "sys_area"."id" FROM "sys_area" RETURNING "public"."sys_user".*),
     sys_user_2
         AS (INSERT INTO "public"."sys_user" ("name", "area_id") SELECT '管理员1' :: text, "sys_area"."id" FROM "sys_area" RETURNING "public"."sys_user".*),
     "sys_user" AS (SELECT * FROM sys_user_1 UNION ALL SELECT * FROM sys_user_2)
SELECT jsonb_build_object('sysUser', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_user_0"."id" AS "id"
                                              FROM (SELECT "sys_user"."id" FROM "sys_user" LIMIT 20) AS "sys_user_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;

-- mutation{sysUser(insert:{name:"管理员1" sysArea:{name:"伊拉克" weight:1 sysUser:{name:"管理员2",sysArea:{name:"俄罗斯"}}}}){id}}
WITH sys_area_0
         AS (INSERT INTO "public"."sys_area" ("name", "weight") SELECT '伊拉克' :: text, '1' :: integer RETURNING "public"."sys_area".*),
     sys_area_1 AS (INSERT INTO "public"."sys_area" ("name") SELECT '俄罗斯' :: text RETURNING "public"."sys_area".*),
     sys_user_2
         AS (INSERT INTO "public"."sys_user" ("name", "area_id", "area_id") SELECT '管理员2' :: text, "sys_area"."id", "sys_area"."id"
                                                                            FROM sys_area_0 "sys_area",
                                                                                 sys_area_1 "sys_area" RETURNING "public"."sys_user".*),
     sys_user_3
         AS (INSERT INTO "public"."sys_user" ("name", "area_id") SELECT '管理员1' :: text, "sys_area"."id" FROM sys_area_0 "sys_area" RETURNING "public"."sys_user".*),
     "sys_area" AS (SELECT * FROM sys_area_0 UNION ALL SELECT * FROM sys_area_1),
     "sys_user" AS (SELECT * FROM sys_user_2 UNION ALL SELECT * FROM sys_user_3)
SELECT jsonb_build_object('sysUser', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_user_0"."id" AS "id"
                                              FROM (SELECT "sys_user"."id" FROM "sys_user" LIMIT 20) AS "sys_user_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true