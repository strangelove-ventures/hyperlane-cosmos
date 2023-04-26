package imt_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
    "testing"
	"github.com/stretchr/testify/require"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
)

type MerkleProof struct {
	Leaf string;
	Index int;
	Path []string;
}

type MerkleVector struct {
	TestName string;
	ExpectedRoot string;
	Leaves []string;
	Proofs []MerkleProof;
}

func TestVectors(t* testing.T) {
	var cases []MerkleVector

	// Open the test cases
	jsonFile, err := os.Open("merkle.json")
	require.Nil(t, err)

	// Read them in
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cases)

	// Test them
	for _, c := range cases {
		t.Run(fmt.Sprintf("Testing case: %s", c.TestName), func(t *testing.T){
			i := imt.Tree{}

			// Insert all of the leaves
			for idx, l := range c.Leaves {
				//Test Vectors used 'ethers.utils.hashMessage(leaf)'
				//to compute the value to insert.
				msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(l), l)
				hash := crypto.Keccak256([]byte(msg))

				// Make sure we get the expected digest
				expectedLeaf, err := hex.DecodeString(c.Proofs[idx].Leaf[2:])
				require.Nil(t, err)
				require.Equal(t, hash[:], expectedLeaf)

				// Insert into the tree
				var h [32]byte
				copy(h[:], hash)
				i.Insert(h)
			}

			// Make sure we've inserted the correct amount
			require.Equal(t, i.Count(), len(c.Leaves))

			// Make sure we've computed the expected root
			expectedRoot, err := hex.DecodeString(c.ExpectedRoot[2:])
			require.Nil(t, err)
			r := i.Root()
			require.Equal(t, r[:], expectedRoot)

			// Verify leaves
			for _, p := range c.Proofs {
				leaf, err := hex.DecodeString(p.Leaf[2:])
				require.Nil(t, err)

				paths := [32][32]byte{}
				for idx, path  := range p.Path {
					pBytes, err := hex.DecodeString(path[2:])
					require.Nil(t, err)
					copy(paths[idx][:], pBytes)
				}

				// Make sure we get the expected branch root
				var l [32]byte
				copy(l[:], leaf)
				proofRoot := imt.BranchRoot(l, paths, p.Index)
				require.Equal(t, proofRoot[:], expectedRoot)
			}
		})
	}
}