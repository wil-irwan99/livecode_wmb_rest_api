package model

type Discount struct {
	Id          int       `gorm:"primaryKey" json:"discountId"`
	Description string    `json:"description"`
	Pct         int       `json:"pct"`
	BaseModel   BaseModel `gorm:"embedded" json:"baseModel"`
}

func (Discount) TableName() string {
	return "m_discount"
}
