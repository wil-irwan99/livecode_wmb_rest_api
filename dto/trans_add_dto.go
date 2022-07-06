package dto

type TransAddDto struct {
	CustId      int    `json:"custId"`
	TransTypeId string `json:"transTypeId"`
	MejaId      int    `json:"tableId"`
	MenuPriceId []int  `json:"menuPriceId"`
	Qty         []int  `json:"qty"`
}
