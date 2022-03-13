CREATE TABLE IF NOT EXISTS users (
  id INT(11) AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name varchar(255),
  status SMALLINT,
  dob DATE,
  gender CHAR(1),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
  id INT(11) AUTO_INCREMENT NOT NULL PRIMARY KEY,
  product_type_id INT(11) NOT NULL,
  operator_id INT(11) NOT NULL,
  code VARCHAR(50),
  name VARCHAR(100),
  status SMALLINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_types (
  id INT(11) AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS operators (
  id INT(11) AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_descriptions (
  id INT(11) AUTO_INCREMENT NOT NULL PRIMARY KEY,
  product_id INT(11) NOT NULL,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payment_methods (
  id INT(11) AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(255),
  status SMALLINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transactions (
  id INT(11) AUTO_INCREMENT NOT NULL PRIMARY KEY,
  user_id INT(11) NOT NULL,
  payment_method_id INT(11) NOT NULL,
  status VARCHAR(10),
  total_qty INT(11),
  total_price NUMERIC(25, 2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transaction_details (
  transaction_id INT(11) NOT NULL,
  product_id INT(11) NOT NULL,
  status VARCHAR(10),
  qty INT(11),
  price NUMERIC(25, 2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE transaction_details
ADD CONSTRAINT PK_TransactionDetail
PRIMARY KEY (transaction_id, product_id);

-- add fk products to product_types
ALTER TABLE products
ADD CONSTRAINT FK_Product_ProductType
FOREIGN KEY (product_type_id)
REFERENCES product_types(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- add fk products to operators
ALTER TABLE products
ADD CONSTRAINT FK_Product_Operator
FOREIGN KEY (operator_id)
REFERENCES operators(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- add fk product_descriptions to products
ALTER TABLE product_descriptions
ADD CONSTRAINT FK_ProductDescription_Product
FOREIGN KEY (product_id)
REFERENCES products(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- add fk transactions to users
ALTER TABLE transactions
ADD CONSTRAINT FK_Transaction_User
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- add fk transactions to payment_methods
ALTER TABLE transactions
ADD CONSTRAINT FK_Transaction_PaymentMethod
FOREIGN KEY (payment_method_id)
REFERENCES payment_methods(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- add fk transaction_details to transactions
ALTER TABLE transaction_details
ADD CONSTRAINT FK_TransactionDetail_Transaction
FOREIGN KEY (transaction_id)
REFERENCES transactions(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- add fk transaction_details to products
ALTER TABLE transaction_details
ADD CONSTRAINT FK_TransactionDetail_Product
FOREIGN KEY (product_id)
REFERENCES products(id)
ON DELETE CASCADE
ON UPDATE CASCADE;
