package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
)

type AbstractIsm interface {
	proto.Message

	Validate() error
	Verify(metadata []byte, message []byte) bool
	Event(origin uint32) sdk.Event
}
