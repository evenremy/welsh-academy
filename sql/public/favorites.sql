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

