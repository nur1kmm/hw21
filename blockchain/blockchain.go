package blockchain

import (
	"log"
)

const dbFile = "blockchain.db"

// Blockchain represents the blockchain structure
type Blockchain struct {
	db *Database
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain() (*Blockchain, error) {
	db, err := NewDatabase(dbFile)
	if err != nil {
		return nil, err
	}

	lastHash, err := db.GetLastBlockHash()
	if err != nil {
		return nil, err
	}

	if lastHash == nil {
		genesis := NewGenesisBlock()
		if err := db.AddBlock(genesis); err != nil {
			return nil, err
		}
		if err := db.UpdateLastBlockHash(genesis.Hash); err != nil {
			return nil, err
		}
	}

	return &Blockchain{db: db}, nil
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) error {
	lastHash, err := bc.db.GetLastBlockHash()
	if err != nil {
		return err
	}

	newBlock := NewBlock(data, lastHash)

	if err := bc.db.AddBlock(newBlock); err != nil {
		return err
	}

	if err := bc.db.UpdateLastBlockHash(newBlock.Hash); err != nil {
		return err
	}

	return nil
}

// GetBlocks returns all blocks in the blockchain
func (bc *Blockchain) GetBlocks() ([]*Block, error) {
	var blocks []*Block
	iter := bc.Iterator()

	for {
		block, err := iter.Next()
		if err != nil {
			if err == ErrBlockNotFound {
				break
			}
			return nil, err
		}

		blocks = append(blocks, block)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return blocks, nil
}

// Iterator returns a blockchain iterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return bc.db.Iterator()
}

// Close closes the database connection
func (bc *Blockchain) Close() {
	bc.db.Close()
}
