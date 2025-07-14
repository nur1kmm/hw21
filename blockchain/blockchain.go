package blockchain

import (
	"go.etcd.io/bbolt"
	"log"
)

// Blockchain represents the blockchain structure
type Blockchain struct {
	db *bbolt.DB
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain() *Blockchain {
	dbFile := "blockchain.db"
	db, err := bbolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// Check if the blocks bucket exists
	err = db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("blocks"))
		if bucket == nil {
			// Create the blocks bucket and genesis block
			bucket, err := tx.CreateBucket([]byte("blocks"))
			if err != nil {
				return err
			}
			
			genesis := NewGenesisBlock()
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}
			
			err = bucket.Put([]byte("l"), genesis.Hash)
			if err != nil {
				return err
			}
		}
		
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{db: db}
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	
	// Get the hash of the last block
	err := bc.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	
	// Create and mine the new block
	newBlock := NewBlock(data, lastHash)
	
	// Save the new block to the database
	err = bc.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}
		
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return err
		}
		
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// GetBlocks returns all blocks in the blockchain
func (bc *Blockchain) GetBlocks() []*Block {
	var blocks []*Block
	
	bci := bc.Iterator()
	
	for {
		block := bci.Next()
		if block == nil {
			break
		}
		blocks = append(blocks, block)
		
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	
	return blocks
}

// Iterator returns a blockchain iterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	var tip []byte
	
	err := bc.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		tip = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	
	return &BlockchainIterator{db: bc.db, currentHash: tip}
}

// Close closes the database connection
func (bc *Blockchain) Close() {
	err := bc.db.Close()
	if err != nil {
		log.Panic(err)
	}
}