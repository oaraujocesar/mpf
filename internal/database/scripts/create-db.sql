DO
$do$
BEGIN
    IF NOT EXISTS (
        SELECT
        FROM pg_catalog.pg_database
        WHERE datname = 'mpf'
    ) THEN
        PERFORM dblink_exec('dbname=postgres', 'CREATE DATABASE mpf');
    END IF;
END
$do$;