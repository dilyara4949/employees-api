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

CREATE INDEX idx_employees_first_name ON employees(first_name);
CREATE INDEX idx_employees_last_name ON employees(last_name);
CREATE INDEX idx_employees_first_last_name ON employees(first_name, last_name);
