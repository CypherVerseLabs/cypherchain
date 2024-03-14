package blockchain

import (
	"math/rand"
	"strconv"
	"time"
)

// generateNewBlockWithDPoS generates a new block in the blockchain using Delegated Proof of Stake.
func generateNewBlockWithDPoS(prevBlock Block, data string, difficulty int, validators []Validator) Block {
	timestamp := time.Now().Unix()
	newBlock := Block{
		Index:      prevBlock.Index + 1,
		Timestamp:  timestamp,
		PrevHash:   prevBlock.Hash,
		Data:       data,
		Nonce:      0,
		Difficulty: difficulty,
	}

	// Select a validator based on DPoS
	delegate := selectDelegate(validators)

	// Update the block with the selected delegate
	newBlock.Validator = delegate.Name

	// Calculate the hash using DPoS
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

// selectDelegate chooses a delegate based on a DPoS mechanism.
func selectDelegate(validators []Validator) Validator {
	rand.Seed(time.Now().UnixNano())
	validatorIndex := rand.Intn(len(validators))
	return validators[validatorIndex]
}
