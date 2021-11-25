package types

import (
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

type EvmParams struct {
	evmtypes.Params
	Height int64 `json:"height" yaml:"height"`
}

// NewEvmParams allows to build a new evmParams instance
func NewEvmParams(params evmtypes.Params, height int64) EvmParams {
	return EvmParams{
		Params: params,
		Height: height,
	}
}
