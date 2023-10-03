package imt_test

import (
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
)

type MerkleProof struct {
	Leaf         string
	Index        uint32
	Path         []string
	ExpectedRoot string
}

type MerkleVector struct {
	TestName     string
	ExpectedRoot string
	Leaves       []string
	Proofs       []MerkleProof
}

//go:embed testdata/merkle.json
var merkleJSON []byte

//go:embed testdata/incremental_merkle.json
var incMerkleJSON []byte

func TestFootGuns(t *testing.T) {
	var i imt.Tree

	emptySlice := []byte{}
	err := i.Insert(emptySlice)
	require.Error(t, err, "nodes must be 32-bytes")

	zeroes_31 := common.FromHex("0x00000000000000000000000000000000000000000000000000000000000000")
	err = i.Insert(zeroes_31)
	require.Error(t, err, "nodes must be 32-bytes")

	zeroes_33 := common.FromHex("0x000000000000000000000000000000000000000000000000000000000000000000")
	err = i.Insert(zeroes_33)
	require.Error(t, err, "nodes must be 32-bytes")
}

func TestVectors(t *testing.T) {
	var cases []MerkleVector

	// Open the test cases
	// Read them in
	err := json.Unmarshal(merkleJSON, &cases)
	require.NoError(t, err)

	// Test them
	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing case: %s", c.TestName), func(t *testing.T) {
			i := imt.Tree{}

			// Insert all of the leaves
			for idx, l := range c.Leaves {
				// Test Vectors used 'ethers.utils.hashMessage(leaf)'
				// to compute the value to insert.
				msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(l), l)
				hash := crypto.Keccak256([]byte(msg))

				// Make sure we get the expected digest
				expectedLeaf, err := hex.DecodeString(c.Proofs[idx].Leaf[2:])
				require.NoError(t, err)
				require.Equal(t, hash[:], expectedLeaf)

				// Insert into the tree
				err = i.Insert(hash)
				require.NoError(t, err)

			}

			// Make sure we've inserted the correct amount
			require.Equal(t, int(i.Count()), len(c.Leaves))

			// Make sure we've computed the expected root
			expectedRoot, err := hex.DecodeString(c.ExpectedRoot[2:])
			require.NoError(t, err)

			r := i.Root()
			require.Equal(t, r[:], expectedRoot)

			// Verify leaves
			for _, p := range c.Proofs {
				leaf, err := hex.DecodeString(p.Leaf[2:])
				require.NoError(t, err)

				paths := [imt.TreeDepth][]byte{}
				for idx, path := range p.Path {
					pBytes, err := hex.DecodeString(path[2:])
					require.NoError(t, err)
					paths[idx] = pBytes
				}

				// Make sure we get the expected branch root
				proofRoot, err := imt.BranchRoot(leaf, paths, p.Index)
				require.NoError(t, err)
				require.Equal(t, proofRoot[:], expectedRoot)
			}
		})
	}
}

func TestIncrementalVectors(t *testing.T) {
	var cases []MerkleVector

	err := json.Unmarshal(incMerkleJSON, &cases)
	require.NoError(t, err)

	// Test them
	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing case: %s", c.TestName), func(t *testing.T) {
			// Verify leaves
			for idx, p := range c.Proofs {
				i := imt.Tree{}
				for index := 0; index <= idx; index++ {
					hash := crypto.Keccak256([]byte(c.Leaves[index]))
					// Make sure we get the expected digest
					expectedLeaf, err := hex.DecodeString(c.Proofs[index].Leaf[2:])
					require.NoError(t, err)
					require.Equal(t, hash[:], expectedLeaf)

					// Insert into the tree
					err = i.Insert(hash)
					require.NoError(t, err)
				}

				// Make sure we've inserted the correct amount
				require.Equal(t, i.Count(), uint32(idx+1))

				leaf, err := hex.DecodeString(p.Leaf[2:])
				require.NoError(t, err)

				// Make sure we've computed the expected root
				expectedRoot, err := hex.DecodeString(p.ExpectedRoot[2:])
				require.NoError(t, err)

				r := i.Root()
				require.Equal(t, r[:], expectedRoot)

				paths := [imt.TreeDepth][]byte{}
				for idx, path := range p.Path {
					pBytes, err := hex.DecodeString(path[2:])
					require.NoError(t, err)
					paths[idx] = pBytes
				}

				// Make sure we get the expected branch root
				proofRoot, err := imt.BranchRoot(leaf, paths, p.Index)
				require.NoError(t, err)
				require.Equal(t, proofRoot[:], expectedRoot)
			}
		})
	}
}
