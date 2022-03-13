-- 6
DELIMITER $$
CREATE TRIGGER delete_all_transaction_details
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN

DECLARE v_transaction_id INT;
SET v_transaction_id = OLD.id;

DELETE FROM transaction_details WHERE transaction_id = v_transaction_id;

END$$
DELIMITER ;

SHOW TRIGGERS;

-- 6 test case
SELECT COUNT(transaction_id) FROM transaction_details WHERE transaction_id = 1;
DELETE FROM transactions WHERE id = 1;
SELECT COUNT(transaction_id) FROM transaction_details WHERE transaction_id = 1;

-- 7
DELIMITER $$
CREATE TRIGGER update_transaction_total_qty
BEFORE DELETE ON transaction_details FOR EACH ROW
BEGIN

DECLARE v_transaction_id INT;
DECLARE new_total_qty INT;

SET v_transaction_id = OLD.transaction_id;
SET new_total_qty = (SELECT SUM(qty) FROM transaction_details WHERE transaction_id = v_transaction_id) - OLD.qty;

UPDATE transactions SET total_qty = new_total_qty WHERE id = v_transaction_id;

END$$
DELIMITER ;

SHOW TRIGGERS;

-- 7 test case
SELECT total_qty FROM transactions WHERE id = 2;
DELETE FROM transaction_details WHERE transaction_id = 2 AND product_id = 4;
SELECT total_qty FROM transactions WHERE id = 2;
