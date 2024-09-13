create table if not exists "sys_user"
(
    id bigserial
        constraint user_pk
            primary key
);

create table if not exists "sys_area"
(
    id bigserial
        constraint area_pk
            primary key
);

create table if not exists "sys_item"
(
    id      bigserial
        constraint item_pk
            primary key,
    user_id bigint
        constraint item_user_id_fk
            references "sys_user"
);

create table if not exists "sys_team"
(
    id      bigserial
        constraint team_pk
            primary key,
    pid     bigint
        constraint team_team_id_fk
            references "sys_team",
    area_id bigint
        constraint team_area_id_fk
            references "sys_area"
);

create table if not exists "sys_edge"
(
    user_id bigint not null
        constraint edge_user_id_fk
            references "sys_user",
    team_id bigint not null
        constraint edge_team_id_fk
            references "sys_team",
    constraint edge_pk
        primary key (user_id, team_id)
);

