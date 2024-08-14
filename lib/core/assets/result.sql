SELECT jsonb_build_object('goods', __sj_0.json, 'teams', __sj_2.json, 'users', __sj_4.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_4.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_4.*) AS json
                                        FROM (SELECT "users_4"."id" AS "id",
                                                     __sj_5.json    AS "goods",
                                                     __sj_6.json    AS "teams"
                                              FROM (SELECT "users"."id" FROM "public"."users" AS "users" LIMIT 20) AS "users_4"
                                                       LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_6.json), '[]') AS json
                                                                                FROM (SELECT to_jsonb(__sr_6.*) AS json
                                                                                      FROM (SELECT "teams_6"."id" AS "id"
                                                                                            FROM (SELECT "teams"."id"
                                                                                                  FROM "public"."teams" AS "teams"
                                                                                                           INNER JOIN sys_users_teams
                                                                                                                      ON (((("sys_users_teams"."user_id") = ("users_4"."id"))))
                                                                                                  WHERE (("teams"."id") = ("sys_users_teams"."team_id"))
                                                                                                  LIMIT 20) AS "teams_6") AS "__sr_6") AS "__sj_6") AS "__sj_6"
                                                                       ON true
                                                       LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_5.json), '[]') AS json
                                                                                FROM (SELECT to_jsonb(__sr_5.*) AS json
                                                                                      FROM (SELECT "goods_5"."id" AS "id"
                                                                                            FROM (SELECT "goods"."id"
                                                                                                  FROM "public"."goods" AS "goods"
                                                                                                  WHERE (("goods"."owner_id") = ("users_4"."id"))
                                                                                                  LIMIT 20) AS "goods_5") AS "__sr_5") AS "__sj_5") AS "__sj_5"
                                                                       ON true) AS "__sr_4") AS "__sj_4") AS "__sj_4"
                         ON true
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_2.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_2.*) AS json
                                        FROM (SELECT "teams_2"."id" AS "id", __sj_3.json AS "user"
                                              FROM (SELECT "teams"."id" FROM "public"."teams" AS "teams" LIMIT 20) AS "teams_2"
                                                       LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_3.json), '[]') AS json
                                                                                FROM (SELECT to_jsonb(__sr_3.*) AS json
                                                                                      FROM (SELECT "users_3"."id" AS "id"
                                                                                            FROM (SELECT "users"."id"
                                                                                                  FROM "public"."users" AS "users"
                                                                                                           INNER JOIN sys_users_teams
                                                                                                                      ON (((("sys_users_teams"."team_id") = ("teams_2"."id"))))
                                                                                                  WHERE (("users"."id") = ("sys_users_teams"."user_id"))
                                                                                                  LIMIT 20) AS "users_3") AS "__sr_3") AS "__sj_3") AS "__sj_3"
                                                                       ON true) AS "__sr_2") AS "__sj_2") AS "__sj_2"
                         ON true
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT __sj_1.json AS "owner"
                                              FROM (SELECT "goods"."owner_id" FROM "public"."goods" AS "goods" LIMIT 20) AS "goods_0"
                                                       LEFT OUTER JOIN LATERAL (SELECT to_jsonb(__sr_1.*) AS json
                                                                                FROM (SELECT "users_1"."id" AS "id"
                                                                                      FROM (SELECT "users"."id"
                                                                                            FROM "public"."users" AS "users"
                                                                                            WHERE (("users"."id") = ("goods_0"."owner_id"))
                                                                                            LIMIT 1) AS "users_1") AS "__sr_1") AS "__sj_1"
                                                                       ON true) AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;
