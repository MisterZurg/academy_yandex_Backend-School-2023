FROM postgres:15.2-alpine

ENV POSTGRES_USER postgres

ENV POSTGRES_PASSWORD password

ENV POSTGRES_DB db

EXPOSE 5432

ADD ../../../internal/db/postgres/postgres_init.sql /docker-entrypoint-initdb.d/