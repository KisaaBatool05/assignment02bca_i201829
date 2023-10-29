package main

import (
	"blockchain"
	"fmt"
	"math/big"
)

func main() {
	// Ask the user for the number of transactions per block.
	fmt.Println("Enter the number of transactions per block:")
	var numberOfTransactionsPerBlock int
	fmt.Scanln(&numberOfTransactionsPerBlock)

	// Set the number of transactions per block.
	blockchain.SetNumberOfTransactionsPerBlock(numberOfTransactionsPerBlock)

	// Create a new blockchain.
	blockchain := blockchain.CreateBlockchain()

	// Get the current block.
	currentBlock := blockchain.Chain[len(blockchain.Chain)-1]

	// Add transactions to the current block.
	for i := 0; i < numberOfTransactionsPerBlock; i++ {
		// Get the next transaction from the user.
		fmt.Println("Enter the next transaction:")
		var transaction string
		fmt.Scanln(&transaction)

		// Add the transaction to the current block.
		currentBlock.Transaction += transaction
	}

	// Create a Merkle tree for the current block.
	merkleTree := blockchain.CreateMerkleTree(currentBlock.Transaction)

	// Calculate the Merkle root.
	merkleRoot := merkleTree.Root()

	// Set the Merkle root for the current block.
	currentBlock.MerkleRoot = merkleRoot

	// Mine the current block.
	pow := blockchain.NewProofOfWork(currentBlock, big.NewInt(1))
	nonce, hash := pow.Mine()

	if nonce > 0 {
		// Add the block to the blockchain.
		blockchain.AddBlock(currentBlock)
	}
}
