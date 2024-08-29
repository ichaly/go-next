create table if not exists "user"
(
    id bigserial
        constraint user_pk
            primary key
);

create table if not exists "area"
(
    id bigserial
        constraint area_pk
            primary key
);

create table if not exists "item"
(
    id      bigserial
        constraint item_pk
            primary key,
    user_id bigint
        constraint item_user_id_fk
            references "user"
);

create table if not exists "team"
(
    id      bigserial
        constraint team_pk
            primary key,
    pid     bigint
        constraint team_team_id_fk
            references "team",
    area_id bigint
        constraint team_area_id_fk
            references "area"
);

create table if not exists "edge"
(
    user_id bigint not null
        constraint edge_user_id_fk
            references "user",
    team_id bigint not null
        constraint edge_team_id_fk
            references "team",
    constraint edge_pk
        primary key (user_id, team_id)
);

