package reference

import (
	"money-recap-app/model/database"
	"money-recap-app/model/dto"
)

type IBankMapper interface {
	ToInsertBank(request dto.RequestCreateBank) database.Bank
	ToDropdownBank(bankData []database.Bank) []dto.ResponseDropdown
}
