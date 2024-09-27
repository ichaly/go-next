-- ManyToOne
-- query{sysUser{key:id sysArea{key:id}}}
SELECT jsonb_build_object('sysUser', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_user_0"."id" AS "key", __sj_1.json AS "sysArea"
                                              FROM (SELECT "sys_user"."id", "sys_user"."area_id"
                                                    FROM "public"."sys_user" AS "sys_user"
                                                    LIMIT 20) AS "sys_user_0"
                                                       LEFT OUTER JOIN LATERAL (SELECT to_jsonb(__sr_1.*) AS json
                                                                                FROM (SELECT "sys_area_1"."id" AS "key"
                                                                                      FROM (SELECT "sys_area"."id"
                                                                                            FROM "public"."sys_area" AS "sys_area"
                                                                                            WHERE (("sys_area"."id") = ("sys_user_0"."area_id"))
                                                                                            LIMIT 1) AS "sys_area_1") AS "__sr_1") AS "__sj_1"
                                                                       ON true) AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;