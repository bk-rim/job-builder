CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE,
    type VARCHAR(30) NOT NULL,
    created_on DATE NOT NULL,
    status VARCHAR(30) NOT NULL,
    executed_on DATE NULL,
    webhook_slack VARCHAR(255) NOT NULL,
);