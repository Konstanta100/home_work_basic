INSERT INTO hw14_sql.users(id, name, email, password)
VALUES ('53e126bd-ef87-4186-855d-29da831d06c2', 'Kostya', 'batuev1@mail.ru', 'qwerty123'),
       ('06fd80ab-3993-41e9-aed0-bf8b3d8e9fd0', 'Tomas', 'batuev2@mail.ru', 'qwerty222'),
       ('ed8dd2e7-591f-4a02-bb0b-b15ad8288a7e', 'Игорь', 'batuev3@mail.ru', 'qwerty333'),
       ('d2cb22e2-f5c7-404f-912c-81180519981c', 'Максим Степанов', 'batuev4@mail.ru', 'qwerty444'),
       ('a7534b72-17ed-4f38-884a-dc4ca793fe04', 'Test Order for B2B', 'batuev5@mail.ru', 'qwerty555'),
       ('91aa2455-f11d-49a7-aa12-5e3d86ad39a9', 'Дядя Коля', 'kbtest1@gmail.com', 'qwerty666'),
       ('2207ab75-ec10-4d7f-9ecd-baf3658c77ac', 'Настенька', 'kbatest2@gmail.com', 'qwerty777'),
       ('17cf58de-525d-4956-a9a5-d05255c4aaf8', 'Гриша', 'kbatest3@gmail.com', 'qwerty888'),
       ('708bc763-d23c-406c-ba29-72fa443ebd19', 'Тимоша', 'kbatest4@gmail.com', 'qwerty999'),
       ('8433d028-7c42-41d0-839c-f090079082c7', 'Sofya', 'kbatest5@gmail.com', 'qwerty987'),
       ('6ecefe44-9522-4efc-96e2-62566d333903', 'Test user', 'kbatest6@gmail.com', 'qwerty999');

INSERT INTO hw14_sql.orders(id, user_id, order_date, total_amount)
VALUES (1, '53e126bd-ef87-4186-855d-29da831d06c2', '2025-01-01 12:00:00', 1000.00),
       (2, '53e126bd-ef87-4186-855d-29da831d06c2', '2025-01-02 12:00:00', 2000.00),
       (3, '53e126bd-ef87-4186-855d-29da831d06c2', '2025-01-05 12:00:00', 3000.00),
       (4, '06fd80ab-3993-41e9-aed0-bf8b3d8e9fd0', '2025-02-01 13:00:00', 4000.00),
       (5, 'ed8dd2e7-591f-4a02-bb0b-b15ad8288a7e', '2025-02-02 14:00:00', 5000.00),
       (6, 'd2cb22e2-f5c7-404f-912c-81180519981c', '2025-02-03 15:00:00', 6000.00),
       (7, 'a7534b72-17ed-4f38-884a-dc4ca793fe04', '2025-02-04 16:00:00', 7000.00),
       (8, 'a7534b72-17ed-4f38-884a-dc4ca793fe04', '2025-02-06 16:00:00', 8000.00),
       (9, 'a7534b72-17ed-4f38-884a-dc4ca793fe04', '2025-02-10 16:00:00', 9000.00),
       (10, '91aa2455-f11d-49a7-aa12-5e3d86ad39a9', '2025-02-05 17:00:00', 10000.00),
       (11, '2207ab75-ec10-4d7f-9ecd-baf3658c77ac', '2025-02-06 18:00:00', 11000.00),
       (12, '17cf58de-525d-4956-a9a5-d05255c4aaf8', '2025-02-07 19:00:00', 12000.00),
       (13, '708bc763-d23c-406c-ba29-72fa443ebd19', '2025-02-08 20:00:00', 13000.00),
       (14, '708bc763-d23c-406c-ba29-72fa443ebd19', '2025-02-11 21:00:00', 140000.00),
       (15, '8433d028-7c42-41d0-839c-f090079082c7', '2025-02-09 21:00:00', 15000.00),
       (16, '6ecefe44-9522-4efc-96e2-62566d333903', '2025-02-10 22:00:00', 16000.00);

INSERT INTO hw14_sql.products(id, name, price)
VALUES (1, 'Молоток', 100.00),
       (2, 'Щипсы', 200.00),
       (3, 'Перфоратор', 300.00),
       (4, 'Кирка', 400.00),
       (5, 'MacBook', 5000.00),
       (6, 'Iphone', 8000.00),
       (7, 'Бумага', 1000.00),
       (8, 'Гарнитура', 800.00),
       (9, 'Пылесос', 900.00),
       (10, 'Пылесос Redmi', 2000.00),
       (11, 'Гвозди в пачке по 500 штук', 3000.00),
       (12, 'Прищепки для белья', 400.00),
       (13, 'Шины Inforce', 6000.00),
       (14, 'Бумеранг Скайвокера', 100000.00);

INSERT INTO hw14_sql.link_orders_to_products(order_id, product_id, count)
VALUES (1, 1, 10),
       (2, 2, 5),
       (2, 1, 3),
       (2, 3, 1),
       (2, 4, 1),
       (3, 10, 1),
       (3, 7, 1),
       (4, 11, 1),
       (4, 8, 1),
       (4, 1, 2),
       (5, 5, 1),
       (6, 13, 1),
       (7, 5, 1),
       (7, 10, 1),
       (8, 6, 1),
       (9, 6, 1),
       (9, 2, 5),
       (10, 5, 2),
       (11, 5, 2),
       (11, 1, 6),
       (11, 4, 1),
       (12, 9, 10),
       (12, 11, 1),
       (13, 7, 6),
       (13, 13, 1),
       (13, 12, 2),
       (13, 2, 1),
       (14, 8, 10),
       (14, 13, 1),
       (15, 5, 2),
       (15, 9, 1),
       (15, 1, 1),
       (15, 12, 10),
       (15, 7, 1),
       (16, 6, 2);


SELECT * FROM hw14_sql.users;
SELECT * FROM hw14_sql.users WHERE name LIKE '%Test%';

SELECT * FROM hw14_sql.products;
SELECT * FROM hw14_sql.products WHERE products.price >= 2000;

SELECT * FROM hw14_sql.orders;
SELECT * FROM hw14_sql.orders WHERE user_id = '53e126bd-ef87-4186-855d-29da831d06c2';

-- EXPLAIN analyse
SELECT
    SUM(distinct orders.total_amount),
    SUM(link_orders_to_products.count),
    ROUND(SUM(products.price * link_orders_to_products.count) / SUM(link_orders_to_products.count), 2)
FROM hw14_sql.orders
         INNER JOIN hw14_sql.link_orders_to_products ON orders.id = link_orders_to_products.order_id
         INNER JOIN hw14_sql.products ON link_orders_to_products.product_id = products.id
WHERE user_id = '53e126bd-ef87-4186-855d-29da831d06c2'
GROUP BY orders.user_id;

UPDATE hw14_sql.products SET name  = 'Джедайский меч', price = 1000000.00 WHERE id = 14;
DELETE FROM hw14_sql.products WHERE id = 14;
UPDATE hw14_sql.users SET password = 'test1111' WHERE id = 'a7534b72-17ed-4f38-884a-dc4ca793fe04';
UPDATE hw14_sql.users SET email    = 'test1111@test-mail.ru', password = '' WHERE id = 'a7534b72-17ed-4f38-884a-dc4ca793fe04';
DELETE FROM hw14_sql.users WHERE name = 'Test user';
DELETE FROM hw14_sql.users WHERE id = 'a7534b72-17ed-4f38-884a-dc4ca793fe04';
DELETE FROM hw14_sql.orders WHERE orders.order_date between '2025-02-01 13:00:00' and '2025-02-06 13:00:00';