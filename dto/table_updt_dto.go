package dto

type TableUpdtDto struct {
	Id               int    `json:"tableId"`
	TableDescription string `json:"tableDescription"`
	IsAvailable      bool   `json:"isAvailable"`
}
