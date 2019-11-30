package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	client "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
	commitment "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment"
)

// SubModuleCdc defines the IBC channel codec.
var SubModuleCdc *codec.Codec

func init() {
	SubModuleCdc = codec.New()
	commitment.RegisterCodec(SubModuleCdc)
	client.RegisterCodec(SubModuleCdc)
	RegisterCodec(SubModuleCdc)
}

// RegisterCodec registers all the necessary types and interfaces for the
// IBC channel.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*PacketDataI)(nil), nil)
	cdc.RegisterConcrete(Packet{}, "ibc/channel/Packet", nil)

	cdc.RegisterConcrete(MsgChannelOpenInit{}, "ibc/channel/MsgChannelOpenInit", nil)
	cdc.RegisterConcrete(MsgChannelOpenTry{}, "ibc/channel/MsgChannelOpenTry", nil)
	cdc.RegisterConcrete(MsgChannelOpenAck{}, "ibc/channel/MsgChannelOpenAck", nil)
	cdc.RegisterConcrete(MsgChannelOpenConfirm{}, "ibc/channel/MsgChannelOpenConfirm", nil)
	cdc.RegisterConcrete(MsgChannelCloseInit{}, "ibc/channel/MsgChannelCloseInit", nil)
	cdc.RegisterConcrete(MsgChannelCloseConfirm{}, "ibc/channel/MsgChannelCloseConfirm", nil)
}

func SetSubModuleCodec(cdc *codec.Codec) {
	SubModuleCdc = cdc
}
