create table if not exists sys_user
(
    id bigserial
        constraint sys_user_pk
            primary key
);

create table if not exists sys_area
(
    id bigserial
        constraint sys_area_pk
            primary key
);

create table if not exists sys_item
(
    id      bigserial
        constraint sys_item_pk
            primary key,
    user_id bigint
        constraint sys_item_sys_user_id_fk
            references sys_user
);

create table if not exists sys_team
(
    id      bigserial
        constraint sys_team_pk
            primary key,
    pid     bigint
        constraint sys_team_sys_team_id_fk
            references sys_team,
    area_id bigint
        constraint sys_team_sys_area_id_fk
            references sys_area
);

create table if not exists sys_edge
(
    user_id bigint not null
        constraint sys_edge_sys_user_id_fk
            references sys_user,
    team_id bigint not null
        constraint sys_edge_sys_team_id_fk
            references sys_team,
    constraint sys_edge_pk
        primary key (user_id, team_id)
);

