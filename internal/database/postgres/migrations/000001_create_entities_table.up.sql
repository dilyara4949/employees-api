CREATE TABLE positions (
                           id VARCHAR PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           salary INT NOT NULL,
                           created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMPTZ
);

CREATE TABLE employees (
                           id VARCHAR PRIMARY KEY,
                           first_name VARCHAR(255) NOT NULL,
                           last_name VARCHAR(255) NOT NULL,
                           position_id VARCHAR REFERENCES positions(id),
                           created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMPTZ
);