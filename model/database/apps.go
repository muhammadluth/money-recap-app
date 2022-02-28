package database

import (
	"time"

	"github.com/uptrace/bun"
)

type MoneyUsedHistory struct {
	bun.BaseModel `bun:"table:apps.money_used_history,alias:mny_usd_hstry"`
	ID            string    `bun:"id,pk,type:uuid,unique"`
	Name          string    `bun:"name,notnull"`
	Description   string    `bun:"description,notnull"`
	MoneyUsedAt   time.Time `bun:"money_used_at,notnull"`
	MoneyCurrency string    `bun:"money_currency,notnull"`
	MoneyAmount   string    `bun:"money_amount,notnull"`
}
