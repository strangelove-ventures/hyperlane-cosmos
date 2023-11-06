package keeper

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

type QueryMsg struct {
	QueryIsmMsg *QueryIsmMsg `json:"ism,omitempty"`
}

type QueryIsmMsg struct{}

type QueryIsmResponse struct {
	IsmId uint32 `json:"ism_id,omitempty"`
}

type ContractMsg struct {
	ContractProcessMsg ContractProcessMsg `json:"process_msg,omitempty"`
}

type ContractProcessMsg struct {
	Origin uint32 `json:"origin"`
	Sender string `json:"sender"`
	Msg    string `json:"msg"`
}

func (k Keeper) getReceiversIsm(ctx sdk.Context, recipient string) (uint32, error) {
	// Iterate through all receivers
	for _, receiver := range k.receivers {
		// If recipient matches the receiver address, query its ISM
		if recipient == receiver.Address() {
			return receiver.QueryIsm(), nil
		}
	}

	// If x/wasm is supported and no receivers match the recipient, query the contract
	if k.cwKeeper != nil {
		recipientBz, err := sdk.AccAddressFromBech32(recipient)
		if err != nil {
			return 0, err
		}

		req := QueryMsg{QueryIsmMsg: &QueryIsmMsg{}}
		reqBz, err := json.Marshal(req)
		if err != nil {
			return 0, err
		}
		resp, err := k.cwKeeper.QuerySmart(ctx, recipientBz, reqBz)
		if err != nil {
			return 0, err
		}

		var ismResp QueryIsmResponse
		err = json.Unmarshal(resp, &ismResp)
		if err != nil {
			return 0, err
		}

		return ismResp.IsmId, nil
	}

	return 0, types.ErrInvalidRecipient
}

func (k Keeper) processMsg(ctx sdk.Context, recipient string, origin uint32, sender string, msg string) error {
	// Iterate through all receivers
	for _, receiver := range k.receivers {
		// If recipient matches the receiver address, process the message
		if recipient == receiver.Address() {
			err := receiver.Process(origin, sender, msg)
			return err
		}
	}

	// If x/wasm is supported and no receivers match the recipient, call the contract's process
	if k.cwKeeper != nil {
		contractMsg := ContractMsg{
			ContractProcessMsg: ContractProcessMsg{
				Origin: origin,
				Sender: sender,
				Msg:    msg,
			},
		}
		encodedMsg, err := json.Marshal(contractMsg)
		if err != nil {
			return err
		}

		receiverAddr, err := sdk.AccAddressFromBech32(recipient)
		if err != nil {
			return errorsmod.Wrap(err, "invalid bech32 recipient")
		}

		// Call the recipient contract
		_, err = k.pcwKeeper.Execute(ctx, receiverAddr, k.mailboxAddr, encodedMsg, sdk.NewCoins())
		if err != nil {
			return err
		}

		return nil
	}

	return types.ErrInvalidRecipient
}
