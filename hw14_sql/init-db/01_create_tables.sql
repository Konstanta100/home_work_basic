CREATE SCHEMA IF NOT EXISTS hw14_sql;

CREATE TABLE IF NOT EXISTS hw14_sql.users
(
    id       uuid PRIMARY KEY default gen_random_uuid(),
    name     TEXT         NOT NULL,
    email    VARCHAR(255) NOT NULL UNIQUE,
    password TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS hw14_sql.orders
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      uuid           NOT NULL,
    order_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount NUMERIC(10, 2) NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES hw14_sql.users (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS hw14_sql.products
(
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT           NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS hw14_sql.link_orders_to_products
(
    order_id   BIGINT,
    product_id BIGINT,
    count      INT NOT NULL,
    CONSTRAINT fk_order
        FOREIGN KEY (order_id)
            REFERENCES hw14_sql.orders (id)
            ON DELETE CASCADE,
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
            REFERENCES hw14_sql.products (id)
            ON DELETE CASCADE
);

CREATE INDEX idx_orders_user_id ON hw14_sql.orders(user_id);
CREATE INDEX idx_link_orders_to_products_product_id ON hw14_sql.link_orders_to_products(product_id);
CREATE INDEX idx_products_price ON hw14_sql.products(price);