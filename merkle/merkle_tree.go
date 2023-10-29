package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

type Node struct {
	Data  string
	Hash  string
	Left  *Node
	Right *Node
}

type MerkleTree struct {
	Root *Node
}

func NewMerkleTree(data []string) *MerkleTree {
	tree := MerkleTree{}

	// Build the Merkle tree from the bottom up
	nodes := make([]*Node, len(data))
	for i := 0; i < len(data); i++ {
		nodes[i] = &Node{Hash: calculateHash(data[i])}
	}

	for len(nodes) > 1 {
		// Combine the nodes into pairs
		newNodes := make([]*Node, len(nodes)/2)
		for i := 0; i < len(newNodes); i++ {
			leftNode := nodes[2*i]
			rightNode := nodes[2*i+1]

			// Create a new node with the hash of the two child nodes
			newNode := &Node{Hash: calculateHash(leftNode.Hash + rightNode.Hash)}
			newNode.Left = leftNode
			newNode.Right = rightNode

			newNodes[i] = newNode
		}

		// Update the nodes list with the new nodes
		nodes = newNodes
	}

	// The root node of the Merkle tree is the last node in the nodes list
	tree.Root = nodes[0]

	return &tree
}
func calculateHash(data string) string {
	// Calculate the hash of the data.
	hash := sha256.Sum256([]byte(data))

	// Convert the [32]byte type to a []byte type.
	bytes := make([]byte, len(hash))
	copy(bytes, hash[:])

	// Encode the hash to a string.
	encodedString := hex.EncodeToString(bytes)

	return encodedString
}

// Get the Merkle root hash of the tree
func (tree *MerkleTree) GetRootHash() string {
	return tree.Root.Hash
}

// Generate a Merkle proof for the given data
func (tree *MerkleTree) GenerateProof(data string) []*Node {
	// Find the node in the Merkle tree that corresponds to the given data
	node := tree.FindNode(data)

	// Generate a list of nodes from the root node to the given node
	proof := []*Node{}
	for node != nil {
		proof = append(proof, node)
		node = node.Parent()
	}

	return proof
}

// Verify the given Merkle proof for the given data
func (tree *MerkleTree) VerifyProof(data string, proof []*Node) bool {
	// Calculate the Merkle root hash using the given proof
	calculatedMerkleRootHash := calculateMerkleRootHash(proof)

	// Compare the calculated Merkle root hash to the Merkle root hash of the tree
	return calculatedMerkleRootHash == tree.GetRootHash()
}

func calculateMerkleRootHash(proof []*Node) string {
	// Calculate the hash of the first two nodes in the proof
	hash := calculateHash(proof[0].Hash + proof[1].Hash)

	// Combine the rest of the nodes in the proof, two at a time, and calculate the hash of each pair
	for i := 2; i < len(proof); i += 2 {
		hash = calculateHash(hash + proof[i].Hash)
	}

	return hash
}

func (node *Node) Parent() *Node {
	if node.Left != nil {
		return node.Left
	} else {
		return node.Right
	}
}

func (tree *MerkleTree) FindNode(data string) *Node {
	// Start at the root node of the Merkle tree.
	node := tree.Root

	// Recursively traverse the Merkle tree until we find the node that corresponds to the given data.
	for node != nil {
		if node.Data == data {
			return node
		}

		if data < node.Data {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	// If we reach this point, then the node was not found in the Merkle tree.
	return nil
}
