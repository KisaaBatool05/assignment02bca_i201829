package blockchain

import (
	"math/big"
	"time"
)

type block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
	DateTime     time.Time
}

type ProofOfWork struct {
	Block  *block
	Target *big.Int
}
