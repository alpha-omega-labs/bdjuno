package database

import (
	"encoding/json"
	"fmt"

	"github.com/forbole/bdjuno/types"
)

// SaveEvmParams allows to store the given evm parameters inside the database
func (db *Db) SaveEvmParams(params types.EvmParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling params: %s", err)
	}

	stmt := `
INSERT INTO evm_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
      	height = excluded.height
WHERE evm_params.height <= excluded.height`
	_, err = db.Sql.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing evm params: %s", err)
	}

	return nil
}
