CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE,
    type VARCHAR(30) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    status VARCHAR(30) NOT NULL,
    executed_on TIMESTAMP NULL,
    webhook_slack VARCHAR(255) NOT NULL,
);