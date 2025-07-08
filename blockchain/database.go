package blockchain

import (
	"go.etcd.io/bbolt"
	"log"
)

// Database represents the blockchain database
type Database struct {
	db *bbolt.DB
}

// NewDatabase creates a new database instance
func NewDatabase(dbFile string) *Database {
	db, err := bbolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// Create the blocks bucket if it doesn't exist
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("blocks"))
		return err
	})
	if err != nil {
		log.Panic(err)
	}

	return &Database{db: db}
}

// Close closes the database connection
func (db *Database) Close() {
	err := db.db.Close()
	if err != nil {
		log.Panic(err)
	}
}

// AddBlock adds a block to the database
func (db *Database) AddBlock(block *Block) {
	err := db.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		serialized := block.Serialize()
		return b.Put(block.Hash, serialized)
	})
	if err != nil {
		log.Panic(err)
	}
}

// GetBlock retrieves a block from the database by its hash
func (db *Database) GetBlock(hash []byte) *Block {
	var block *Block

	err := db.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		serialized := b.Get(hash)
		if serialized == nil {
			return nil
		}
		block = DeserializeBlock(serialized)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return block
}

// GetBlockchainIterator returns an iterator for the blockchain
func (db *Database) GetBlockchainIterator() *BlockchainIterator {
	return &BlockchainIterator{db: db.db, currentHash: []byte{}}
}