-- init ec_backend
create database ec_backend;

\c ec_backend;

create user ec_backend with encrypted password 'scut2023';

grant all privileges on database ec_backend to ec_backend;

CREATE EXTENSION pg_trgm;

CREATE EXTENSION btree_gin;

-- apps
create table if not exists apps (
    id bigserial not null primary key,
    uid uuid default gen_random_uuid() not null,
    team_id int not null,
    name text not null,
    component_id text not null,
    config jsonb,
    created_at timestamp not null,
    updated_at timestamp not null
);

--app_data
CREATE TABLE IF NOT EXISTS app_data (
     aid bytea NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

--tables
CREATE TABLE IF NOT EXISTS tables (
    tid bytea NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    app_aid bytea,
    FOREIGN KEY (app_aid) REFERENCES app_data(aid)
);

CREATE TABLE IF NOT EXISTS table_fields (
    tid bytea NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    PRIMARY KEY (tid, name),
    FOREIGN KEY (tid) REFERENCES tables(tid)
);

CREATE TABLE IF NOT EXISTS table_records (
    record_id SERIAL PRIMARY KEY,
    entity_id INT NOT NULL,  -- 添加这个字段来标识同一实体的不同记录
    tid bytea NOT NULL,
    field_name VARCHAR(255) NOT NULL,
    field_value TEXT,
    FOREIGN KEY (tid) REFERENCES tables(tid),
    FOREIGN KEY (tid,field_name) REFERENCES table_fields(tid,name)
);

GRANT ALL PRIVILEGES ON TABLE apps TO ec_backend;
GRANT ALL PRIVILEGES ON TABLE  tables TO ec_backend;
GRANT ALL PRIVILEGES ON TABLE app_data TO ec_backend;
GRANT ALL PRIVILEGES ON TABLE table_fields TO ec_backend;
GRANT ALL PRIVILEGES ON TABLE table_records TO ec_backend;