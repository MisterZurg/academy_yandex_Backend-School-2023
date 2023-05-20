-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS couriers
(
    courier_id      SERIAL NOT NULL PRIMARY KEY,
    courier_type    VARCHAR NOT NULL,
    regions         NUMERIC[] NOT NULL,
    working_hours   VARCHAR[] NOT NULL
);

CREATE TABLE IF NOT EXISTS orders
(
    order_id        SERIAL NOT NULL PRIMARY KEY,
    weight          NUMERIC NOT NULL,
    regions         NUMERIC NOT NULL,
    delivery_hours  VARCHAR[] NOT NULL,
    cost            NUMERIC NOT NULL,
    completed_time  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS couriers_groups
(
    courier_id      SERIAL NOT NULL,
    group_id        NUMERIC NOT NULL
);

CREATE TABLE IF NOT EXISTS groups_orders
(
    group_id        SERIAL NOT NULL,
    order_id        NUMERIC NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE couriers;

DROP TABLE orders;

DROP TABLE couriers_groups;

DROP TABLE groups_orders;
-- +goose StatementEnd
