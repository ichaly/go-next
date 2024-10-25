WITH "sys_area"
         AS (INSERT INTO "public"."sys_area" ("name", "pid") SELECT '巴黎' :: text, '1' :: bigint RETURNING "public"."sys_area".*)
SELECT jsonb_build_object('sysArea', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_area_0"."id" AS "id"
                                              FROM (SELECT "sys_area"."id" FROM "sys_area" LIMIT 20) AS "sys_area_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true