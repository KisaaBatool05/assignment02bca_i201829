package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

const NumberOfTransactionsPerBlock int = 10

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
	DateTime     time.Time
}

type Blockchain struct {
	Chain []Block
}

func CreateBlockchain() *Blockchain {
	var genesisBlock Block

	genesisBlock = Block{
		Transaction:  "Genesis block",
		Nonce:        0,
		PreviousHash: "",
		Hash:         calculateHash(&genesisBlock),
		DateTime:     time.Now(),
	}

	blockchain := Blockchain{}
	blockchain.Chain = append(blockchain.Chain, genesisBlock)
	return &blockchain
}

func (blockchain *Blockchain) AddBlock(newBlock Block) {
	blockchain.Chain = append(blockchain.Chain, newBlock)
}

func (blockchain *Blockchain) DisplayBlocks() {
	for _, block := range blockchain.Chain {
		fmt.Println("------------------------------")
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("DateTime: %s\n", block.DateTime.String())
		fmt.Println("------------------------------")
	}
}

func (blockchain *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	blockchain.Chain[blockIndex].Transaction = newTransaction
}

func (blockchain *Blockchain) VerifyChain() bool {
	for i := 1; i < len(blockchain.Chain); i++ {
		currentBlock := blockchain.Chain[i]
		previousBlock := blockchain.Chain[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}

		if calculateHash(&currentBlock) != currentBlock.Hash {
			return false
		}
	}

	return true
}

func calculateHash(block *Block) string {
	data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
