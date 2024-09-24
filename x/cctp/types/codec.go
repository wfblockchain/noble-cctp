/*
 * Copyright (c) 2023, © Circle Internet Financial, LTD.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAcceptOwner{}, "/cctp.Msg/AcceptOwner", nil)
	cdc.RegisterConcrete(&MsgAddRemoteTokenMessenger{}, "/cctp.Msg/AddRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgDepositForBurn{}, "/cctp.Msg/DepositForBurn", nil)
	cdc.RegisterConcrete(&MsgDepositForBurnWithCaller{}, "/cctp.Msg/DepositForBurnWithCaller", nil)
	cdc.RegisterConcrete(&MsgDisableAttester{}, "/cctp.Msg/DisableAttester", nil)
	cdc.RegisterConcrete(&MsgEnableAttester{}, "/cctp.Msg/EnableAttester", nil)
	cdc.RegisterConcrete(&MsgLinkTokenPair{}, "/cctp.Msg/LinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgPauseBurningAndMinting{}, "/cctp.Msg/PauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgPauseSendingAndReceivingMessages{}, "/cctp.Msg/PauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgReceiveMessage{}, "/cctp.Msg/ReceiveMessage", nil)
	cdc.RegisterConcrete(&MsgRemoveRemoteTokenMessenger{}, "/cctp.Msg/RemoveRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgReplaceDepositForBurn{}, "/cctp.Msg/ReplaceDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgReplaceMessage{}, "/cctp.Msg/ReplaceMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessage{}, "/cctp.Msg/SendMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessageWithCaller{}, "/cctp.Msg/SendMessageWithCaller", nil)
	cdc.RegisterConcrete(&MsgUnlinkTokenPair{}, "/cctp.Msg/UnlinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgUnpauseBurningAndMinting{}, "/cctp.Msg/UnpauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgUnpauseSendingAndReceivingMessages{}, "/cctp.Msg/UnpauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "/cctp.Msg/UpdateOwner", nil)
	cdc.RegisterConcrete(&MsgUpdateAttesterManager{}, "/cctp.Msg/UpdateAttesterManager", nil)
	cdc.RegisterConcrete(&MsgUpdateTokenController{}, "/cctp.Msg/UpdateTokenController", nil)
	cdc.RegisterConcrete(&MsgUpdatePauser{}, "/cctp.Msg/UpdatePauser", nil)
	cdc.RegisterConcrete(&MsgUpdateMaxMessageBodySize{}, "/cctp.Msg/UpdateMaxMessageBodySize", nil)
	cdc.RegisterConcrete(&MsgSetMaxBurnAmountPerMessage{}, "/cctp.Msg/SetMaxBurnAmountPerMessage", nil)
	cdc.RegisterConcrete(&MsgUpdateSignatureThreshold{}, "/cctp.Msg/UpdateSignatureThreshold", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAcceptOwner{},
		&MsgAddRemoteTokenMessenger{},
		&MsgDepositForBurn{},
		&MsgDepositForBurnWithCaller{},
		&MsgDisableAttester{},
		&MsgEnableAttester{},
		&MsgLinkTokenPair{},
		&MsgPauseBurningAndMinting{},
		&MsgPauseSendingAndReceivingMessages{},
		&MsgReceiveMessage{},
		&MsgRemoveRemoteTokenMessenger{},
		&MsgReplaceDepositForBurn{},
		&MsgReplaceMessage{},
		&MsgSendMessage{},
		&MsgSendMessageWithCaller{},
		&MsgUnlinkTokenPair{},
		&MsgUnpauseBurningAndMinting{},
		&MsgUnpauseSendingAndReceivingMessages{},
		&MsgUpdateOwner{},
		&MsgUpdateAttesterManager{},
		&MsgUpdateTokenController{},
		&MsgUpdatePauser{},
		&MsgUpdateMaxMessageBodySize{},
		&MsgSetMaxBurnAmountPerMessage{},
		&MsgUpdateSignatureThreshold{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
