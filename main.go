package main

import (
	"fmt"
)

func main() {
	fmt.Println("Blockchain implementation in Go")
	
	// Create a new blockchain
	bc := NewBlockchain()
	
	// Add some blocks to the blockchain
	bc.AddBlock("First block")
	bc.AddBlock("Second block")
	bc.AddBlock("Third block")
	
	// Print all blocks
	for i, block := range bc.GetBlocks() {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("  Timestamp: %d\n", block.Timestamp)
		fmt.Printf("  Data: %s\n", block.Data)
		fmt.Printf("  Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("  Hash: %x\n", block.Hash)
		fmt.Println()
	}
}