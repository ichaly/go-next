-- {sysArea{id name sysArea(find:"parents"){id name}}}
SELECT jsonb_build_object('sysArea', __sj_0.json) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]') AS json
                                  FROM (SELECT to_jsonb(__sr_0.*) AS json
                                        FROM (SELECT "sys_area_0"."id"   AS "id",
                                                     "sys_area_0"."name" AS "name",
                                                     __sj_1.json         AS "sysArea"
                                              FROM (SELECT "sys_area"."id", "sys_area"."name"
                                                    FROM "public"."sys_area" AS "sys_area"
                                                    LIMIT 20) AS "sys_area_0"
                                                       LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_1.json), '[]') AS json
                                                                                FROM (SELECT to_jsonb(__sr_1.*) AS json
                                                                                      FROM (SELECT "sys_area_1"."id" AS "id", "sys_area_1"."name" AS "name"
                                                                                            FROM (WITH RECURSIVE "__rcte_sys_area"
                                                                                                                     AS ((SELECT "sys_area"."id", "sys_area"."name", "sys_area"."pid"
                                                                                                                          FROM "public"."sys_area" AS "sys_area"
                                                                                                                          WHERE ("sys_area"."id") = ("sys_area_0"."id")
                                                                                                                          LIMIT 1)
                                                                                                                         UNION ALL
                                                                                                                         SELECT "sys_area"."id", "sys_area"."name", "sys_area"."pid"
                                                                                                                         FROM "public"."sys_area" AS "sys_area",
                                                                                                                              "__rcte_sys_area"
                                                                                                                         WHERE ((("__rcte_sys_area"."pid") IS NOT NULL) AND
                                                                                                                                (("__rcte_sys_area"."pid") != ("__rcte_sys_area"."id")) AND
                                                                                                                                (("sys_area"."id") = ("__rcte_sys_area"."pid"))))
                                                                                                  SELECT "sys_area"."id" AS "id", "sys_area"."name" AS "name"
                                                                                                  FROM (SELECT * FROM "__rcte_sys_area" OFFSET 1) AS "sys_area"
                                                                                                  LIMIT 20) AS "sys_area_1") AS "__sr_1") AS "__sj_1") AS "__sj_1"
                                                                       ON true) AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true