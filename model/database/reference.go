package database

import (
	"time"

	"github.com/uptrace/bun"
)

type Bank struct {
	bun.BaseModel `bun:"table:reference.bank,alias:bnk"`
	ID            string    `bun:"id,pk,type:uuid,unique,notnull,default:gen_random_uuid()"`
	BankID        string    `bun:"bank_id,unique,notnull"`
	BankName      string    `bun:"bank_name,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,nullzero"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero"`
	DeletedAt     time.Time `bun:"deleted_at,soft_delete,nullzero"`
}
