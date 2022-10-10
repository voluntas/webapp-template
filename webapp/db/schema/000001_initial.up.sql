-- ここは適当に変えること

CREATE TABLE account (
    pk SERIAL PRIMARY KEY,

    display_name VARCHAR(255) NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);