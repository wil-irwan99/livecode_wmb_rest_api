package model

type TransType struct {
	Id          string    `gorm:"primaryKey;check:id='EI' OR id='TA'" json:"transTypeId"`
	Description string    `json:"description"`
	BaseModel   BaseModel `gorm:"embedded" json:"baseModel"`
}

func (TransType) TableName() string {
	return "m_trans_type"
}
