-- 1
SELECT * FROM transactions WHERE user_id = 1
UNION
SELECT * FROM transactions WHERE user_id = 2;

-- 2
SELECT SUM(total_price) FROM transactions WHERE user_id = 1;

-- 3
SELECT COUNT(transaction_id) FROM transaction_details WHERE product_id IN (SELECT id FROM products WHERE product_type_id = 2);

-- 4
SELECT p.*, pt.name AS pt_name FROM products p INNER JOIN product_types pt ON p.product_type_id = pt.id;

-- 5
SELECT t.*, p.name AS p_name, u.name AS u_name FROM transactions t JOIN products p JOIN users u;

-- 6 & 7 ada di file problem2-function-up.sql

-- 8
-- add data product baru yang belum ada di transactions
INSERT INTO products (
  id,
  product_type_id,
  operator_id,
  code,
  name,
  status
)
VALUES (9, 1, 3, 'RT3', 'Vacuum Cleaner', '200');

SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transaction_details);

DELETE FROM products WHERE id = 9;