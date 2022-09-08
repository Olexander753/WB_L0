DROP TABLE IS EXISTS Models

CREATE TABLE Models(
	order_uid 			VARCHAR(256)
	track_number 		VARCHAR(256) 
	entry        		VARCHAR(256)
	locale 				VARCHAR(256) 
	internal_signature 	VARCHAR(256) 
	customer_id 		VARCHAR(256) 
	delivery_service 	VARCHAR(256)
	shardkey 			int    
	sm_id 				int    
	date_created 		VARCHAR(256)
	oof_shard 			int    
	delivery_id 		int
	transction 
);

CREATE TABLE Delivery(
	id 		int
	name    VARCHAR(256)
	phone   VARCHAR(256)
	zip     int 
	city    VARCHAR(256)
	address VARCHAR(256)
	region  VARCHAR(256)
	email   VARCHAR(256)
);
		
CREATE TABLE	Payment (
	transaction   VARCHAR(256)
	request_id    int    
	currency      VARCHAR(256)
	provider      VARCHAR(256)
	amount        int    
	payment_dt    int    
	bank          VARCHAR(256)
	delivery_cost int    
	goods_total   int    
	custom_fee    int    
);

CREATE TABLE Item (
	chrt_id      	int   
	track_number	VARCHAR(256)
	price        	int    
	rid           	VARCHAR(256)
	name          	VARCHAR(256)
	sale         	int    
	size         	int    
	total_price  	int    
	nm_id        	int    
	brand        	VARCHAR(256) 
	status       	int    
);