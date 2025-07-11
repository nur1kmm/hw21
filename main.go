package main

import (
	"fmt"
	
	"github.com/nur1kmm/hw21/blockchain"
)

// main is the command-line entry point for the blockchain
func main() {
	fmt.Println("Blockchain implementation in Go (Command-line version)")
	
	// Create a new blockchain
	bc := blockchain.NewBlockchain()
	defer bc.Close()
	
	// Add some blocks to the blockchain
	bc.AddBlock("First block")
	bc.AddBlock("Second block")
	bc.AddBlock("Third block")
	
	// Print all blocks
	bci := bc.Iterator()
	
	for {
		block := bci.Next()
		if block == nil {
			break
		}
		
		fmt.Printf("Block %x:\n", block.Hash)
		fmt.Printf("  Timestamp: %d\n", block.Timestamp)
		fmt.Printf("  Data: %s\n", block.Data)
		fmt.Printf("  Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("  Hash: %x\n", block.Hash)
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Println()
		
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}