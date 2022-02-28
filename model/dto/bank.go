package dto

type RequestCreateBank struct {
	BankID   string `json:"bank_id" form:"bank_id"`
	BankName string `json:"bank_name" form:"bank_name"`
}
