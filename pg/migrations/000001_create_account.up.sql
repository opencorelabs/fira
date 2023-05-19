create type account_namespace as enum (
    'consumer',
    'developer'
);

create type credentials_type as enum (
    'email_password',
    'oauth2'
);

create table if not exists accounts (
    account_id bigint not null primary key,
    namespace account_namespace not null,
    valid boolean not null,
    credentials_type credentials_type not null,
    credentials jsonb not null,
    name text not null,
    avatar_url text not null,
    email text not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null
);

create index if not exists accounts_namespace_idx on accounts (namespace);
create index if not exists credentials_type_idx on accounts using gin (credentials);
