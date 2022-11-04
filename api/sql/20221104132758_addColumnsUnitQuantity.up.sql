alter table quantity
    add column unit     varchar(30) not null default 'some',
    add column quantity float;