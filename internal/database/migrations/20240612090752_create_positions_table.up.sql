create positions (
    id varchar primary key,
    name varchar,
    salary int
);

create employees (
       id varchar primary key,
       firstname varchar,
       lastname varchar,
       position_id references positions(id)
);