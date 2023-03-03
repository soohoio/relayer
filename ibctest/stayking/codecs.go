package stayking

import (
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	rlystayking "github.com/cosmos/relayer/v2/relayer/chains/cosmos/stayking"
	"github.com/strangelove-ventures/ibctest/v5/chain/cosmos"
)

func Encoding() *simappparams.EncodingConfig {
	cfg := cosmos.DefaultEncoding()

	rlystayking.RegisterInterfaces(cfg.InterfaceRegistry)

	return &cfg
}
