package config

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/module"

	ethermint "github.com/tharsis/ethermint/encoding"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	irmtypes "github.com/tharsis/evmos/x/intrarelayer/types"
)

// MakeEncodingConfig creates an EncodingConfig to properly handle all the messages
func MakeEncodingConfig(managers []module.BasicManager) func() params.EncodingConfig {
	return func() params.EncodingConfig {
		manager := mergeBasicManagers(managers)
		encoding := ethermint.MakeConfig(manager)
		evmtypes.RegisterInterfaces(encoding.InterfaceRegistry)
		irmtypes.RegisterInterfaces(encoding.InterfaceRegistry)

		return encoding
	}
}

// mergeBasicManagers merges the given managers into a single module.BasicManager
func mergeBasicManagers(managers []module.BasicManager) module.BasicManager {
	var union = module.BasicManager{}
	for _, manager := range managers {
		for k, v := range manager {
			union[k] = v
		}
	}
	return union
}
