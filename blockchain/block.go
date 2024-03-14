package blockchain

import (
 "bytes"
 "crypto/sha256"
 "encoding/hex"
 "strconv"
 "time"
)

// Block represents a block in the blockchain.
type Block struct {
 Index        int
 Timestamp    int64
 PrevHash     string
 Data         string
 Nonce        int
 Difficulty   int
 Hash         string
}
