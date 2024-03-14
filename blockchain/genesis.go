package blockchain

import (
	"strconv"
	"time"
)

// createGenesisBlock creates the first block in the blockchain (genesis block).
func createGenesisBlock(difficulty int, data string) Block {
	timestamp := time.Now().Unix()
	genesisBlock := Block{
		Index:      0,
		Timestamp:  timestamp,
		PrevHash:   "0",
		Data:       data, // Customize the data field if needed
		Nonce:      0,
		Difficulty: difficulty,
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}

