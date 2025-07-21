package blockchain

import (
	"errors"
	"go.etcd.io/bbolt"
)

const (
	blocksBucket = "blocks"
	lastHashKey  = "l"
)

var (
	ErrBlockNotFound = errors.New("block not found")
)

// Database represents the blockchain database
type Database struct {
	db *bbolt.DB
}

// NewDatabase creates a new database instance
func NewDatabase(dbFile string) (*Database, error) {
	db, err := bbolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(blocksBucket))
		return err
	})
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// Close closes the database connection
func (db *Database) Close() {
	db.db.Close()
}

// GetLastBlockHash retrieves the hash of the last block in the blockchain
func (db *Database) GetLastBlockHash() ([]byte, error) {
	var lastHash []byte
	err := db.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte(lastHashKey))
		return nil
	})
	return lastHash, err
}

// AddBlock adds a block to the database
func (db *Database) AddBlock(block *Block) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		return b.Put(block.Hash, block.Serialize())
	})
}

// UpdateLastBlockHash updates the last block hash in the database
func (db *Database) UpdateLastBlockHash(hash []byte) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		return b.Put([]byte(lastHashKey), hash)
	})
}

// GetBlock retrieves a block from the database by its hash
func (db *Database) GetBlock(hash []byte) (*Block, error) {
	var block *Block
	err := db.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		blockData := b.Get(hash)
		if blockData == nil {
			return ErrBlockNotFound
		}
		block = DeserializeBlock(blockData)
		return nil
	})
	return block, err
}

// Iterator returns a blockchain iterator
func (db *Database) Iterator() *BlockchainIterator {
	lastHash, _ := db.GetLastBlockHash()
	return &BlockchainIterator{
		currentHash: lastHash,
		db:          db.db,
	}
}
