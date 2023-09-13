package legacy_multisig

import (
	"fmt"
	"reflect"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

var _ types.AbstractIsm = (*LegacyMultiSig)(nil)

func (i *LegacyMultiSig) Event(origin uint32) sdk.Event {
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

func (i *LegacyMultiSig) Validate() error {
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

func (i *LegacyMultiSig) Verify(metadata []byte, message []byte) (bool, error) {
	if !VerifyMerkleProof(metadata, message) {
		return false, types.ErrVerifyMerkleProof
	}
	if !i.VerifyValidatorSignatures(metadata, message) {
		return false, types.ErrVerifyValidatorSignatures
	}

	return true, nil
}

func (i *LegacyMultiSig) VerifyValidatorSignatures(metadata []byte, message []byte) bool {
	if i.Threshold == 0 {
		return false
	}

	// checkpoint digest
	digest := Digest(common.Origin(message), OriginMailbox(metadata), Root(metadata), Index(metadata))

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
		for validatorIndex < validatorCount {
			fmt.Println("Signer: ", hexutil.Encode(signerAddress[:]))
			fmt.Println("Val: ", i.ValidatorPubKeys[validatorIndex])
			valAddress, err := hexutil.Decode(i.ValidatorPubKeys[validatorIndex])
			if err != nil {
				return false
			}
			if reflect.DeepEqual(signerAddress.Bytes(), valAddress) {
				break
			}
			validatorIndex++
		}
		// Fail if we never found a match
		if validatorIndex >= validatorCount {
			return false
		}
		validatorIndex++
	}
	return true
}

func VerifyMerkleProof(metadata []byte, message []byte) bool {
	proof := Proof(metadata)
	paths := [imt.TreeDepth][]byte{}
	for i := 0; i < imt.TreeDepth; i++ {
		paths[i] = proof[i*32 : (i+1)*32]
	}

	calculatedRoot, err := imt.BranchRoot(
		common.Id(message),
		paths,
		common.Nonce(message),
	)
	if err != nil {
		return false
	}

	return reflect.DeepEqual(calculatedRoot, Root(metadata))
}
