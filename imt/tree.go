package imt

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	TreeDepth = 32
	MaxLeaves = ^uint32(0)
)

type Tree struct {
	Branch [TreeDepth][]byte
	count  uint32
}

// Insert inserts node into the Merkle Tree
func (t *Tree) Insert(node []byte) ([32][]byte, error) {
	if t.count >= MaxLeaves {
		return [32][]byte{}, errors.New("merkle tree full")
	}

	if len(node) != 32 {
		return [32][]byte{}, errors.New("must be 32-bytes")
	}

	t.count += 1
	size := t.count
	for i := 0; i < TreeDepth; i++ {
		if (size & 1) == 1 {
			t.Branch[i] = node
			return t.Branch, nil
		}
		temp := append(t.Branch[i][:], node...)
		node = crypto.Keccak256(temp)
		size /= 2
	}

	return [32][]byte{}, errors.New("unreachable")
}

// Count returns the number of inserts performed on the Tree
func (t *Tree) Count() uint32 {
	return t.count
}

// Print dumps the tree (for debugging)
func (t *Tree) Print() {
	for i := 0; i < TreeDepth; i++ {
		fmt.Printf("%02d: %X\n", i, t.Branch[i])
	}
}

// Root calculates and returns Tree's current root
func (t *Tree) Root() []byte {
	return t.RootWithContext(ZeroHashes())
}

// RootWithCtx calculates and returns Tree's current root given array of zeroes
func (t *Tree) RootWithContext(zeroes [][]byte) []byte {
	index := t.count

	// We start with a 32-bit zero slice
	current := make([]byte, 32)
	for i := 0; i < TreeDepth; i++ {
		ithBit := (index >> i) & 0x01
		next := t.Branch[i]

		var temp []byte
		if ithBit == 1 {
			temp = append(next, current...)
		} else {
			temp = append(current, zeroes[i]...)
		}
		current = crypto.Keccak256(temp)
	}
	return current
}

// BranchRoot calculates and returns the merkle root for the given leaf item, a merkle branch, and the index of item in the tree
func BranchRoot(item []byte, branch [TreeDepth][]byte, index uint32) ([]byte, error) {
	if len(item) != 32 {
		return nil, errors.New("must be 32-bytes")
	}

	current := make([]byte, 32)
	copy(current, item)
	for i := 0; i < TreeDepth; i++ {
		ithBit := (index >> i) & 0x01
		next := make([]byte, 32)
		copy(next, branch[i])

		var temp []byte
		if ithBit == 1 {
			temp = append(next, current...)
		} else {
			temp = append(current, next...)
		}
		current = crypto.Keccak256(temp)
	}
	return current, nil
}

// GetProofForNextIndex returns the proof for the next index.
func (t *Tree) GetProofForNextIndex() (proof [TreeDepth][32]byte) {
	zeroHashes := ZeroHashes()
	index := t.count
	for i := 0; i < TreeDepth; i++ {
		ithBit := (index >> i) & 0x01
		if ithBit == 1 {
			copy(proof[i][:], t.Branch[i])
		} else {
			copy(proof[i][:], zeroHashes[i])
		}
	}
	return proof
}

// ZeroHashes returns the array of TreeDepth zero hashes
func ZeroHashes() [][]byte {
	zeroes := [][]byte{
		common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000"),
		common.FromHex("0xad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5"),
		common.FromHex("0xb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30"),
		common.FromHex("0x21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85"),
		common.FromHex("0xe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a19344"),
		common.FromHex("0x0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d"),
		common.FromHex("0x887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a1968"),
		common.FromHex("0xffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f83"),
		common.FromHex("0x9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af"),
		common.FromHex("0xcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e0"),
		common.FromHex("0xf9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a5"),
		common.FromHex("0xf8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf892"),
		common.FromHex("0x3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c"),
		common.FromHex("0xc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb"),
		common.FromHex("0x5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc"),
		common.FromHex("0xda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d2"),
		common.FromHex("0x2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f"),
		common.FromHex("0xe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a"),
		common.FromHex("0x5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0"),
		common.FromHex("0xb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0"),
		common.FromHex("0xc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2"),
		common.FromHex("0xf4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd9"),
		common.FromHex("0x5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e377"),
		common.FromHex("0x4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652"),
		common.FromHex("0xcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef"),
		common.FromHex("0x0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d"),
		common.FromHex("0xb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0"),
		common.FromHex("0x838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e"),
		common.FromHex("0x662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e"),
		common.FromHex("0x388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea322"),
		common.FromHex("0x93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735"),
		common.FromHex("0x8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9"),
	}

	return zeroes
}
