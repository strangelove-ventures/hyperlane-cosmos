package keeper_test

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"

	"github.com/golang/mock/gomock"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

// trackMockBalances sets up expected calls on the Mock BankKeeper, and also
// locally tracks accounts balances (not modules balances).
func trackMockBalances(bankKeeper *govtestutil.MockBankKeeper) {
	balances := make(map[string]sdk.Coins)

	// We don't track module account balances.
	bankKeeper.EXPECT().MintCoins(gomock.Any(), minttypes.ModuleName, gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToModule(gomock.Any(), minttypes.ModuleName, types.ModuleName, gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().SendCoins(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	// But we do track normal account balances.
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), types.ModuleName, gomock.Any()).DoAndReturn(func(_ sdk.Context, sender sdk.AccAddress, _ string, coins sdk.Coins) error {
		newBalance, negative := balances[sender.String()].SafeSub(coins...)
		if negative {
			return fmt.Errorf("not enough balance")
		}
		balances[sender.String()] = newBalance
		return nil
	}).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ sdk.Context, module string, rcpt sdk.AccAddress, coins sdk.Coins) error {
		balances[rcpt.String()] = balances[rcpt.String()].Add(coins...)
		return nil
	}).AnyTimes()
	bankKeeper.EXPECT().GetAllBalances(gomock.Any(), gomock.Any()).DoAndReturn(func(_ sdk.Context, addr sdk.AccAddress) sdk.Coins {
		return balances[addr.String()]
	}).AnyTimes()
	bankKeeper.EXPECT().GetBalance(gomock.Any(), gomock.Any(), sdk.DefaultBondDenom).DoAndReturn(func(_ sdk.Context, addr sdk.AccAddress, _ string) sdk.Coin {
		balances := balances[addr.String()]
		for _, balance := range balances {
			if balance.Denom == sdk.DefaultBondDenom {
				return balance
			}
		}
		return sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(0))
	}).AnyTimes()
}
