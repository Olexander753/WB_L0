DROP TABLE IS EXISTS Models

CREATE TABLE Models(
	order_uid VARCHAR(256)
	track_number VARCHAR(256) 
	entry        VARCHAR(256));

CREATE TABLE Delivery(
Name    string 
		Phone   string
		Zip     int 
		City    string
		Address string
		Region  string
		Email   string 
)
		


CREATE TABLE	Payment (
		Transaction   string `json:"transaction"`
		Request_id    int    `json:"request_id"`
		Currency      string `json:"currency"`
		Provider      string `json:"provider"`
		Amount        int    `json:"amount"`
		Payment_dt    int    `json:"payment_dt"`
		Bank          string `json:"bank"`
		Delivery_cost int    `json:"delivery_cost"`
		Goods_total   int    `json:"goods_total"`
		Custom_fee    int    `json:"custom_fee"`
)
    	Items              []Item `json:"items"`
	Locale             string `json:"locale"`
	Internal_signature string `json:"internal_signature"`
	Customer_id        string `json:"customer_id"`
	Delivery_service   string `json:"delivery_service"`
	Shardkey           int    `json:"shardkey"`
	Sm_id              int    `json:"sm_id"`
	Date_created       string `json:"date_created"`
	Oof_shard          int    `json:"oof_shard"`
}

type Item struct {
	Chrt_id      int    `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int    `json:"sale"`
	Size         int    `json:"size"`
	Total_price  int    `json:"total_price"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}

)