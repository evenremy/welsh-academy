create table stages
(
    recipe      bigint  not null
        constraint recipe_fk
            references recipes,
    "order"     integer not null,
    description text,
    constraint stages_pk
        primary key (recipe, "order")
);

alter table stages
    owner to postgres;

