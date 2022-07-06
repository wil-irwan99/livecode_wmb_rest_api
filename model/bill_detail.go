package model

import (
	"encoding/json"
)

type BillDetail struct {
	Id          int       `gorm:"primaryKey" json:"billDetailId"`
	BillID      int       `json:"billId"`
	Bill        Bill      `gorm:"foreignKey:BillID" json:"bill"`
	MenuPriceID int       `json:"menuPriceId"`
	MenuPrice   MenuPrice `gorm:"foreignKey:MenuPriceID" json:"menuPrice"`
	Qty         int       `json:"qty"`
	BaseModel   BaseModel `gorm:"embedded" json:"baseModel"`
}

func (BillDetail) TableName() string {
	return "t_bill_detail"
}

func (b *BillDetail) ToString() string {
	billDetail, err := json.MarshalIndent(b, "", " ")
	if err != nil {
		return ""
	}
	return string(billDetail)
}
