-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS hw15_go_sql;

CREATE TABLE IF NOT EXISTS hw15_go_sql.users
(
    id       uuid PRIMARY KEY default general.new_uuid(),
    name     TEXT         NOT NULL,
    email    VARCHAR(255) NOT NULL UNIQUE,
    password TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS hw15_go_sql.orders
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      uuid           NOT NULL,
    append_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount NUMERIC(10, 2) NOT NULL,
    CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES hw15_go_sql.users (id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS hw15_go_sql.products
(
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT           NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS hw15_go_sql.link_order_to_product
(
    id         BIGSERIAL PRIMARY KEY,
    order_id   BIGINT,
    product_id BIGINT,
    count      BIGINT NOT NULL,
    CONSTRAINT fk_order
    FOREIGN KEY (order_id)
    REFERENCES hw15_go_sql.orders (id)
    ON DELETE CASCADE,
    CONSTRAINT fk_product
    FOREIGN KEY (product_id)
    REFERENCES hw15_go_sql.products (id)
    ON DELETE CASCADE
);

CREATE INDEX idx_orders_user_id ON hw15_go_sql.orders(user_id);
CREATE INDEX idx_link_order_to_product_id ON hw15_go_sql.link_order_to_product(product_id);
CREATE INDEX idx_products_price ON hw15_go_sql.products(price);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS hw15_go_sql.idx_products_price;
DROP INDEX IF EXISTS hw15_go_sql.idx_link_order_to_product_product_id;
DROP INDEX IF EXISTS hw15_go_sql.idx_orders_user_id;

DROP TABLE IF EXISTS hw15_go_sql.link_order_to_product;
DROP TABLE IF EXISTS hw15_go_sql.products;
DROP TABLE IF EXISTS hw15_go_sql.orders;
DROP TABLE IF EXISTS hw15_go_sql.users;

DROP SCHEMA IF EXISTS hw15_go_sql;
-- +goose StatementEnd
