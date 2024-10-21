SELECT jsonb_build_object('sysArea', __sj_0.json, 'sysArea_cursor', __sj_0.__cursor) AS __root
FROM ((SELECT true)) AS __root_x
         LEFT OUTER JOIN LATERAL (SELECT COALESCE(jsonb_agg(__sj_0.json), '[]')                  AS json,
                                         CONCAT('gj/6715ffc4:', CONCAT_WS(',', 0, MAX(__cur_0))) as __cursor
                                  FROM (SELECT to_jsonb(__sr_0.*) - '__cur_0' AS json, __cur_0
                                        FROM (SELECT "sys_area_0"."id"                     AS "id",
                                                     "sys_area_0"."name"                   AS "name",
                                                     LAST_VALUE("sys_area_0"."id") OVER () AS __cur_0
                                              FROM (WITH __cur
                                                             AS (SELECT a[2] :: bigint AS "id" FROM STRING_TO_ARRAY('0,5', ',') AS a)
                                                    SELECT "sys_area"."id", "sys_area"."name"
                                                    FROM "public"."sys_area" AS "sys_area",
                                                         __cur
                                                    WHERE ((("__cur"."id") IS NULL) OR (("sys_area"."id") < ("__cur"."id")))
                                                    ORDER BY "sys_area"."id" DESC
                                                    LIMIT 5) AS "sys_area_0") AS "__sr_0") AS "__sj_0") AS "__sj_0"
                         ON true;