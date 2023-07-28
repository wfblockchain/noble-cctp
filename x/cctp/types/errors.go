package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cctp module sentinel errors
var (
	ErrUnauthorized             = sdkerrors.Register(ModuleName, 30, "unauthorized")
	ErrMint                     = sdkerrors.Register(ModuleName, 31, "tokens can not be minted")
	ErrBurn                     = sdkerrors.Register(ModuleName, 32, "tokens can not be burned")
	ErrAttesterAlreadyFound     = sdkerrors.Register(ModuleName, 33, "attester is already present")
	ErrAuthorityNotSet          = sdkerrors.Register(ModuleName, 34, "authority not set")
	ErrMalformedField           = sdkerrors.Register(ModuleName, 35, "field cannot be empty or nil")
	ErrReceiveMessage           = sdkerrors.Register(ModuleName, 36, "err in receive message")
	ErrDisableAttester          = sdkerrors.Register(ModuleName, 37, "err in disable attester")
	ErrUpdateSignatureThreshold = sdkerrors.Register(ModuleName, 38, "err in update signature threshold")
	ErrMinterAllowanceNotFound  = sdkerrors.Register(ModuleName, 39, "minter allowance not found")
	ErrTokenPairAlreadyFound    = sdkerrors.Register(ModuleName, 40, "token pair already exists")
	ErrTokenPairNotFound        = sdkerrors.Register(ModuleName, 41, "token pair not found")
	ErrSendMessage              = sdkerrors.Register(ModuleName, 42, "error in send message")
	ErrSendMessageWithCaller    = sdkerrors.Register(ModuleName, 43, "error in send message with caller")
	ErrDepositForBurn           = sdkerrors.Register(ModuleName, 44, "error in deposit for burn")
	ErrInvalidDestinationCaller = sdkerrors.Register(ModuleName, 45, "malformed destination caller")
	ErrSignatureVerification    = sdkerrors.Register(ModuleName, 46, "unable to verify signature")
	ErrReplaceMessage           = sdkerrors.Register(ModuleName, 47, "error in replace message")
	ErrDuringPause              = sdkerrors.Register(ModuleName, 48, "error while trying to pause or unpause")
	ErrInvalidAmount            = sdkerrors.Register(ModuleName, 49, "invalid amount")
)
