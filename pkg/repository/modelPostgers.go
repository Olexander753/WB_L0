package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Olexander753/WB_L0/internal/cach"
	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/jmoiron/sqlx"
)

type modelPostgres struct {
	db *sqlx.DB
	ce *cach.Cach
}

func NewModelPostgres(db *sqlx.DB, ce *cach.Cach) *modelPostgres {
	return &modelPostgres{db: db,
		ce: ce}
}

func (m *modelPostgres) InsertModel(ctx context.Context, model schema.Model) (string, error) {
	var uid string
	insertModelOrder_uid := model.Order_uid
	model.Order_uid = ""
	b, err := json.Marshal(model)
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("INSERT INTO %s values($1, $2) RETURNING order_uid ;", ModelsTable)
	row := m.db.QueryRow(query, insertModelOrder_uid, b)
	if err := row.Scan(&uid); err != nil {
		return "", err
	}
	m.ce.Models[insertModelOrder_uid] = model
	return uid, nil
	// var uid string
	// query := fmt.Sprintf("DO $$\nDECLARE\n\ttotal_rows integer;\nBEGIN\n\tINSERT INTO %s values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);\n\tINSERT INTO %s values($12,$13,$14,$15,$16,$17,$18,$19,$20);\n\tINSERT INTO %s values($21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31);\n\tINSERT INTO %s values($32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,#43);\n\tGET DIAGNOSTICS total_rows := ROW_COUNT;\n\tIF total_rows != 4 THEN\n\t\tROLLBACK;\n\tELSE COMMIT;\n\t\tEND IF;\n\tRETURN Model.order_uid;\nEND $$;", modelsTable, deliveryTable, paymentTable, itemTable)
	// row := m.db.QueryRow(query, model.Order_uid, model.Track_number, model.Entry, model.Locale, model.Internal_signature, model.Customer_id, model.Delivery_service, model.Shardkey, model.Sm_id, model.Date_created, model.Oof_shard,
	// 	model.Delivery.Id, model.Delivery.Name, model.Delivery.Phone, model.Delivery.Zip, model.Delivery.City, model.Delivery.Address, model.Delivery.Region, model.Delivery.Email, model.Order_uid,
	// 	model.Payment.Transaction, model.Payment.Request_id, model.Payment.Currency, model.Payment.Provider, model.Payment.Amount, model.Payment.Payment_dt, model.Payment.Bank, model.Payment.Delivery_cost, model.Payment.Goods_total, model.Payment.Custom_fee, model.Order_uid,
	// 	model.Items.Chrt_id, model.Items.Track_number, model.Items.Price, model.Items.Rid, model.Items.Name, model.Items.Sale, model.Items.Size, model.Items.Total_price, model.Items.Nm_id, model.Items.Brand, model.Items.Status, model.Order_uid)
	// if err := row.Scan(&uid); err != nil {
	// 	return "", err
	// }
	// return uid, nil
}

func (m *modelPostgres) SelectModel(ctx context.Context, order_uid string) (schema.Model, error) {
	// var mod struct {
	// 	Order_uid string `json:"order_uid" db:"order_uid"`
	// 	Body      string `json:"body" db:"boody"`
	// }
	//
	// query := fmt.Sprintf("SELECT * FROM %s WHERE order_uid = $1", ModelsTable)
	// err := m.db.Get(&mod, query, uid)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return model, err
	// }
	// fmt.Println(mod)
	// err = json.Unmarshal([]byte(mod.Body), &model)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return model, err
	// }
	model, err := m.ce.GetModelByOrder_uid(order_uid)
	return model, err
}
