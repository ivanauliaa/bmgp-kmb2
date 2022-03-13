ALTER TABLE transaction_details
DROP FOREIGN KEY FK_TransactionDetail_Product;

ALTER TABLE transaction_details
DROP FOREIGN KEY FK_TransactionDetail_Transaction;

ALTER TABLE transactions
DROP FOREIGN KEY FK_Transaction_PaymentMethod;

ALTER TABLE transactions
DROP FOREIGN KEY FK_Transaction_User;

ALTER TABLE product_descriptions
DROP FOREIGN KEY FK_ProductDescription_Product;

ALTER TABLE products
DROP FOREIGN KEY FK_Product_Operator;

ALTER TABLE products
DROP FOREIGN KEY FK_Product_ProductType;

DROP TABLE IF EXISTS transaction_details, transactions, 
payment_methods, product_descriptions, operators, 
product_types, products, users;