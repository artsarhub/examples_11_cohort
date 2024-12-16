CREATE TABLE orders
(
    order_id      INTEGER PRIMARY KEY,
    customer_name TEXT,
    customer_city TEXT,
    product_name  TEXT,
    product_price REAL,
    order_date    DATE
);

INSERT INTO orders (order_id, customer_name, customer_city, product_name, product_price, order_date)
VALUES (1, 'Alice', 'New York', 'Laptop', 1000.00, '2024-11-01'),
       (2, 'Bob', 'Los Angeles', 'Smartphone', 500.00, '2024-11-02'),
       (3, 'Alice', 'New York', 'Tablet', 300.00, '2024-11-03'),
       (4, 'Charlie', 'Chicago', 'Laptop', 1000.00, '2024-11-04'),
       (5, 'Bob', 'Los Angeles', 'Tablet', 300.00, '2024-11-05');

CREATE TABLE customers
(
    customer_id INTEGER PRIMARY KEY,
    name       TEXT,
    city       TEXT
);

INSERT INTO customers (customer_id, name, city)
VALUES (1, 'Alice', 'New York'),
       (2, 'Bob', 'Los Angeles'),
       (3, 'Charlie', 'Chicago');

CREATE TABLE products
(
    product_id INTEGER PRIMARY KEY,
    name      TEXT,
    price     REAL
);

INSERT INTO products (product_id, name, price)
VALUES (1, 'Laptop', 1000.00),
       (2, 'Smartphone', 500.00),
       (3, 'Tablet', 300.00);

CREATE TABLE new_orders
(
    order_id    INTEGER PRIMARY KEY,
    customer_id INTEGER,
    product_id  INTEGER,
    order_date  DATE,
    FOREIGN KEY (customer_id) REFERENCES customers (customer_id),
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);

INSERT INTO new_orders (order_id, customer_id, product_id, order_date)
VALUES (1, 1, 1, '2024-11-01'),
       (2, 2, 2, '2024-11-02'),
       (3, 1, 3, '2024-11-03'),
       (4, 3, 1, '2024-11-04'),
       (5, 2, 3, '2024-11-05');

DROP TABLE orders;

------------------------------------------------------------------------------------------------------------------------

ALTER TABLE new_orders RENAME TO orders;

SELECT customer_id AS id, name FROM customers WHERE name = 'Alice';

SELECT COUNT(*) AS total_orders FROM orders;

SELECT o.order_id, c.name, p.name AS product_name, o.order_date
FROM orders o
         LEFT JOIN customers c ON o.customer_id = c.customer_id
         LEFT JOIN products p ON o.product_id = p.product_id
WHERE c.name = 'Alice';

WITH cte AS (SELECT c.name, COUNT(o.order_id) AS total_orders
             FROM orders o
                      JOIN customers c ON o.customer_id = c.customer_id
             GROUP BY c.name)
SELECT *
FROM cte
WHERE name = 'Alice';

explain query plan
SELECT c.name, SUM(p.price) AS total_spent, COUNT(o.order_id) AS total_orders
FROM orders o
         JOIN customers c ON o.customer_id = c.customer_id
         JOIN products p ON o.product_id = p.product_id
GROUP BY c.name
ORDER BY total_spent DESC;

begin;
delete from customers where customer_id = 1;
rollback;

select * from customers;
