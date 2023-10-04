package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/crypto"
)

// This is an ECDSA public key (secp256k1) in uncompressed format (65 bytes).
// However, in Ethereum the leading byte is omitted, so 64 bytes total.
// The Cosmos hyperlane module is compatible with the same format.
const (
	ETHEREUM_PUB_KEY_LEN = 64
	ETHEREUM_ADDR_LEN    = 20 // The ethereum address format is the last 20 bytes of the Keccak256 hashed public key
)

func encodePackedAnnouncement(origin uint32, originMailbox []byte) ([]byte, error) {
	if len(originMailbox) != 32 {
		return nil, sdkerrors.Wrapf(ErrPackAnnouncement, "origin mailbox is %d bytes, expected 32", len(originMailbox))
	}

	var packed []byte
	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, origin)
	packed = append(packed, originBytes...)
	packed = append(packed, originMailbox...)
	packed = append(packed, []byte("HYPERLANE_ANNOUNCEMENT")...)
	return packed, nil
}

func hash(packed []byte) []byte {
	return crypto.Keccak256(packed)
}

func getAddress(pubKey []byte) ([]byte, error) {
	if len(pubKey) != ETHEREUM_PUB_KEY_LEN {
		return nil, fmt.Errorf("provided bytes %s is not a valid public key", hex.EncodeToString(pubKey))
	}
	hashedKey := hash(pubKey)
	return hashedKey[len(hashedKey)-20:], nil
}

func getAnnouncementDigest(origin uint32, originMailbox []byte, storageLocation string) ([]byte, error) {
	pack, err := encodePackedAnnouncement(origin, originMailbox)
	if err != nil {
		return nil, err
	}
	domainHash := hash(pack)
	var packedDomainHashStorageLoc []byte
	packedDomainHashStorageLoc = append(packedDomainHashStorageLoc, domainHash...)
	packedDomainHashStorageLoc = append(packedDomainHashStorageLoc, []byte(storageLocation)...)
	return toEthSignedMessageHash(hash(packedDomainHashStorageLoc)), nil
}

func toEthSignedMessageHash(packedHash []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(packedHash), packedHash)
	return hash([]byte(msg))
}

func verifyAnnouncementDigest(digest []byte, signature []byte, expectedSigner []byte) error {
	sigPublicKey, err := crypto.Ecrecover(digest, signature)
	if err != nil {
		return err
	}

	addr, err := getAddress(sigPublicKey)
	if err != nil {
		return err
	}

	if bytes.Equal(expectedSigner, addr) {
		return nil
	}

	return errors.New("Signature does not match the declared validator")
}
