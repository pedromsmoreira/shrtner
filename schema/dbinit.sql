SET sql_safe_updates = FALSE;

USE defaultdb
DROP DATABASE IF EXISTS shrtener CASCADE;
CREATE DATABASE IF NOT EXISTS shrtener;

USE shrtener;

CREATE TABLE urls (
    id STRING(36) PRIMARY KEY,
    datecreated TIMESTAMP,
    url STRING,
    shortened STRING
)

