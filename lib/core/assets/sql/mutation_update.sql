-- mutation{sysUser(update:{name:"管理员",where:{id:{equals:1}}}){id}}
WITH "sys_user"
         AS (UPDATE "public"."sys_user" SET ("name") = (SELECT '管理员' :: text) WHERE (("sys_user"."id") = '1') RETURNING "public"."sys_user".*)
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

-- mutation{sysUser(update:{name:"管理员",sysArea:{name:"国外1",weight:1}}where:{id:{equals:1}}){id}}
WITH "sys_user"
         AS (UPDATE "public"."sys_user" SET ("name") = (SELECT '管理员' :: text) WHERE (("sys_user"."id") = '1') RETURNING "public"."sys_user".*),
     "sys_area" AS (UPDATE "public"."sys_area" SET ("weight", "name", "id") =
             (SELECT '1' :: integer, '国外1' :: text, "_x_sys_user"."area_id"
              FROM "sys_user" _x_sys_user) FROM "sys_user" _x_sys_user WHERE (("sys_area"."id") = ("_x_sys_user"."area_id")) RETURNING "public"."sys_area".*)
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


WITH "sys_user"
         AS (UPDATE "public"."sys_user" SET ("name") = (SELECT '管理员' :: text) WHERE (("sys_user"."id") = '1') RETURNING "public"."sys_user".*)
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


WITH "sys_user"
         AS (WITH "sys_area" AS (SELECT "id" FROM "sys_area" WHERE (("sys_area"."id") = '2') LIMIT 1) UPDATE "public"."sys_user" SET ("name") = (SELECT '管理员' :: text FROM "sys_area" _x_sys_area) WHERE (("sys_user"."id") = '1') RETURNING "public"."sys_user".*)
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