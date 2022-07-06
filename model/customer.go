package model

type Customer struct {
	Id            int         `gorm:"primaryKey" json:"customerId"`
	CustomerName  string      `json:"customerName"`
	MobilePhoneNo string      `json:"mobilePhoneNo"`
	IsMember      bool        `json:"isMember"`
	Discounts     []*Discount `gorm:"many2many:m_customer_discount" json:"discounts"`
	BaseModel     BaseModel   `gorm:"embedded" json:"baseModel"`
}

func (Customer) TableName() string {
	return "m_customer"
}
