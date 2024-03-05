CREATE SCHEMA IF NOT EXISTS task_schema;

CREATE TABLE IF NOT EXISTS task_schema.task (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100),
    description TEXT,
    resolved BOOL DEFAULT false
);