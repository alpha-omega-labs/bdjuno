package intrarelayer

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/desmos-labs/juno/types"
	"github.com/forbole/bdjuno/database"
	irmtypes "github.com/tharsis/evmos/x/intrarelayer/types"
)

// TODO: do something with these messages
func HandleMsg(
	tx *juno.Tx, index int, msg sdk.Msg,
	cdc codec.Codec, db *database.Db,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch msg.(type) {
	// Converters
	case *irmtypes.MsgConvertCoin:
		return nil
	case *irmtypes.MsgConvertERC20:
		return nil
	default:
		return nil
	}
}
