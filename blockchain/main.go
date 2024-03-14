package main

import (
	"fmt"
	"https://github.com/CypherVerseLabs/cypherchain"
)

func main() {
	difficulty := 3 // Adjust the difficulty level as needed

	// Test Proof of Work
	powBlockchain := []blockchain.Block{blockchain.CreateGenesisBlock(difficulty)}

	// Generate a new block with some data using Proof of Work
	newBlockData := "Block Data"
	newPowBlock := blockchain.GenerateNewBlockWithPoW(powBlockchain[len(powBlockchain)-1], newBlockData, difficulty)
	powBlockchain = append(powBlockchain, newPowBlock)

	// Print the Proof of Work blockchain
	fmt.Println("Proof of Work Blockchain:")
	printBlockchain(powBlockchain)

	// Test Proof of Stake
	posBlockchain := []blockchain.Block{blockchain.CreateGenesisBlockForPoS(difficulty)}

	// Generate a new block with some data using Proof of Stake
	newPosBlock := blockchain.GenerateNewBlockWithPoS(posBlockchain[len(posBlockchain)-1], newBlockData, difficulty, []string{"Validator 1", "Validator 2", "Validator 3"})
	posBlockchain = append(posBlockchain, newPosBlock)

	// Print the Proof of Stake blockchain
	fmt.Println("Proof of Stake Blockchain:")
	printBlockchain(posBlockchain)
}

func printBlockchain(bc []blockchain.Block) {
	for _, block := range bc {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Difficulty: %d\n", block.Difficulty)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

