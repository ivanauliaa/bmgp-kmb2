-- create users table
CREATE TABLE IF NOT EXISTS users (
  id BIGINT AUTO_INCREMENT PRIMARY KEY
);

-- create products, product_types, operators, product_descriptions, payment_methods table
CREATE TABLE IF NOT EXISTS products (
  id BIGINT AUTO_INCREMENT PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS product_types (
  id BIGINT AUTO_INCREMENT PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS operators (
  id BIGINT AUTO_INCREMENT PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS product_descriptions (
  id BIGINT AUTO_INCREMENT PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS payment_methods (
  id BIGINT AUTO_INCREMENT PRIMARY KEY
);

-- create couriers table (id, name, created_at, updated_at)
CREATE TABLE IF NOT EXISTS couriers (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- add ongkos_dasar column in couriers
ALTER TABLE couriers ADD ongkos_dasar FLOAT;

-- rename couriers -> shippings
ALTER TABLE couriers RENAME TO shippings;

-- drop shippings table
DROP TABLE IF EXISTS shippings;

-- [create relationships section]

-- create payment_method_descriptions table with id and payment_method_id columns
CREATE TABLE IF NOT EXISTS payment_method_descriptions (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  payment_method_id BIGINT
);

-- create addresses table with id and user_id columns
CREATE TABLE IF NOT EXISTS addresses (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT
);

-- create user_payment_method_details table with id, user_id, and payment_method_id columns
CREATE TABLE IF NOT EXISTS user_payment_method_details (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT,
  payment_method_id BIGINT
);

-- add foreign key constraint in payment_method_descriptions
ALTER TABLE payment_method_descriptions
ADD CONSTRAINT FK_PaymentMethod_PaymentMethodDescription
FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id);

-- add foreign key constraint in addresses
ALTER TABLE addresses
ADD CONSTRAINT FK_User_Address
FOREIGN KEY (user_id) REFERENCES users(id);

-- add foreign key constraint in user_payment_method_details
ALTER TABLE user_payment_method_details
ADD CONSTRAINT FK_User_PaymentMethodDetail
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE user_payment_method_details
ADD CONSTRAINT FK_PaymentMethod_PaymentMethodDetail
FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id);

-- [end of create relationships section]