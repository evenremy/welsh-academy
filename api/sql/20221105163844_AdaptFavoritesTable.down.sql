alter table quantity
    drop constraint favorites_pk,
    drop constraint unique_favorites;

alter table quantity
    drop column id;

alter table quantity
    add constraint quantity_pk
        primary key (recipe, ingredient);