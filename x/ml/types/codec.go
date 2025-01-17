package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGenerate{}, "ml/Generate", nil)
	cdc.RegisterConcrete(&MsgConditional{}, "ml/Conditional", nil)
	cdc.RegisterConcrete(&MsgLlm{}, "ml/Llm", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenerate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConditional{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLlm{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
