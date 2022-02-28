package reference

import "money-recap-app/model/database"

type IBankRepository interface {
	CreateBankDB(traceId string, bankData database.Bank) error
	GetListBankDB(traceId string) (int, []database.Bank, error)
}
