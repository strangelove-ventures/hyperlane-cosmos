package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper expected staking keeper
type StakingKeeper interface {
	BondDenom(ctx sdk.Context) string
}
