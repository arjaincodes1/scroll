package utils

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/crypto"
)

// ComputeBatchID compute a unique hash for a batch using "endBlockHash" & "endBlockHash in last batch"
// & "batch height", following the logic in `_computeBatchId` in contracts/src/L1/rollup/ZKRollup.sol
func ComputeBatchID(endBlockHash common.Hash, lastEndBlockHash common.Hash, index *big.Int) string {
	indexBytes := make([]byte, 32)
	return crypto.Keccak256Hash(
		endBlockHash.Bytes(),
		lastEndBlockHash.Bytes(),
		index.FillBytes(indexBytes),
	).String()
}