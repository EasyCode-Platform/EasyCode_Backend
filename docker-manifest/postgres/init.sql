-- init ec-backend


create database ec-backend;

\c ec-backend;

create user ec-backend with encrypted password 'scut2023';

grant all privileges on database ec-backend to ec-backend;

CREATE EXTENSION pg_trgm;

CREATE EXTENSION btree_gin;

-- apps
create table if not exists apps (
    id                      bigserial                       not null primary key,
    uid                     uuid default gen_random_uuid()  not null,
    team_id                 bigserial                       not null,
    name                    text                  not null,
    component_id            text                            not null,
    config                  jsonb,
    created_at              timestamp                       not null,
    updated_at              timestamp                       not null,
);
