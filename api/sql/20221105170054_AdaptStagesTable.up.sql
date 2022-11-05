alter table stages
    add id BIGSERIAL not null,
    drop constraint stages_pk,
    add constraint stages_pk
        primary key (id),
    add constraint unique_ordered_stage
        unique (recipe,"order");