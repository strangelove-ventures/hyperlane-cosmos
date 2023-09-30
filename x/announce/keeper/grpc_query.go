package keeper

import (
	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

var _ types.QueryServer = (*Keeper)(nil)
