package intrarelayer

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/desmos-labs/juno/types"
	"github.com/forbole/bdjuno/database"
	irmtypes "github.com/tharsis/evmos/x/intrarelayer/types"
)

// TODO: do something with this messages
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
		// switch cosmosMsg := msg.(govtypes.Content) {
		// // Proposals
		// case *irmtypes.RegisterCoinProposal:
		// 	return nil
		// case *irmtypes.RegisterERC20Proposal:
		// 	return nil
		// case *irmtypes.ToggleTokenRelayProposal:
		// 	return nil
		// case *irmtypes.UpdateTokenPairERC20Proposal:
		// 	return nil
	}

	return nil
}

// func HandleMsg(
// 	tx *juno.Tx, index int, msg sdk.Msg,
// 	govClient govtypes.QueryClient,
// 	cdc codec.Codec, db *database.Db,
// ) error {
// 	if len(tx.Logs) == 0 {
// 		return nil
// 	}

// 	switch cosmosMsg := msg.(type) {
// 	case *govtypes.MsgSubmitProposal:
// 		return handleMsgSubmitProposal(tx, index, cosmosMsg, govClient, cdc, db)

// 	case *govtypes.MsgDeposit:
// 		return handleMsgDeposit(tx, cosmosMsg, govClient, db)

// 	case *govtypes.MsgVote:
// 		return handleMsgVote(tx, cosmosMsg, db)
// 	}

// 	return nil
// }
