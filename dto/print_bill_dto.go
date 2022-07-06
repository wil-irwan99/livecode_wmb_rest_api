package dto

import "time"

type PrintBill struct {
	BillId     int         `json:"bill_id"`
	TransDate  time.Time   `'json:"trans_date"`
	Orders     []OrderList `json:"orders"`
	Discount   int         `json:"member_discount"`
	TotalPrice float32     `json:"total_price"`
}

type OrderList struct {
	Name  string `json:"menu"`
	Qty   int    `json:"qty"`
	Price int    `json:"price"`
}
