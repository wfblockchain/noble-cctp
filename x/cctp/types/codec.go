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
	cdc.RegisterConcrete(&MsgAcceptOwner{}, "circle.cctp.v1.MsgAcceptOwner", nil)
	cdc.RegisterConcrete(&MsgAddRemoteTokenMessenger{}, "circle.cctp.v1.MsgAddRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgDepositForBurn{}, "circle.cctp.v1.MsgDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgDepositForBurnWithCaller{}, "circle.cctp.v1.MsgDepositForBurnWithCaller", nil)
	cdc.RegisterConcrete(&MsgDisableAttester{}, "circle.cctp.v1.MsgDisableAttester", nil)
	cdc.RegisterConcrete(&MsgEnableAttester{}, "circle.cctp.v1.MsgEnableAttester", nil)
	cdc.RegisterConcrete(&MsgLinkTokenPair{}, "circle.cctp.v1.MsgLinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgPauseBurningAndMinting{}, "circle.cctp.v1.MsgPauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgPauseSendingAndReceivingMessages{}, "circle.cctp.v1.MsgPauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgReceiveMessage{}, "circle.cctp.v1.MsgReceiveMessage", nil)
	cdc.RegisterConcrete(&MsgRemoveRemoteTokenMessenger{}, "circle.cctp.v1.MsgRemoveRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgReplaceDepositForBurn{}, "circle.cctp.v1.MsgReplaceDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgReplaceMessage{}, "circle.cctp.v1.MsgReplaceMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessage{}, "circle.cctp.v1.MsgSendMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessageWithCaller{}, "circle.cctp.v1.MsgSendMessageWithCaller", nil)
	cdc.RegisterConcrete(&MsgUnlinkTokenPair{}, "circle.cctp.v1.MsgUnlinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgUnpauseBurningAndMinting{}, "circle.cctp.v1.MsgUnpauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgUnpauseSendingAndReceivingMessages{}, "circle.cctp.v1.MsgUnpauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "circle.cctp.v1.MsgUpdateOwner", nil)
	cdc.RegisterConcrete(&MsgUpdateAttesterManager{}, "circle.cctp.v1.MsgUpdateAttesterManager", nil)
	cdc.RegisterConcrete(&MsgUpdateTokenController{}, "circle.cctp.v1.MsgUpdateTokenController", nil)
	cdc.RegisterConcrete(&MsgUpdatePauser{}, "circle.cctp.v1.MsgUpdatePauser", nil)
	cdc.RegisterConcrete(&MsgUpdateMaxMessageBodySize{}, "circle.cctp.v1.MsgUpdateMaxMessageBodySize", nil)
	cdc.RegisterConcrete(&MsgSetMaxBurnAmountPerMessage{}, "circle.cctp.v1.MsgSetMaxBurnAmountPerMessage", nil)
	cdc.RegisterConcrete(&MsgUpdateSignatureThreshold{}, "circle.cctp.v1.MsgUpdateSignatureThreshold", nil)
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
