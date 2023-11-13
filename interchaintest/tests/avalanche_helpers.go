package tests

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
)

func SendFunds(
	ec ethclient.Client,
	ctx context.Context,
	senderPrivKey *ecdsa.PrivateKey,
	amount ibc.WalletAmount) (*types.Transaction, error) {
	senderAddr := crypto.PubkeyToAddress(senderPrivKey.PublicKey)

	chainID, err := ec.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get chainID: %w", err)
	}

	senderNonce, err := ec.AcceptedNonceAt(ctx, senderAddr)
	if err != nil {
		return nil, fmt.Errorf("can't get nonce: %w", err)
	}

	gasPrice, err := ec.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get gas price: %w", err)
	}

	toAddress := common.HexToAddress(amount.Address)
	utx := types.NewTransaction(senderNonce, toAddress, big.NewInt(amount.Amount), 21000, gasPrice, nil)
	signedTx, err := types.SignTx(utx, types.NewEIP155Signer(chainID), senderPrivKey)
	if err != nil {
		return nil, fmt.Errorf("can't sign transaction: %w", err)
	}

	return signedTx, ec.SendTransaction(ctx, signedTx)
}
