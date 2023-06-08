package types

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	legacy "github.com/strangelove-ventures/hyperlane-cosmos/x/common_legacy"
)

var _ AbstractIsm = (*MerkleRootMultiSig)(nil)

func (i *MerkleRootMultiSig) IsmType() string {
	return MerkleRootMultiSigType
}

func (i *MerkleRootMultiSig) Event(origin uint32) sdk.Event {
	originStr := strconv.FormatUint(uint64(origin), 10)
	thresholdStr := strconv.FormatUint(uint64(i.Threshold), 10)
	eventAttributes := []sdk.Attribute{}
	eventAttributes = append(eventAttributes, sdk.NewAttribute(AttributeKeyOrigin, originStr))
	eventAttributes = append(eventAttributes, sdk.NewAttribute(AttributeKeyThreshold, thresholdStr))
	for index := 0; index < len(i.ValidatorPubKeys); index++ {
		eventAttributes = append(eventAttributes, sdk.NewAttribute(
			AttributeKeyValidator,
			i.ValidatorPubKeys[index],
		))
	}
	return sdk.NewEvent(
		EventTypeSetDefaultIsm,
		eventAttributes...
	)
}

func (i *MerkleRootMultiSig) Validate() error {
	if i.Threshold == 0 {
		return ErrInvalidThreshold
	}
	for _, validator := range i.ValidatorPubKeys {
		len := len(validator)
		if len < 42 || len > 66 { // Will be 21-66 bytes
			return ErrInvalidValSet
		}
	}
	return nil
}

func (i *MerkleRootMultiSig) Verify(metadata []byte, message []byte) bool {
	return VerifyMerkleProof(metadata, message) && i.VerifyValidatorSignatures(metadata, message)
}

func (i *MerkleRootMultiSig) VerifyValidatorSignatures(metadata []byte, message []byte) bool {
	if i.Threshold == 0 {
		return false
	}

	// checkpoint digest
	digest := legacy.Digest(common.Origin(message), legacy.OriginMailbox(metadata), legacy.Root(metadata), legacy.Index(metadata))

	validatorCount := len(i.ValidatorPubKeys)
	validatorIndex := 0
	// Assumes that signatures are ordered by validator
	for index := uint32(0); index < i.Threshold; index++ {
		// get signer
		signer, err := crypto.SigToPub(digest, legacy.SignatureAt(metadata, index))
		if err != nil {
			fmt.Println("signer recover error: ", err)
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

func VerifyMerkleProof(metadata []byte, message []byte) bool {
	proof := legacy.Proof(metadata)
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

	return reflect.DeepEqual(calculatedRoot, legacy.Root(metadata))
}
