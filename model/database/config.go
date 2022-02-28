package database

import (
	"encoding/json"

	"github.com/uptrace/bun"
)

type JsonConfig struct {
	bun.BaseModel `bun:"table:config.json_config,alias:jsn_cnfg"`
	ID            string          `bun:"id,pk,type:uuid,unique"`
	Name          string          `bun:"name,notnull,unique"`
	Config        json.RawMessage `bun:"config,notnull"`
}
