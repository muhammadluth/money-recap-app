package repository

import (
	"context"
	"fmt"
	"money-recap-app/app/log"
	"money-recap-app/model/database"
	"money-recap-app/modules/reference"

	"github.com/uptrace/bun"
)

type BankRepository struct {
	dbPostgres *bun.DB
}

func NewBankRepository(dbPostgres *bun.DB) reference.IBankRepository {
	return &BankRepository{dbPostgres}
}

func (r *BankRepository) CreateBankDB(traceId string, bankData database.Bank) error {
	if result, err := r.dbPostgres.NewInsert().Model(&bankData).Exec(context.Background()); err != nil {
		log.Error(traceId, err)
		return err
	} else if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		err = fmt.Errorf("no rows affected : %v ON INSERT reference.bank", rowsAffected)
		log.Error(traceId, err)
		return err
	}
	return nil
}

func (r *BankRepository) GetListBankDB(traceId string) (count int, data []database.Bank, err error) {
	count, err = r.dbPostgres.NewSelect().Model(&data).ScanAndCount(context.Background())
	if err != nil {
		log.Error(traceId, err)
		return 0, nil, err
	}
	return count, data, nil
}
