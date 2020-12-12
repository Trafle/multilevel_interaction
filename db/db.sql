 DROP SCHEMA IF EXISTS payment_system;
 CREATE SCHEMA IF NOT EXISTS payment_system;
 USE payment_system;
 
 CREATE TABLE accounts(
	id int NOT NULL AUTO_INCREMENT,
    balance int NOT NULL DEFAULT 0,
    lastOperationTime datetime NULL,
    PRIMARY KEY(id)
 );
 
 INSERT INTO accounts(balance, lastOperationTime)
 VALUES
	(2000, '2015-11-05 14:29:36.11'),
	(1000, '2019-09-29 15:14:18.32'),
    (5300, '2019-11-14 10:22:10.32'),
    (3450, '2020-12-12 18:11:24.45');
