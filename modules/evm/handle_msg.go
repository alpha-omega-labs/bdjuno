package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forbole/bdjuno/database"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/desmos-labs/juno/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

// HandleMsg allows to handle the different utils related to the gov module
func HandleMsg(
	tx *juno.Tx, index int, msg sdk.Msg,
	cdc codec.Codec, db *database.Db,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *evmtypes.MsgEthereumTx:
		return handleMsgEthereumTx(tx, index, cosmosMsg, cdc, db)
	}

	return nil
}

// handleMsgEthereumTx allows to properly handle a MsgEthereumTx
func handleMsgEthereumTx(
	tx *juno.Tx, index int, msg *evmtypes.MsgEthereumTx, cdc codec.Codec, db *database.Db,
) error {
	// TODO: figure it out how the addresses index works and add the transaction to it.
	var data evmtypes.TxData
	err := cdc.UnpackAny(msg.Data, &data)
	if err != nil {
		return err
	}

	hash := msg.AsTransaction().Hash()

	fmt.Println(hash.Hex())

	// Sender address is obtained from the signature and cached to msg.From
	sender, err := msg.GetSender(data.GetChainID())
	if err != nil {
		return err
	}

	senderBech32 := sdk.AccAddress(sender.Bytes()).String()
	senderHex := sender.Hex()

	to := data.GetTo()

	if to != nil {
		toBech32 := sdk.AccAddress(to.Bytes()).String()
		toHex := to.Hex()
		_ = []string{senderBech32, senderHex, toBech32, toHex}
		return nil
	}

	// when the recipient address **IS** defined, we are performing a EVM Contract creation
	contract := crypto.CreateAddress(sender, data.GetNonce())
	contractBech32 := sdk.AccAddress(contract.Bytes()).String()
	contractHex := contract.Hex()
	_ = []string{senderBech32, senderHex, contractBech32, contractHex}

	return nil
}
