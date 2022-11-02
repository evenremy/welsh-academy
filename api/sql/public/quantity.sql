create table quantity
(
    recipe     bigint not null
        constraint recipe_fk
            references recipes,
    ingredient bigint not null
        constraint ingredient_fk
            references ingredients,
    constraint quantity_pk
        primary key (recipe, ingredient)
);

alter table quantity
    owner to postgres;

