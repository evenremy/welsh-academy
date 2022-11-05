alter table quantity
    add id BIGSERIAL not null;

alter table quantity
    drop constraint quantity_pk;

alter table quantity
    add constraint quantity_pk
        primary key (id);

alter table quantity
    add constraint unique_quantity
        unique (recipe,ingredient);