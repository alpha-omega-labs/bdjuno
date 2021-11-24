package evm

import (
	"context"

	"github.com/desmos-labs/juno/client"
	"github.com/rs/zerolog/log"

	"github.com/forbole/bdjuno/database"
	"github.com/forbole/bdjuno/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

// HandleBlock handles a new block by updating any eventually open proposal's status and tally result
func HandleBlock(
	height int64, evmClient evmtypes.QueryClient, db *database.Db,
) error {
	go updateParams(height, evmClient, db)

	return nil
}

func updateParams(height int64, evmClient evmtypes.QueryClient, db *database.Db) {
	header := client.GetHeightRequestHeader(height)
	evmParamsRes, err := evmClient.Params(
		context.Background(),
		&evmtypes.QueryParamsRequest{},
		header,
	)
	if err != nil {
		log.Error().Str("module", "evm").Err(err).
			Int64("height", height).
			Msg("error while getting params")
		return
	}

	err = db.SaveEvmParams(types.NewEvmParamsInline(evmParamsRes.GetParams(), height))
	if err != nil {
		log.Error().Str("module", "distribution").Err(err).
			Int64("height", height).
			Msg("error while saving params")
		return
	}
	return
}
