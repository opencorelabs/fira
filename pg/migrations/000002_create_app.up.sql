create table if not exists apps (
    app_id bigint not null primary key,
    name text not null,
    account_id bigint not null references accounts(account_id) on delete cascade,
    tokens jsonb not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null
);

create index if not exists apps_account_id_idx on apps(account_id, created_at desc);
