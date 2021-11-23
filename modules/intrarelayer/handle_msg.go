package intrarelayer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forbole/bdjuno/database"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/desmos-labs/juno/types"
	irmtypes 	"github.com/tharsis/evmos/x/intrarelayer/types"

)

// TODO: do something with this messages
func HandleMsg(
	tx *juno.Tx, index int, msg sdk.Msg,
	cdc codec.Codec, db *database.Db,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
		switch msg := msg.(type) {
			// Converters
			case *irmtypes.MsgConvertCoin:
				return nil
			case *irmtypes.MsgConvertERC20:
				return nil
			// Proposals
			case *types.RegisterCoinProposal:
				return nil
			case *types.RegisterERC20Proposal:
				return nil
			case *types.ToggleTokenRelayProposal:
				return nil
			case *types.UpdateTokenPairERC20Proposal:
				return nil
			default:
				return nil
		}
	}

	return nil
}
