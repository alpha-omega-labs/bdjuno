package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

// EvmParams contains the data of the x/evm module parameters
type EvmParams struct {
	EvmDenom     string      `json:"evm_denom" yaml:"evm_denom"`
	EnableCreate bool        `json:"enable_create" yaml:"enable_create"`
	EnableCall   bool        `json:"enable_call" yaml:"enable_call"`
	ExtraEIPs    []int64     `json:"extra_eips" yaml:"extra_eips"`
	ChainConfig  ChainConfig `json:"chain_config" yaml:"chain_config"`
	Height       int64       `json:"height" yaml:"height"`
}

type EvmParamsInline struct {
	evmtypes.Params
	Height int64 `json:"height" yaml:"height"`
}

// NewEvmParamsInline allows to build a new DistributionParams instance
func NewEvmParamsInline(params evmtypes.Params, height int64) EvmParamsInline {
	return EvmParamsInline{
		Params: params,
		Height: height,
	}
}

// ChainConfig defines the Ethereum ChainConfig parameters using *sdk.Int values
// instead of *big.Int.
type ChainConfig struct {
	// Homestead switch block (nil no fork, 0 = already homestead)
	HomesteadBlock sdk.Int `json:"homestead_block" yaml:"homestead_block"`
	// TheDAO hard-fork switch block (nil no fork)
	DAOForkBlock sdk.Int `json:"dao_fork_block" yaml:"dao_fork_block"`
	// Whether the nodes supports or opposes the DAO hard-fork
	DAOForkSupport bool `json:"dao_fork_support" yaml:"dao_fork_support"`
	// EIP150 implements the Gas price changes
	// (https://github.com/ethereum/EIPs/issues/150) EIP150 HF block (nil no fork)
	EIP150Block sdk.Int `json:"eip150_block" yaml:"eip150_block"`
	// EIP150 HF hash (needed for header only clients as only gas pricing changed)
	EIP150Hash string `json:"eip150_hash" yaml:"byzantium_block"`
	// EIP155Block HF block
	EIP155Block sdk.Int `json:"eip155_block" yaml:"eip155_block"`
	// EIP158 HF block
	EIP158Block sdk.Int `json:"eip158_block" yaml:"eip158_block"`
	// Byzantium switch block (nil no fork, 0 = already on byzantium)
	ByzantiumBlock sdk.Int `json:"byzantium_block" yaml:"byzantium_block"`
	// Constantinople switch block (nil no fork, 0 = already activated)
	ConstantinopleBlock sdk.Int `json:"constantinople_block" yaml:"constantinople_block"`
	// Petersburg switch block (nil same as Constantinople)
	PetersburgBlock sdk.Int `json:"petersburg_block" yaml:"petersburg_block"`
	// Istanbul switch block (nil no fork, 0 = already on istanbul)
	IstanbulBlock sdk.Int `json:"istanbul_block" yaml:"istanbul_block"`
	// Eip-2384 (bomb delay) switch block (nil no fork, 0 = already activated)
	MuirGlacierBlock sdk.Int `json:"muir_glacier_block" yaml:"muir_glacier_block"`
	// Berlin switch block (nil = no fork, 0 = already on berlin)
	BerlinBlock sdk.Int `json:"berlin_block" yaml:"berlin_block"`
	// London switch block (nil = no fork, 0 = already on london)
	LondonBlock sdk.Int `json:"london_block" yaml:"london_block"`
}

// NewEvmParams allows to build a new EvmParams instance
func NewEvmParams(evmDenom string, enableCreate bool, enableCall bool, extraEips []int64, chainConfig ChainConfig, height int64) *EvmParams {
	return &EvmParams{
		EvmDenom:     evmDenom,
		EnableCreate: enableCreate,
		EnableCall:   enableCall,
		ExtraEIPs:    extraEips,
		ChainConfig:  chainConfig,
		Height:       height,
	}
}
