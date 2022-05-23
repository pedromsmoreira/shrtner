SET sql_safe_updates = FALSE;

USE defaultdb;
DROP DATABASE IF EXISTS shrtener CASCADE;
CREATE DATABASE IF NOT EXISTS shrtener;

USE shrtener;

CREATE TABLE urls (
    tenant STRING(36),
    id STRING(36),
    created_date TIMESTAMP,
    modified_date TIMESTAMP,
    original_url STRING,
    short_url STRING,
    PRIMARY KEY (tenant, id)
)

