DROP TABLE Item;
DROP TABLE Payment;
DROP TABLE Delivery;
DROP TABLE Model;

CREATE TABLE Model(
	order_uid 			VARCHAR(256) PRIMARY KEY,
	track_number 		VARCHAR(256) NOT NULL,
	entry        		VARCHAR(256) NOT NULL,
	locale 				VARCHAR(256) NOT NULL,
	internal_signature 	VARCHAR(256) NOT NULL,
	customer_id 		VARCHAR(256) NOT NULL,
	delivery_service 	VARCHAR(256) NOT NULL,
	shardkey 			INT NOT NULL,
	sm_id 				INT NOT NULL,
	date_created 		VARCHAR(256) NOT NULL,
	oof_shard 			INT NOT NULL);


CREATE TABLE Delivery(
	id 		INT PRIMARY KEY,
	name    VARCHAR(256) NOT NULL,
	phone   VARCHAR(256) NOT NULL,
	zip     INT NOT NULL,
	city    VARCHAR(256) NOT NULL,
	address VARCHAR(256) NOT NULL,
	region  VARCHAR(256) NOT NULL,
	email   VARCHAR(256) NOT NULL,
	order_uid VARCHAR(256) REFERENCES Model NOT NUll);
		
CREATE TABLE	Payment (
	transaction   VARCHAR(256) PRIMARY KEY,
	request_id    INT NOT NULL,
	currency      VARCHAR(256) NOT NULL,
	provider      VARCHAR(256) NOT NULL,
	amount        INT NOT NULL,
	payment_dt    INT NOT NULL,
	bank          VARCHAR(256) NOT NULL,
	delivery_cost INT NOT NULL,
	goods_total   INT NOT NULL,
	custom_fee    INT NOT NULL,
	order_uid VARCHAR(256) REFERENCES Model NOT NUll);

CREATE TABLE Item (
	chrt_id      	INT PRIMARY KEY,
	track_number	VARCHAR(256) NOT NULL,
	price        	INT NOT NULL,    
	rid           	VARCHAR(256) NOT NULL,
	name          	VARCHAR(256) NOT NULL,
	sale         	INT NOT NULL,    
	size         	INT NOT NULL,   
	total_price  	INT NOT NULL,    
	nm_id        	INT NOT NULL,    
	brand        	VARCHAR(256) NOT NULL, 
	status       	INT NOT NULL,
	order_uid VARCHAR(256) REFERENCES Model NOT NUll);


DO $$
DECLARE
        total_rows integer;
BEGIN
        INSERT INTO Item values();
        INSERT INTO Payment values();
        INSERT INTO Delivery values();
        INSERT INTO Model values();
        GET DIAGNOSTICS total_rows := ROW_COUNT;
        IF total_rows != 4 THEN
                ROLLBACK;
        ELSE COMMIT;
                END IF;
		RETURN Model.order_uid 
END $$;



CREATE TABLE Model(
	order_uid 			VARCHAR(256) PRIMARY KEY,
	body json);