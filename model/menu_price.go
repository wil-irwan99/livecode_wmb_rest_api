package model

type MenuPrice struct {
	Id        int       `gorm:"primaryKey" json:"menuPriceId"`
	MenuID    int       `json:"menuId"`
	Menu      Menu      `gorm:"foreignKey:MenuID" json:"menu"`
	Price     int       `json:"price"`
	BaseModel BaseModel `gorm:"embedded" json:"baseModel"`
}

func (MenuPrice) TableName() string {
	return "m_menu_price"
}
