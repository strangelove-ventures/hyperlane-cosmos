package keeper

import (
	"github.com/cosmos/cosmos-sdk/types"
)

func quoteGasPayment(destination uint32, destinationGasAmount types.Int) types.Int {
	//TODO: Implement spot price
	//TODO: Implement overhead lookup
	return types.ZeroInt()
}

// Callers will specify how much DESTINATION gas they will pay in the MsgPayForGas (PayForGas function).
// This amount needs to be translated into the local chain's native denom, so the caller can pay the relayer.
func getGasPaymentDenom() string {
	//TODO: What is the easiest way to look this up?
	return "stake"
}
