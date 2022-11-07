alter table stages
    drop constraint recipe_fk,
    add constraint recipe_fk
        foreign key (recipe) references recipes
            on delete cascade;

alter table quantity
    drop constraint recipe_fk,
    add constraint recipe_fk
        foreign key (recipe) references recipes
            on delete cascade;

alter table favorites
    drop constraint users_fk,
    add constraint users_fk
        foreign key ("user") references users
            on delete cascade,
    drop constraint recipe_fk,
    add constraint recipe_fk
        foreign key (recipe) references recipes
            on delete cascade;