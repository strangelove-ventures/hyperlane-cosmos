package keeper

import (
	"context"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// Dispatch defines a rpc handler method for MsgDispatch
func (k Keeper) Dispatch(goCtx context.Context, msg *types.MsgDispatch) (*types.MsgDispatchResponse, error) {
	panic("Implement me")
}

// Process defines a rpc handler method for MsgProcess
func (k Keeper) Process(goCtx context.Context, msg *types.MsgProcess) (*types.MsgProcessResponse, error) {
	panic("Implement me")
}

// StoreCode defines a rpc handler method for MsgStoreCode
func (k Keeper) SetDefaultIsm(goCtx context.Context, msg *types.MsgSetDefaultIsm) (*types.MsgSetDefaultIsmResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Signer {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority: expected %s, got %s", k.authority, msg.Signer)
	}

	panic("Implement me")

	/*ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeSetDefaultIsm,
			sdk.NewAttribute(AttributeKeySetDefaultIsm, hex.EncodeToString(ismHash)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		),
	})

	return &types.MsgSetDefaultIsmResponse{}, nil*/
}
