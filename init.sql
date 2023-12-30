CREATE ROLE orders_user WITH LOGIN SUPERUSER PASSWORD '123Qweqwe.';
CREATE DATABASE order_database;

\c order_database;

CREATE TABLE public.orders (
                               order_id varchar(100) NOT NULL,
                               creation_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               "json" json NOT NULL,
                               CONSTRAINT orders_pkey PRIMARY KEY (order_id)
);