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

create table ingredients
(
    name varchar(100) not null,
    id   bigserial
        constraint ingredients_pk
            primary key
);

alter table ingredients
    owner to postgres;

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

create table favorites
(
    "user" bigint not null
        constraint users_fk
            references users,
    recipe bigint not null
        constraint recipe_fk
            references recipes,
    constraint favorites_pk
        primary key ("user", recipe)
);

alter table favorites
    owner to postgres;

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