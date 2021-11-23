package evm

import (
	"github.com/forbole/bdjuno/database"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represent x/intrarelayer module
type Module struct {
	encodingConfig *params.EncodingConfig
	db             *database.Db
}

// NewModule returns a new Module instance
func NewModule(
	encodingConfig *params.EncodingConfig, db *database.Db,
) *Module {
	return &Module{
		encodingConfig: encodingConfig,
		db:             db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "intrarelayer"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *types.Tx) error {
	return HandleMsg(tx, index, msg, m.encodingConfig.Marshaler, m.db)
}
