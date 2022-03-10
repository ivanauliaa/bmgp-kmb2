-- [create relationships section]

-- drop foreign key constraint in user_payment_method_details
ALTER TABLE user_payment_method_details
DROP FOREIGN KEY FK_PaymentMethod_PaymentMethodDetail;

ALTER TABLE user_payment_method_details
DROP FOREIGN KEY FK_User_PaymentMethodDetail;

-- drop foreign key constraint in addresses
ALTER TABLE addresses
DROP FOREIGN KEY FK_User_Address;

-- drop foreign key constraint in payment_method_descriptions
ALTER TABLE payment_method_descriptions
DROP FOREIGN KEY FK_PaymentMethod_PaymentMethodDescription;

-- drop user_payment_method_details table
DROP TABLE IF EXISTS user_payment_method_details;

-- drop addresses table
DROP TABLE IF EXISTS addresses;

-- drop payment_method_descriptions table
DROP TABLE IF EXISTS payment_method_descriptions;

-- [end of create relationships section]

-- restore shippings table
CREATE TABLE IF NOT EXISTS shippings (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  ongkos_dasar FLOAT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- rename shippings -> couriers
ALTER TABLE shippings RENAME TO couriers;

-- drop ongkos_dasar_column
ALTER TABLE couriers DROP COLUMN IF EXISTS ongkos_dasar;

-- drop couriers table
DROP TABLE IF EXISTS couriers;

-- drop products, product_types, operators, product_descriptions, payment_methods table
DROP TABLE IF EXISTS products, product_types, operators, product_descriptions, payment_methods;

-- drop users table
DROP TABLE IF EXISTS users;