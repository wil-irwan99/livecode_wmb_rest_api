package model

type Table struct {
	Id               int       `gorm:"primaryKey" json:"tableId"`
	TableDescription string    `json:"tableDescription"`
	IsAvailable      bool      `json:"isAvailable"`
	BaseModel        BaseModel `gorm:"embedded" json:"baseModel"`
}

func (Table) TableName() string {
	return "m_table"
}
