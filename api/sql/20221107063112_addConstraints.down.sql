alter table stages
    drop constraint recipe_fk,
    add constraint recipe_fk;

alter table quantity
    drop constraint recipe_fk,
    add constraint recipe_fk;

alter table favorites
    drop constraint users_fk,
    add constraint users_fk,
    drop constraint recipe_fk,
    add constraint recipe_fk;