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
	errorof "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddRemoteTokenMessenger = "add_remote_token_messenger"

var _ sdk.Msg = &MsgAddRemoteTokenMessenger{}

func NewMsgAddRemoteTokenMessenger(from string, domainId uint32, address []byte) *MsgAddRemoteTokenMessenger {
	return &MsgAddRemoteTokenMessenger{
		From:     from,
		DomainId: domainId,
		Address:  address,
	}
}

func (msg *MsgAddRemoteTokenMessenger) Route() string {
	return RouterKey
}

func (msg *MsgAddRemoteTokenMessenger) Type() string {
	return TypeMsgAddRemoteTokenMessenger
}

func (msg *MsgAddRemoteTokenMessenger) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgAddRemoteTokenMessenger) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRemoteTokenMessenger) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorof.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	return nil
}
