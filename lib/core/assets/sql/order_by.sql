-- OneToMany
-- query{sysArea(orderBy:{name:desc,id:asc}){key:id sysUser(orderBy:{id:desc}){key:id}}}
SELECT jsonb_build_object('sysArea', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_area_0"."id" AS "key", __sj_1.json AS "sysUser"
                                              FROM (SELECT "sys_area"."id", "sys_area"."name"
                                                    FROM "public"."sys_area" AS "sys_area"
                                                    ORDER BY "sys_area"."name" DESC, "sys_area"."id" ASC
                                                    LIMIT 20) AS "sys_area_0"
                                                       LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json
                                                                                FROM (SELECT to_jsonb(__sr_1.*) AS json
                                                                                      FROM (SELECT "sys_user_1"."id" AS "key"
                                                                                            FROM (SELECT "sys_user"."id"
                                                                                                  FROM "public"."sys_user" AS "sys_user"
                                                                                                  WHERE (("sys_user"."area_id") = ("sys_area_0"."id"))
                                                                                                  ORDER BY "sys_user"."id" DESC
                                                                                                  LIMIT 20) AS "sys_user_1") AS "__sr_1") AS "__sj_1") AS "__sj_1"
                                                                       ON true) AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;