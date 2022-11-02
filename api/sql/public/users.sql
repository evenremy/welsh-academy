create table users
(
    username varchar(100)          not null,
    is_admin boolean default false not null,
    id       bigserial
        constraint users_pk
            primary key
);

alter table users
    owner to postgres;

create unique index username_index
    on users (username);

