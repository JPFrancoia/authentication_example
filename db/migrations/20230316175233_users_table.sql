-- migrate:up

create extension if not exists "uuid-ossp";

create type provider_type as enum (
    'facebook'
);

create table users (
    user_id uuid not null default uuid_generate_v4 (),
    creation_time timestamp with time zone not null default now(),
    provider provider_type not null,
    email text not null,
    primary key (user_id),
    unique (email, provider)
);

create index idx_hash_user_id on users using hash (user_id);

-- migrate:down
drop index idx_hash_user_id;

drop table users;

drop type provider_type;

