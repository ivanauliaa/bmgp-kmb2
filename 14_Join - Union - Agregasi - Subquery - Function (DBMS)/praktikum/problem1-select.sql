-- 2a
SELECT * FROM users WHERE gender = 'M';

-- 2b
SELECT * FROM products WHERE id = 3;

-- 2c
SELECT * FROM users created_at WHERE created_at <= CURRENT_TIMESTAMP AND created_at >= (CURRENT_TIMESTAMP - INTERVAL 7 DAY) AND name LIKE '%a%';

-- 2d
SELECT COUNT(id) FROM users WHERE gender = 'F';

-- 2e
SElECT * FROM users ORDER BY name ASC;

-- 2f
SELECT * FROM products LIMIT 5;