package model

type Menu struct {
	Id        int       `gorm:"primaryKey" json:"menuId"`
	MenuName  string    `json:"menuName"`
	BaseModel BaseModel `gorm:"embedded" json:"baseModel"`
}

func (Menu) TableName() string {
	return "m_menu"
}
