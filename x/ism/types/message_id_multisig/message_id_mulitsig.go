package message_id_multisig

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

var _ types.AbstractIsm = (*MessageIdMultiSig)(nil)

func (i *MessageIdMultiSig) Event(origin uint32) sdk.Event {
	originStr := strconv.FormatUint(uint64(origin), 10)
	thresholdStr := strconv.FormatUint(uint64(i.Threshold), 10)
	eventAttributes := []sdk.Attribute{}
	eventAttributes = append(eventAttributes, sdk.NewAttribute(types.AttributeKeyOrigin, originStr))
	eventAttributes = append(eventAttributes, sdk.NewAttribute(types.AttributeKeyThreshold, thresholdStr))
	for index := 0; index < len(i.ValidatorPubKeys); index++ {
		eventAttributes = append(eventAttributes, sdk.NewAttribute(
			types.AttributeKeyValidator,
			i.ValidatorPubKeys[index],
		))
	}
	return sdk.NewEvent(
		types.EventTypeSetDefaultIsm,
		eventAttributes...,
	)
}

func (i *MessageIdMultiSig) Validate() error {
	if i.Threshold == 0 {
		return types.ErrInvalidThreshold
	}
	for _, validator := range i.ValidatorPubKeys {
		len := len(validator)
		if len < 42 || len > 66 { // Will be 21-66 bytes
			return types.ErrInvalidValSet
		}
	}
	return nil
}

func (i *MessageIdMultiSig) Verify(metadata []byte, message []byte) bool {
	return i.VerifyValidatorSignatures(metadata, message)
}

func (i *MessageIdMultiSig) VerifyValidatorSignatures(metadata []byte, message []byte) bool {
	if i.Threshold == 0 {
		return false
	}

	// checkpoint digest
	digest := Digest(common.Origin(message), OriginMailbox(metadata), 
						Root(metadata), common.Nonce(message), common.Id(message))

	validatorCount := len(i.ValidatorPubKeys)
	validatorIndex := 0
	// Assumes that signatures are ordered by validator
	for index := uint32(0); index < i.Threshold; index++ {
		// get signer
		signer, err := crypto.SigToPub(digest, SignatureAt(metadata, index))
		if err != nil {
			return false
		}
		// fmt.Println("Signer: ", hex.EncodeToString(signer))
		signerAddress := crypto.PubkeyToAddress(*signer)
		// Loop through remaining validators until we find a match
		for validatorIndex < validatorCount &&
			hexutil.Encode(signerAddress.Bytes()) == i.ValidatorPubKeys[validatorIndex] {
			validatorIndex++
		}
		// Fail if we never found a match
		if validatorIndex >= validatorCount {
			fmt.Println("never found match")
			return false
		}
		validatorIndex++
	}
	return true
}
