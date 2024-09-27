-- ManyToMany
-- query{sysUser{key:id team:sysTeam{key:id}}}
SELECT jsonb_build_object('sysUser', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_user_0"."id" AS "key", __sj_1.json AS "team"
                                              FROM (SELECT "sys_user"."id"
                                                    FROM "public"."sys_user" AS "sys_user"
                                                    LIMIT 20) AS "sys_user_0"
                                                       LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json
                                                                                FROM (SELECT to_jsonb(__sr_1.*) AS json
                                                                                      FROM (SELECT "sys_team_1"."id" AS "key"
                                                                                            FROM (SELECT "sys_team"."id"
                                                                                                  FROM "public"."sys_team" AS "sys_team"
                                                                                                           INNER JOIN sys_edge ON (((("sys_edge"."user_id") = ("sys_user_0"."id"))))
                                                                                                  WHERE (("sys_team"."id") = ("sys_edge"."team_id"))
                                                                                                  LIMIT 20) AS "sys_team_1") AS "__sr_1") AS "__sj_1") AS "__sj_1"
                                                                       ON true) AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;