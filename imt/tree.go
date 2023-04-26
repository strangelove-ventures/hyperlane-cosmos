package imt

import (
	"fmt"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"errors"
)

const (
	TreeDepth = 32
	MaxLeaves = 4294967295
)

type Tree struct {
	branch [TreeDepth][32]byte;
	count int;
}

func (t* Tree) Insert(node [32]byte) error {
	if t.count >= MaxLeaves {
		return errors.New("merkle tree full")
	}

	t.count += 1
	size := t.count
	for i := 0; i < TreeDepth; i++ {
		if (size & 1) == 1 {
			t.branch[i] = node;
			return nil;
		}
		temp := make([]byte, 64)
		copy(temp[0:], t.branch[i][0:])
		copy(temp[32:], node[0:])
		hash := crypto.Keccak256(temp);
		copy(node[0:], hash[0:]);
		size /= 2;
	}

	return errors.New("unreachable")
}

func (t* Tree) Count() int {
	return t.count
}

func (t* Tree) Print() {
	for i := 0; i < TreeDepth; i ++ {
		fmt.Printf("%02d: %X\n", i, t.branch[i])
	}
}

func (t* Tree) Root() [32]byte { 
	return t.RootWithContext(t.ZeroHashes())
}

func (t* Tree) RootWithContext(zeroes [][]byte) [32]byte {
	current := [32]byte{}
	index := t.count;
	for i := 0; i < TreeDepth; i++ {
		ithBit := (index >> i) & 0x01;
		next := t.branch[i];

		temp := make([]byte, 64)
		if ithBit == 1 {
			copy(temp[0:], next[0:])
			copy(temp[32:], current[0:])
		} else {
			copy(temp[0:], current[0:])
			copy(temp[32:], zeroes[i][0:])
		}
		hash := crypto.Keccak256(temp);
		copy(current[0:], hash[0:]);
	}
	return current;
}

func BranchRoot(item [32]byte, branch [TreeDepth][32]byte, index int) [32]byte {
	current := [32]byte{}
	copy(current[0:], item[0:])
	for i := 0; i < TreeDepth; i++ {
		ithBit := (index >> i) & 0x01;
		next := branch[i];

		temp := make([]byte, 64)
		if ithBit == 1 {
			copy(temp[0:], next[0:])
			copy(temp[32:], current[0:])
		} else {
			copy(temp[0:], current[0:])
			copy(temp[32:], next[0:])
		}
		hash := crypto.Keccak256(temp);
		copy(current[0:], hash[0:]);
	}
	return current;
}

func (t* Tree) ZeroHashes() [][]byte {
	zeroes := []string{
	"0000000000000000000000000000000000000000000000000000000000000000",
	"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
	"b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30",
	"21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85",
	"e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a19344",
	"0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d",
	"887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a1968",
	"ffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f83",
	"9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af",
	"cefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e0",
	"f9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a5",
	"f8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf892",
	"3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c",
	"c1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb",
	"5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc",
	"da7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d2",
	"2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f",
	"e1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a",
	"5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0",
	"b46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0",
	"c65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2",
	"f4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd9",
	"5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e377",
	"4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652",
	"cdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef",
	"0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d",
	"b8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0",
	"838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e",
	"662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e",
	"388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea322",
	"93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735",
	"8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9",
	}

	zBytes := [][]byte{}

	for _,i := range zeroes {
		b, _ := hex.DecodeString(i)
		zBytes = append(zBytes, b)
	}
	return zBytes;
}