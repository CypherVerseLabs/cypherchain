package blockchain

import (
	"strconv"
	"time"
)

// generateNewBlockWithPoS generates a new block in the blockchain using Proof of Stake.
func generateNewBlockWithPoS(prevBlock Block, data string, difficulty int, validators []string) Block {
	timestamp := time.Now().Unix()
	newBlock := Block{
		Index:      prevBlock.Index + 1,
		Timestamp:  timestamp,
		PrevHash:   prevBlock.Hash,
		Data:       data,
		Nonce:      0,
		Difficulty: difficulty,
	}

	// Select a random validator
	randValidatorIndex := rand.Intn(len(validators))
	validator := validators[randValidatorIndex]

	newBlock.Hash = calculateHash(newBlock, validator)

	return newBlock
}

// createGenesisBlockForPoS creates the first block in the blockchain (genesis block) for Proof of Stake.
func createGenesisBlockForPoS(difficulty int, data string) Block {
	timestamp := time.Now().Unix()
	genesisBlock := Block{
		Index:      0,
		Timestamp:  timestamp,
		PrevHash:   "0",
		Data:       data,
		Nonce:      0,
		Difficulty: difficulty,
	}
	
	genesisBlock.Validators = validators
	
	genesisBlock.Hash = calculateHash(genesisBlock)

	return genesisBlock
}

