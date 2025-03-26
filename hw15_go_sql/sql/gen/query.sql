-- name: UserCreate :one
insert into hw15_go_sql.users(name, email, password)
VALUES ($1, $2, $3) returning id;

-- name: Users :many
select * from hw15_go_sql.users limit $1 offset $2;

-- name: UserById :one
select * from hw15_go_sql.users where id = $1 limit $2;

-- name: ProductCreate :one
insert into hw15_go_sql.products(name, price)
VALUES ($1, $2) returning id;

-- name: Products :many
select * from hw15_go_sql.products limit $1 offset $2;

-- name: ProductUpdatePrice :one
update hw15_go_sql.products set price = $2 where id = $1 returning id;

-- name: ProductsByIds :many
select * from hw15_go_sql.products where id = ANY($1::bigint[]) limit $2 ;

-- name: OrderCreate :one
insert into hw15_go_sql.orders(user_id, total_amount)
VALUES ($1, $2) returning id;

-- name: Orders :many
select * from hw15_go_sql.orders limit $1 offset $2;

-- name: OrderById :one
select * from hw15_go_sql.orders where id = $1 limit $2;

-- name: OrderUpdateTotalAmount :one
update hw15_go_sql.orders set total_amount = $2 where id = $1 returning id;

-- name: LinkOrderToProductCreate :one
insert into hw15_go_sql.link_order_to_product(order_id, product_id, count)
VALUES ($1, $2, $3) returning id;

-- name: LinkOrderToProductByOrderId :many
select * from hw15_go_sql.link_order_to_product where order_id = $1;


