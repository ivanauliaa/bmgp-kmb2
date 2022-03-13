-- 1a
INSERT INTO operators (id, name)
VALUES (1, 'Operator 1'), (2, 'Operator 2'),
(3, 'Operator 3'), (4, 'Operator 4'),
(5, 'Operator 5');

-- 1b
INSERT INTO product_types (id, name)
VALUES (1, 'Rumah Tangga'), (2, 'Elektronik'),
(3, 'Bahan Bangunan');

-- 1c
INSERT INTO products (
  id,
  product_type_id,
  operator_id,
  code,
  name,
  status
)
VALUES (1, 1, 3, 'RT1', 'Water Mop', '200'),
(2, 1, 3, 'RT2', 'Rak Sepatu', '200');

-- 1d
INSERT INTO products (
  id,
  product_type_id,
  operator_id,
  code,
  name,
  status
)
VALUES (3, 2, 1, 'E1', 'LED TV', '404'),
(4, 2, 1, 'E2', 'Soundbar', '200'),
(5, 2, 1, 'E3', 'Mesin Cuci', '200');

-- 1e
INSERT INTO products (
  id,
  product_type_id,
  operator_id,
  code,
  name,
  status
)
VALUES (6, 3, 4, 'BB1', 'Keramik', '404'),
(7, 3, 4, 'BB2', 'Wastafel', '404'),
(8, 3, 4, 'BB3', 'Cat Tembok', '404');

-- 1f
INSERT INTO product_descriptions (
  id, product_id, description
) VALUES (1, 1, 'Lorem'),
(2, 2, 'Ipsum'),
(3, 3, 'Dolor'),
(4, 4, 'Sit'),
(5, 5, 'Amet'),
(6, 6, 'Consectetuer'),
(7, 7, 'Adipiscing'),
(8, 8, 'Elit');

-- 1g
INSERT INTO payment_methods (id, name, status)
VALUES (1, 'OVO', 200), (2, 'Gopay', 200),
(3, 'ShopeePay', 200);

-- 1h
INSERT INTO users (
  id, name, status, dob, gender
)
VALUES (1, 'Ivan', 200, '2000-10-27', 'M'),
(2, 'Aulia', 200, '2000-10-28', 'M'),
(3, 'Rahman', 200, '2000-10-29', 'F'),
(4, 'Karim', 200, '2000-10-30', 'M'),
(5, 'Benzema', 200, '2000-10-31', 'F');

-- 1i
INSERT INTO transactions (
  id,
  user_id,
  payment_method_id,
  status,
  total_qty,
  total_price
) VALUES (1, 1, 1, 'PENDING', 6, 60000),
(2, 1, 2, 'PENDING', 15, 150000),
(3, 1, 3, 'CANCELLED', 16, 160000),
(4, 2, 1, 'PENDING', 9, 90000),
(5, 2, 2, 'SUCCESS', 18, 180000),
(6, 2, 3, 'CANCELLED', 6, 60000),
(7, 3, 1, 'SUCCESS', 8, 80000),
(8, 3, 2, 'PENDING', 18, 180000),
(9, 3, 3, 'PENDING', 11, 110000),
(10, 4, 1, 'CANCELLED', 12, 120000),
(11, 4, 2, 'PENDING', 21, 210000),
(12, 4, 3, 'PENDING', 6, 60000),
(13, 5, 1, 'CANCELLED', 15, 150000),
(14, 5, 2, 'PENDING', 16, 160000),
(15, 5, 3, 'SUCCESS', 9, 90000);

-- 1j
INSERT INTO transaction_details (
  transaction_id,
  product_id,
  status,
  qty,
  price
) VALUES (1, 1, 'AVAILABLE', 1, 10000),
(1, 2, 'PRE-ORDER', 2, 20000),
(1, 3, 'AVAILABLE', 3, 30000),
(2, 4, 'AVAILABLE', 4, 40000),
(2, 5, 'AVAILABLE', 5, 50000),
(2, 6, 'AVAILABLE', 6, 60000),
(3, 7, 'AVAILABLE', 7, 70000),
(3, 8, 'PRE-ORDER', 8, 80000),
(3, 1, 'AVAILABLE', 1, 10000),
(4, 2, 'AVAILABLE', 2, 20000),
(4, 3, 'AVAILABLE', 3, 30000),
(4, 4, 'AVAILABLE', 4, 40000),
(5, 5, 'PRE-ORDER', 5, 50000),
(5, 6, 'AVAILABLE', 6, 60000),
(5, 7, 'AVAILABLE', 7, 70000),
(6, 8, 'AVAILABLE', 8, 80000),
(6, 1, 'AVAILABLE', 1, 10000),
(6, 2, 'AVAILABLE', 2, 20000),
(7, 3, 'PRE-ORDER', 3, 30000),
(7, 4, 'AVAILABLE', 4, 40000),
(7, 1, 'AVAILABLE', 1, 10000),
(8, 5, 'AVAILABLE', 5, 50000),
(8, 6, 'AVAILABLE', 6, 60000),
(8, 7, 'AVAILABLE', 7, 70000),
(9, 8, 'AVAILABLE', 8, 80000),
(9, 1, 'AVAILABLE', 1, 10000),
(9, 2, 'AVAILABLE', 2, 20000),
(10, 3, 'AVAILABLE', 3, 30000),
(10, 4, 'PRE-ORDER', 4, 40000),
(10, 5, 'AVAILABLE', 5, 50000),
(11, 6, 'AVAILABLE', 6, 60000),
(11, 7, 'AVAILABLE', 7, 70000),
(11, 8, 'AVAILABLE', 8, 80000),
(12, 1, 'AVAILABLE', 1, 10000),
(12, 2, 'PRE-ORDER', 2, 20000),
(12, 3, 'AVAILABLE', 3, 30000),
(13, 4, 'AVAILABLE', 4, 40000),
(13, 5, 'AVAILABLE', 5, 50000),
(13, 6, 'AVAILABLE', 6, 60000),
(14, 7, 'AVAILABLE', 7, 70000),
(14, 8, 'PRE-ORDER', 8, 80000),
(14, 1, 'AVAILABLE', 1, 10000),
(15, 2, 'AVAILABLE', 2, 20000),
(15, 3, 'AVAILABLE', 3, 30000),
(15, 4, 'AVAILABLE', 4, 40000);