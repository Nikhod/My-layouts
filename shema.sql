create table notes
(
    id         bigserial                     not null primary key,
    content    json                          not null,
    user_id    integer references users (id) not null,
    active     boolean                       not null default true,
    created_at timestamptz                   not null default current_timestamp,
    updated_at timestamptz                   not null default current_timestamp,
    deleted_at timestamptz                   not null default current_timestamp
);

create table users
(
    id         bigserial   not null primary key,
    name       text        not null,
    login      text        not null unique,
    password   text        not null,
    active     boolean     not null default true,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz not null default current_timestamp
);

alter table users
add tokens text not null;