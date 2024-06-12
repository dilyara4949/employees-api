
create table positions (
                           id varchar primary key,
                           name varchar,
                           salary int
                               created_at timestamptz default current_timestamp,
                           updated_at timestamptz default current_timestamp,
);

create table employees (
                           id varchar primary key,
                           firstname varchar,
                           lastname varchar,
                           position_id references positions(id)
                               created_at timestamptz default current_timestamp,
                           updated_at timestamptz default current_timestamp,
);

CREATE INDEX idx_employees_first_name ON employees(first_name);
CREATE INDEX idx_employees_last_name ON employees(last_name);
CREATE INDEX idx_employees_first_last_name ON employees(first_name, last_name);
