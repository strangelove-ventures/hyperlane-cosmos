package types

import (
	proto "github.com/cosmos/gogoproto/proto"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	MerkleRootMultiSigType string = "merkle-root-multisig"
	MessageIdMultiSigType  string = "message-id-multisig"
)

type AbstractIsm interface {
	proto.Message

	IsmType() string
	Validate() error
	Verify(metadata []byte, message []byte) bool
	Event(origin uint32) sdk.Event
}
