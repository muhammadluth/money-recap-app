package mapper

import (
	"money-recap-app/model/database"
	"money-recap-app/model/dto"
	"money-recap-app/modules/reference"
)

type BankMapper struct {
}

func NewBankMapper() reference.IBankMapper {
	return &BankMapper{}
}

func (m *BankMapper) ToInsertBank(request dto.RequestCreateBank) database.Bank {
	payload := database.Bank{
		BankID:   request.BankID,
		BankName: request.BankName,
	}
	return payload
}

func (m *BankMapper) ToDropdownBank(bankData []database.Bank) (response []dto.ResponseDropdown) {
	for _, row := range bankData {
		payload := dto.ResponseDropdown{
			ID:    row.ID,
			Label: row.BankName,
			Value: row.BankID,
		}
		response = append(response, payload)
	}
	return response
}
