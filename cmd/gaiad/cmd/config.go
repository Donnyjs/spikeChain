package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gaia/v8/app/params"
)

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := params.AccountAddressPrefix + "pub"
	validatorAddressPrefix := params.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := params.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := params.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := params.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(params.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
}
