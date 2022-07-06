package dto

import "time"

type BillHarian struct {
	Today      time.Time   `json:"date"`
	Orders     []OrderList `json:"orders"`
	TotalPrice float32     `json:"total_price"`
}
