package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// CalculateHash calculates the hash of a block.
func CalculateHash(block *Block) string {
	record := strconv.Itoa(block.Index) +
		strconv.FormatInt(block.Timestamp, 10) +
		block.PrevHash +
		block.Data +
		strconv.Itoa(block.Nonce) +
		strconv.Itoa(block.Difficulty)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

