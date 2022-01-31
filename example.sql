drop table users;

create table users (
    id serial primary key,
    name varchar(255),
    score integer,
    created_at timestamp not null
);
