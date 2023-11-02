package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
			fmt.Println("Receiver ISM ID response: ", receiver.QueryIsm())
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

		fmt.Println("Contract ISM ID response: ", ismResp.IsmId)
		return ismResp.IsmId, nil
	}

	return 0, sdkerrors.ErrInvalidAddress
}

func (k Keeper) processMsg(ctx sdk.Context, recipient string, origin uint32, sender string, msg string) error {
	// Iterate through all receivers
	for _, receiver := range k.receivers {
		// If recipient matches the receiver address, process the message
		if recipient == receiver.Address() {
			fmt.Println("Receiver process")
			receiver.Process(origin, sender, msg)
			return nil
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
			return sdkerrors.Wrap(err, "contract")
		}

		// Call the recipient contract
		_, err = k.pcwKeeper.Execute(ctx, receiverAddr, k.mailboxAddr, encodedMsg, sdk.NewCoins())
		if err != nil {
			fmt.Println("Contract err: ", err) // TODO: remove, debug only
			return err
		}

		fmt.Println("Contract process")
		return nil
	}

	return sdkerrors.ErrInvalidAddress
}
