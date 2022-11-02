create table recipes
(
    title       varchar(100) not null,
    description text,
    owner       bigint
        constraint owner_fk
            references users,
    id          bigserial
        constraint recipes_pk
            primary key
);

alter table recipes
    owner to postgres;

