package main

// Blockchain represents the blockchain structure
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{blocks: []*Block{NewGenesisBlock()}}
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	// No need to call SetHash anymore as it's done in NewBlock with PoW
	bc.blocks = append(bc.blocks, newBlock)
}

// GetBlocks returns all blocks in the blockchain
func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}