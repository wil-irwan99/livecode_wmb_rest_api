package dto

type TableAddDto struct {
	TableDescription string `json:"tableDescription"`
	IsAvailable      bool   `json:"isAvailable"`
}
