/*
Package merklee provides a simple implementation of a Merkle tree.

A Merkle tree is a binary tree in which every leaf node is labelled with the cryptographic hash of a data block, 
and every non-leaf node is labelled with the cryptographic hash of the labels of its child nodes. 
This structure allows efficient and secure verification of the contents of large data structures.

The main features of this package include:
- Creating a Merkle tree from a list of data blocks.
- Generating the Merkle root, which is the hash at the root of the tree.
- Verifying the inclusion of a data block in the tree using a Merkle proof.

Example usage:

	package main

	import (
		"fmt"
		"github.com/zacksfF/Zackchain/crypto/merklee"
	)

	func main() {
		data := [][]byte{
			[]byte("block1"),
			[]byte("block2"),
			[]byte("block3"),
			[]byte("block4"),
		}

		tree := merklee.NewMerkleTree(data)
		fmt.Printf("Merkle Root: %x\n", tree.Root())

		proof, _ := tree.GenerateProof(data[0])
		fmt.Printf("Proof for block1: %x\n", proof)
	}

*/
package merklee