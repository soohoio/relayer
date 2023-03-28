package stayking

import (
	"github.com/cometbft/cometbft/proto/tendermint/crypto"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
)

// Needed for cosmos sdk Msg implementation in messages.go.

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitQueryResponse{}, "/stayking.interchainquery.MsgSubmitQueryResponse", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitQueryResponse{},
	)
}

func RegisterCometTypes() {
	proto.RegisterType((*crypto.Proof)(nil), "cometbft.crypto.Proof")
	proto.RegisterType((*crypto.ValueOp)(nil), "cometbft.crypto.ValueOp")
	proto.RegisterType((*crypto.DominoOp)(nil), "cometbft.crypto.DominoOp")
	proto.RegisterType((*crypto.ProofOp)(nil), "cometbft.crypto.ProofOp")
	proto.RegisterType((*crypto.ProofOps)(nil), "cometbft.crypto.ProofOps")
}

func init() {
	RegisterCometTypes()
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
