-- 3a
UPDATE products
SET name = 'product dummy'
WHERE id = 1;

SELECT * FROM products where id = 1;

-- 3b
UPDATE transaction_details
SET qty = 3
WHERE product_id = 3;

SELECT * FROM transaction_details WHERE product_id = 3;