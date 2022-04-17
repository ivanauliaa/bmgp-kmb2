FROM mysql:8.0.23

COPY ./database/migration.sql /docker-entrypoint-initdb.d/

