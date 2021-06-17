FROM bitnami/postgresql:12.7.0
LABEL maintainer="dev@plantbook.icu" version="0.0.1"
ADD ./database/createdb/000000_create_database.sql /docker-entrypoint-initdb.d/
ADD ./database/migrations_schema/*.up.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_DB plantbook_admin
ENV POSTGRES_USER plantbook_admin
ENV POSTGRES_PASSWORD mypassword
