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
	cdc.RegisterConcrete(&MsgAcceptOwner{}, "circle.cctp.v1/AcceptOwner", nil)
	cdc.RegisterConcrete(&MsgAddRemoteTokenMessenger{}, "circle.cctp.v1/AddRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgDepositForBurn{}, "circle.cctp.v1/DepositForBurn", nil)
	cdc.RegisterConcrete(&MsgDepositForBurnWithCaller{}, "circle.cctp.v1/DepositForBurnWithCaller", nil)
	cdc.RegisterConcrete(&MsgDisableAttester{}, "circle.cctp.v1/DisableAttester", nil)
	cdc.RegisterConcrete(&MsgEnableAttester{}, "circle.cctp.v1/EnableAttester", nil)
	cdc.RegisterConcrete(&MsgLinkTokenPair{}, "circle.cctp.v1/LinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgPauseBurningAndMinting{}, "circle.cctp.v1/PauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgPauseSendingAndReceivingMessages{}, "circle.cctp.v1/PauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgReceiveMessage{}, "circle.cctp.v1/ReceiveMessage", nil)
	cdc.RegisterConcrete(&MsgRemoveRemoteTokenMessenger{}, "circle.cctp.v1/RemoveRemoteTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgReplaceDepositForBurn{}, "circle.cctp.v1/ReplaceDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgReplaceMessage{}, "circle.cctp.v1/ReplaceMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessage{}, "circle.cctp.v1/SendMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessageWithCaller{}, "circle.cctp.v1/SendMessageWithCaller", nil)
	cdc.RegisterConcrete(&MsgUnlinkTokenPair{}, "circle.cctp.v1/UnlinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgUnpauseBurningAndMinting{}, "circle.cctp.v1/UnpauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgUnpauseSendingAndReceivingMessages{}, "circle.cctp.v1/UnpauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "circle.cctp.v1/UpdateOwner", nil)
	cdc.RegisterConcrete(&MsgUpdateAttesterManager{}, "circle.cctp.v1/UpdateAttesterManager", nil)
	cdc.RegisterConcrete(&MsgUpdateTokenController{}, "circle.cctp.v1/UpdateTokenController", nil)
	cdc.RegisterConcrete(&MsgUpdatePauser{}, "circle.cctp.v1/UpdatePauser", nil)
	cdc.RegisterConcrete(&MsgUpdateMaxMessageBodySize{}, "circle.cctp.v1/UpdateMaxMessageBodySize", nil)
	cdc.RegisterConcrete(&MsgSetMaxBurnAmountPerMessage{}, "circle.cctp.v1/SetMaxBurnAmountPerMessage", nil)
	cdc.RegisterConcrete(&MsgUpdateSignatureThreshold{}, "circle.cctp.v1/UpdateSignatureThreshold", nil)
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
