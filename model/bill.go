package model

import (
	"time"
)

type Bill struct {
	Id          int       `gorm:"primaryKey" json:"billId"`
	TransDate   time.Time `json:"transDate"`
	CustomerID  int       `json:"customerId"`
	Customer    Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	TableID     int       `json:"tableId"`
	Table       Table     `gorm:"foreignKey:TableID" json:"table"`
	TransTypeID string    `json:"transTypeId"`
	TransType   TransType `gorm:"foreignKey:TransTypeID" json:"transType"`
	BaseModel   BaseModel `gorm:"embedded" json:"baseModel"`
}

func (Bill) TableName() string {
	return "t_bill"
}
