package evm

import (
	"github.com/forbole/bdjuno/database"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
	_ modules.BlockModule   = &Module{}
)

// Module represent x/evm module
type Module struct {
	encodingConfig *params.EncodingConfig
	db             *database.Db
	evmClient      evmtypes.QueryClient
}

// NewModule returns a new Module instance
func NewModule(
	encodingConfig *params.EncodingConfig, db *database.Db, evmClient evmtypes.QueryClient,
) *Module {
	return &Module{
		encodingConfig: encodingConfig,
		db:             db,
		evmClient:      evmClient,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "evm"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *types.Tx) error {
	return HandleMsg(tx, index, msg, m.encodingConfig.Marshaler, m.db)
}

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(b *tmctypes.ResultBlock, _ []*types.Tx, _ *tmctypes.ResultValidators) error {
	return HandleBlock(b.Block.Height, m.evmClient, m.db)
}
