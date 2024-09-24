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
	cdc.RegisterConcrete(&MsgAcceptOwner{}, "cctp.MsgAcceptOwner", nil)
	cdc.RegisterConcrete(&MsgAddRemoteTokenMessenger{}, "cctp.MsgAddRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgDepositForBurn{}, "cctp.MsgDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgDepositForBurnWithCaller{}, "cctp.MsgDepositForBurnWithCaller", nil)
	cdc.RegisterConcrete(&MsgDisableAttester{}, "cctp.MsgDisableAttester", nil)
	cdc.RegisterConcrete(&MsgEnableAttester{}, "cctp.MsgEnableAttester", nil)
	cdc.RegisterConcrete(&MsgLinkTokenPair{}, "cctp.MsgLinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgPauseBurningAndMinting{}, "cctp.MsgPauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgPauseSendingAndReceivingMessages{}, "cctp.MsgPauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgReceiveMessage{}, "cctp.MsgReceiveMessage", nil)
	cdc.RegisterConcrete(&MsgRemoveRemoteTokenMessenger{}, "cctp.MsgRemoveRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgReplaceDepositForBurn{}, "cctp.MsgReplaceDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgReplaceMessage{}, "cctp.MsgReplaceMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessage{}, "cctp.MsgSendMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessageWithCaller{}, "cctp.MsgSendMessageWithCaller", nil)
	cdc.RegisterConcrete(&MsgUnlinkTokenPair{}, "cctp.MsgUnlinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgUnpauseBurningAndMinting{}, "cctp.MsgUnpauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgUnpauseSendingAndReceivingMessages{}, "cctp.MsgUnpauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "cctp.MsgUpdateOwner", nil)
	cdc.RegisterConcrete(&MsgUpdateAttesterManager{}, "cctp.MsgUpdateAttesterManager", nil)
	cdc.RegisterConcrete(&MsgUpdateTokenController{}, "cctp.MsgUpdateTokenController", nil)
	cdc.RegisterConcrete(&MsgUpdatePauser{}, "cctp.MsgUpdatePauser", nil)
	cdc.RegisterConcrete(&MsgUpdateMaxMessageBodySize{}, "cctp.MsgUpdateMaxMessageBodySize", nil)
	cdc.RegisterConcrete(&MsgSetMaxBurnAmountPerMessage{}, "cctp.MsgSetMaxBurnAmountPerMessage", nil)
	cdc.RegisterConcrete(&MsgUpdateSignatureThreshold{}, "cctp.MsgUpdateSignatureThreshold", nil)
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
