package schema

import "fmt"

type Model struct {
	Order_uid    string `json:"order_uid,omitempty" binding:"required"`
	Track_number string `json:"track_number" binding:"required"`
	Entry        string `json:"entry"`
	Delivery     struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction   string `json:"transaction"`
		Request_id    string `json:"request_id"`
		Currency      string `json:"currency"`
		Provider      string `json:"provider"`
		Amount        int    `json:"amount"`
		Payment_dt    int    `json:"payment_dt"`
		Bank          string `json:"bank"`
		Delivery_cost int    `json:"delivery_cost"`
		Goods_total   int    `json:"goods_total"`
		Custom_fee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		Chrt_id      int    `json:"chrt_id"`
		Track_number string `json:"track_number"`
		Price        int    `json:"price"`
		Rid          string `json:"rid"`
		Name         string `json:"name"`
		Sale         int    `json:"sale"`
		Size         string `json:"size"`
		Total_price  int    `json:"total_price"`
		Nm_id        int    `json:"nm_id"`
		Brand        string `json:"brand"`
		Status       int    `json:"status"`
	} `json:"items"`
	Locale             string `json:"locale"`
	Internal_signature string `json:"internal_signature"`
	Customer_id        string `json:"customer_id"`
	Delivery_service   string `json:"delivery_service"`
	Shardkey           string `json:"shardkey"`
	Sm_id              int    `json:"sm_id"`
	Date_created       string `json:"date_created"`
	Oof_shard          string `json:"oof_shard"`
}

func (m *Model) PrintModel() string {
	var s string
	s += fmt.Sprintln("<ul>{")
	s += fmt.Sprintf(" <ul><li>\"order_uid\": \"%s\",</li>", m.Order_uid)
	s += fmt.Sprintf(" <li>\"track_number\": \"%s\",</li>", m.Track_number)
	s += fmt.Sprintf(" <li>\"entry\": \"%s\",</li>", m.Entry)

	s += fmt.Sprintln(" <li>\"delivery\": {</li>")
	s += fmt.Sprintf("  <ul><li>\"name\": \"%s\",</li>", m.Delivery.Name)
	s += fmt.Sprintf("  <li>\"phone\": \"%s\",</li>", m.Delivery.Phone)
	s += fmt.Sprintf("  <li>\"zip\": \"%s\",</li>", m.Delivery.Zip)
	s += fmt.Sprintf("  <li>\"city\": \"%s\",</li>", m.Delivery.City)
	s += fmt.Sprintf("  <li>\"address\": \"%s\",</li>", m.Delivery.Address)
	s += fmt.Sprintf("  <li>\"region\": \"%s\",</li>", m.Delivery.Region)
	s += fmt.Sprintf("  <li>\"email\": \"%s\",</li></ul>", m.Delivery.Email)
	s += fmt.Sprintln(" },")

	s += fmt.Sprintln(" <li>\"payment\": {</li>")
	s += fmt.Sprintf("  <ul><li>\"transaction\": \"%s\",</li>", m.Payment.Transaction)
	s += fmt.Sprintf("  <li>\"request_id\": \"%s\",</li>", m.Payment.Request_id)
	s += fmt.Sprintf("  <li>\"currency\": \"%s\",</li>", m.Payment.Currency)
	s += fmt.Sprintf("  <li>\"provider\": \"%s\",</li>", m.Payment.Provider)
	s += fmt.Sprintf("  <li>\"amount\": %v,</li>", m.Payment.Amount)
	s += fmt.Sprintf("  <li>\"payment_dt\": %v,</li>", m.Payment.Payment_dt)
	s += fmt.Sprintf("  <li>\"bank\": \"%s\",</li>", m.Payment.Bank)
	s += fmt.Sprintf("  <li>\"delivery_cost\": %v,</li>", m.Payment.Delivery_cost)
	s += fmt.Sprintf("  <li>\"goods_total\": %v,</li>", m.Payment.Goods_total)
	s += fmt.Sprintf("  <li>\"custom_fee\": %v,</li>", m.Payment.Custom_fee)
	s += fmt.Sprintln(" },</ul>")

	s += fmt.Sprintln(" <li>\"items\": [</li>")
	for _, val := range m.Items {
		s += fmt.Sprintln("  <ul><li>{</li>")
		s += fmt.Sprintf("   <li>\"chrt_id\": %v,</li>", val.Chrt_id)
		s += fmt.Sprintf("   <li>\"track_number\": \"%s\",</li>", val.Track_number)
		s += fmt.Sprintf("   <li>\"price\": %v,</li>", val.Price)
		s += fmt.Sprintf("   <li>\"rid\": \"%s\",</li>", val.Rid)
		s += fmt.Sprintf("   <li>\"name\": \"%s\",</li>", val.Name)
		s += fmt.Sprintf("   <li>\"sale\": %v,</li>", val.Sale)
		s += fmt.Sprintf("   <li>\"size\": \"%s\",</li>", val.Size)
		s += fmt.Sprintf("   <li>\"total_price\": %v,</li>", val.Total_price)
		s += fmt.Sprintf("   <li>\"nm_id\": %v,</li>", val.Nm_id)
		s += fmt.Sprintf("   <li>\"brand\": \"%s\",</li>", val.Brand)
		s += fmt.Sprintf("   <li>\"status\": %v,</li>", val.Status)
		s += fmt.Sprintln("  <li>},</li></ul>")
	}
	s += fmt.Sprintln(" <li>],</li>")

	s += fmt.Sprintf(" <li>\"locale\": \"%s\",</li>", m.Locale)
	s += fmt.Sprintf(" <li>\"internal_signature\": \"%s\",</li>", m.Internal_signature)
	s += fmt.Sprintf(" <li>\"customer_id\": \"%s\",</li>", m.Customer_id)
	s += fmt.Sprintf(" <li>\"delivery_service\": \"%s\",</li>", m.Delivery_service)
	s += fmt.Sprintf(" <li>\"shardkey\": \"%s\",</li>", m.Shardkey)
	s += fmt.Sprintf(" <li>\"sm_id\": %v,</li>", m.Sm_id)
	s += fmt.Sprintf(" <li>\"date_created\": \"%s\",</li>", m.Date_created)
	s += fmt.Sprintf(" <li>\"oof_shard\": \"%s\",</li></ul>", m.Oof_shard)
	s += fmt.Sprintln("}</ul>")
	return s
}

func (m *Model) Valid() error {
	if m.Order_uid == "" {
		err := fmt.Errorf("not valid model")
		return err
	}
	return nil
}
