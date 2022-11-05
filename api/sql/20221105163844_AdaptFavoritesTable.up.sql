alter table favorites
    add id BIGSERIAL not null;

alter table favorites
    drop constraint favorites_pk;

alter table favorites
    add constraint favorites_pk
        primary key (id);

alter table favorites
    add constraint unique_favorites
        unique ("user", recipe);