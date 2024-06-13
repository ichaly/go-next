SELECT c.table_name,
       c.column_name,
       c.data_type,
       CAST(c.is_nullable AS BOOLEAN),
       pk.is_primary,
       fk.is_foreign,
       obj_description(c.table_name::regclass)                     table_description,
       col_description(c.table_name::regclass, c.ordinal_position) column_description
FROM information_schema.columns c
         LEFT JOIN (SELECT conrelid::regclass::text AS table_name,
                           attnum,
                           TRUE                     AS is_primary
                    FROM pg_constraint
                             JOIN pg_attribute ON conrelid = attrelid AND attnum = ANY (conkey)
                    WHERE contype = 'p') pk ON c.table_name = pk.table_name AND c.ordinal_position = pk.attnum
         LEFT JOIN (SELECT conrelid::regclass::text AS table_name,
                           attnum,
                           TRUE                     AS is_foreign
                    FROM pg_constraint
                             JOIN pg_attribute ON conrelid = attrelid AND attnum = ANY (conkey)
                    WHERE contype = 'f') fk ON c.table_name = fk.table_name AND c.ordinal_position = fk.attnum
WHERE c.table_schema = 'public'
order by c.table_name, c.ordinal_position