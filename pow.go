package blockchain

import (
 "math"
 "math/big"
 "strconv"
 "time"
)

// generateNewBlockWithPoW generates a new block in the blockchain using Proof of Work.
func generateNewBlockWithPoW(prevBlock Block, data string, difficulty int) Block {
 var nonce int
 timestamp := time.Now().Unix()
 newBlock := Block{
  Index:        prevBlock.Index + 1,
  Timestamp:    timestamp,
  PrevHash:     prevBlock.Hash,
  Data:         data,
  Nonce:        0,
  Difficulty:   difficulty,
 }

 // Perform the proof of work
 target := big.NewInt(1)
 target.Lsh(target, uint(256-difficulty))
 for nonce < math.MaxInt64 {
  newBlock.Nonce = nonce
  hash := calculateHash(newBlock)

  hashInt := new(big.Int)
  hashInt.SetString(hash, 16)

  if hashInt.Cmp(target) == -1 {
   newBlock.Hash = hash
   break
  } else {
   nonce++
  }
 }

 return newBlock
}
