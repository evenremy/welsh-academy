create table ingredients
(
    name varchar(100) not null,
    id   bigserial
        constraint ingredients_pk
            primary key
);

alter table ingredients
    owner to postgres;

