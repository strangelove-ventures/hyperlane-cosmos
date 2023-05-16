package keeper

import (
	//"encoding/hex"

	"fmt"
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	authority  string
	defaultIsm MultiSigIsm
}

type MultiSigIsm struct {
	ValPubKeys [][]byte
	Threshold  uint32
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, authority string) Keeper {
	return Keeper{
		cdc:       cdc,
		storeKey:  key,
		authority: authority,
	}
}

func (k Keeper) Verify(metadata, message []byte) bool {
	// Look up recipient contract's ISM, if 0, use default multi sig (just use default for now)
	ism := k.defaultIsm
	return VerifyMerkleProof(metadata, message) && VerifyValidatorSignatures(metadata, message, ism)
}

func VerifyMerkleProof(metadata []byte, message []byte) bool {
	proof := types.Proof(metadata)
	paths := [imt.TreeDepth][]byte{}
	for i := 0; i < imt.TreeDepth; i++ {
		paths[i] = proof[i*32 : (i+1)*32]
	}

	calculatedRoot, err := imt.BranchRoot(
		types.Id(message),
		paths,
		types.Nonce(message),
	)
	if err != nil {
		return false
	}

	return reflect.DeepEqual(calculatedRoot, types.Root(metadata))
}

func VerifyValidatorSignatures(metadata []byte, message []byte, ism MultiSigIsm) bool {

	if ism.Threshold == 0 {
		return false
	}

	// checkpoint digest
	digest := types.Digest(types.Origin(message), types.OriginMailbox(metadata), types.Root(metadata), types.Index(metadata))

	validatorCount := len(ism.ValPubKeys)
	validatorIndex := 0
	// Assumes that signatures are ordered by validator
	for i := uint32(0); i < ism.Threshold; i++ {
		// get signer
		signer, err := crypto.SigToPub(digest, types.SignatureAt(metadata, i))
		if err != nil {
			fmt.Println("signer recover error: ", err)
			return false
		}
		//fmt.Println("Signer: ", hex.EncodeToString(signer))
		signerAddress := crypto.PubkeyToAddress(*signer)
		// Loop through remaining validators until we find a match
		for validatorIndex < validatorCount && !reflect.DeepEqual(signerAddress.Bytes(), ism.ValPubKeys[validatorIndex]) {
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
