-- +goose Up
-- +goose StatementBegin
INSERT INTO hw15_go_sql.users(id, name, email, password)
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

INSERT INTO hw15_go_sql.products(name, price)
VALUES ( 'Молоток', 100.00),
       ( 'Щипсы', 200.00),
       ( 'Перфоратор', 300.00),
       ( 'Кирка', 400.00),
       ( 'MacBook', 5000.00),
       ( 'Iphone', 8000.00),
       ( 'Бумага', 1000.00),
       ( 'Гарнитура', 800.00),
       ( 'Пылесос', 900.00),
       ( 'Пылесос Redmi', 2000.00),
       ( 'Гвозди в пачке по 500 штук', 3000.00),
       ( 'Прищепки для белья', 400.00),
       ( 'Шины Inforce', 6000.00),
       ( 'Бумеранг Скайвокера', 100000.00);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM hw15_go_sql.products;
DELETE FROM hw15_go_sql.users;
-- +goose StatementEnd
